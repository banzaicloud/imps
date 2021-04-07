package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	logrintegration "logur.dev/integration/logr"
	"logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/banzaicloud/imps/controllers"
	"github.com/banzaicloud/imps/internal/errorhandler"
	"github.com/banzaicloud/imps/internal/log"
	"github.com/banzaicloud/imps/pkg/ecr"
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
)

var (
	scheme              = runtime.NewScheme()
	setupLog            = ctrl.Log.WithName("setup")
	ErrInvalidReference = errors.New("invalid resource reference name")
	ErrNoSourceSecrets  = errors.New("no source secrets are specified")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	// +kubebuilder:scaffold:scheme
}

type Config struct {
	Log log.Config
}

func refToNamespacedName(name string) (*types.NamespacedName, error) {
	parts := strings.Split(name, ".")
	if len(parts) <= 1 || len(parts) > 2 {
		return nil, errors.WrapWithDetails(ErrInvalidReference, "reference", name)
	}

	return &types.NamespacedName{
		Namespace: parts[0],
		Name:      parts[1],
	}, nil
}

func main() {
	Configure(viper.GetViper(), pflag.CommandLine)
	var periodicReconcileInterval int
	var targetSecretString string
	var metricsAddr string

	sourceSecretStrings := pflag.StringArray("source-secret", nil, "Source secrets specified in <namespace>.<secret-name> format")
	pflag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	pflag.IntVar(&periodicReconcileInterval, "periodic-reconcile-interval", 300, "The interval in seconds in which controller reconciles are run periodically.")
	pflag.StringVar(&targetSecretString, "target-secret", "", "Target secret specifies what secret to create containing the image pull secrets. Format: namespace.secret-name")
	pflag.Parse()

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		setupLog.Error(err, "failed to unmarshal configuration")
		os.Exit(1)
	}

	// Parse command line arguments
	targetSecret, err := refToNamespacedName(targetSecretString)
	if err != nil {
		setupLog.Error(err, "failed to parse target secret name")
		os.Exit(1)
	}

	if len(*sourceSecretStrings) == 0 {
		setupLog.Error(ErrNoSourceSecrets, "please specify source secrets")
		os.Exit(1)
	}

	sourceSecrets := []types.NamespacedName{}
	for _, sourceSecertString := range *sourceSecretStrings {
		sourceSecret, err := refToNamespacedName(sourceSecertString)
		if err != nil {
			setupLog.Error(err, "failed to parse source secret name")
			os.Exit(1)
		}
		sourceSecrets = append(sourceSecrets, *sourceSecret)
	}

	// Create logger (first thing after configuration loading)
	logger := log.NewLogger(config.Log)
	ctrl.SetLogger(logrintegration.New(logger))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		MetricsBindAddress:      metricsAddr,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	ecrLogger := logur.WithField(logger, "controller", "ecr_token_refresh")
	ecr.Initialize(ecrLogger)

	errorHandler := errorhandler.New(logger)

	periodicReconcileIntervalDuration := time.Duration(periodicReconcileInterval) * time.Second

	refresherLogger := logur.WithField(logger, "controller", "imagepullsecrets")
	refresherReconciler := &controllers.RefresherReconciler{
		Client:                    mgr.GetClient(),
		Log:                       refresherLogger,
		ErrorHandler:              errorHandler,
		Scheme:                    mgr.GetScheme(),
		ResourceReconciler:        reconciler.NewReconcilerWith(mgr.GetClient(), reconciler.WithLog(logrintegration.New(refresherLogger))),
		PeriodicReconcileInterval: periodicReconcileIntervalDuration,
		SourceSecrets:             sourceSecrets,
		TargetSecret:              *targetSecret,
	}

	if err = refresherReconciler.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "imagepullsecrets")
		os.Exit(1)
	}

	setupLog.Info("starting manager")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

const FriendlyServiceName = "imps"

func Configure(v *viper.Viper, p *pflag.FlagSet) {
	v.AllowEmptyEnv(true)
	p.Init(FriendlyServiceName, pflag.ExitOnError)
	pflag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", FriendlyServiceName)
		pflag.PrintDefaults()
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	log.ConfigureLoggingFlags(v, p)

	_ = v.BindPFlags(p)
}
