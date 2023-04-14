// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	// +kubebuilder:scaffold:imports
	"github.com/banzaicloud/operator-tools/pkg/reconciler"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	logrintegration "logur.dev/integration/logr"
	"logur.dev/logur"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/banzaicloud/imps/api/v1alpha1"
	"github.com/banzaicloud/imps/controllers"
	"github.com/banzaicloud/imps/internal/errorhandler"
	"github.com/banzaicloud/imps/internal/log"
	"github.com/banzaicloud/imps/pkg/ecr"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = v1alpha1.AddToScheme(scheme)

	// +kubebuilder:scaffold:scheme
}

type Config struct {
	Log log.Config
}

func main() {
	Configure(viper.GetViper(), pflag.CommandLine)

	var metricsAddr string
	var periodicReconcileInterval int
	var enableLeaderElection bool
	var configNamespace string

	pflag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	pflag.IntVar(&periodicReconcileInterval, "periodic-reconcile-interval", 300, "The interval in seconds in which controller reconciles are run periodically.")
	pflag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	pflag.StringVar(&configNamespace, "config-namespace", "backyards-system", "The namespace in which internal resources should be created for leader election.")

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		setupLog.Error(err, "failed to unmarshal configuration")
		os.Exit(1)
	}

	pflag.Parse()

	periodicReconcileIntervalDuration := time.Duration(periodicReconcileInterval) * time.Second

	// Create logger (first thing after configuration loading)
	logger := log.NewLogger(config.Log)
	ctrl.SetLogger(logrintegration.New(logger))

	errorHandler := errorhandler.New(logger)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                  scheme,
		MetricsBindAddress:      metricsAddr,
		Port:                    9443,
		LeaderElection:          enableLeaderElection,
		LeaderElectionID:        "73de1ad9.banzaicloud.io",
		LeaderElectionNamespace: configNamespace,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	ecrLogger := logur.WithField(logger, "controller", "ecr_token_refresh")
	ecr.Initialize(ecrLogger)

	impsLogger := logur.WithField(logger, "controller", "imagepullsecrets")
	impsReconciler := &controllers.ImagePullSecretReconciler{
		Client:                    mgr.GetClient(),
		Log:                       impsLogger,
		ErrorHandler:              errorHandler,
		ResourceReconciler:        reconciler.NewReconcilerWith(mgr.GetClient(), reconciler.WithLog(logrintegration.New(impsLogger))),
		Recorder:                  mgr.GetEventRecorderFor("imagepullsecrets-controller"),
		PeriodicReconcileInterval: periodicReconcileIntervalDuration,
	}

	if err = impsReconciler.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "imagepullsecrets")
		os.Exit(1)
	}

	// +kubebuilder:scaffold:builder

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
