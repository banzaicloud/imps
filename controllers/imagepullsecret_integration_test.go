package controllers

// import (
// 	"log"
// 	"os"
// 	"path"
// 	"testing"
// 	"time"

// 	"k8s.io/apimachinery/pkg/runtime"
// 	"k8s.io/client-go/kubernetes"
// 	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
// 	"k8s.io/client-go/rest"
// 	logrintegration "logur.dev/integration/logr"
// 	"logur.dev/logur"
// 	ctrl "sigs.k8s.io/controller-runtime"
// 	"sigs.k8s.io/controller-runtime/pkg/envtest"

// 	"github.com/banzaicloud/imps/api/v1alpha1"
// 	"github.com/banzaicloud/imps/internal/errorhandler"
// 	logging "github.com/banzaicloud/imps/internal/log"
// 	"github.com/banzaicloud/imps/pkg/ecr"
// 	"github.com/cisco-open/operator-tools/pkg/reconciler"
// )

// var (
// 	k8sClient kubernetes.Interface
// 	testEnv   *envtest.Environment
// 	cfg       *rest.Config
// )

// func TestMain(m *testing.M) {
// 	setup()
// 	code := m.Run()
// 	teardown()
// 	os.Exit(code)
// }

// func setup() {
// 	var err error

// 	testEnv = &envtest.Environment{
// 		CRDDirectoryPaths: []string{path.Join("..", "config", "crd", "bases")},
// 	}

// 	cfg, err = testEnv.Start()
// 	if err != nil {
// 		log.Fatalf("error starting envtest: %v", err)
// 	}

// 	k8sClient, err = kubernetes.NewForConfig(cfg)
// 	if err != nil {
// 		log.Fatalf("error creating Kubernetes client: %v", err)
// 	}
// }

// func teardown() {
// 	err := testEnv.Stop()
// 	if err != nil {
// 		log.Fatalf("error stopping envtest: %v", err)
// 	}
// }

// func TestImagePullSecret_Setup(t *testing.T) {
// 	scheme := runtime.NewScheme()
// 	_ = clientgoscheme.AddToScheme(scheme)

// 	_ = v1alpha1.AddToScheme(scheme)

// 	logger := logging.NewLogger(logging.Config{})
// 	ctrl.SetLogger(logrintegration.New(logger))

// 	errorHandler := errorhandler.New(logger)

// 	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
// 		Scheme:                  scheme,
// 		MetricsBindAddress:      "0",
// 		Port:                    9443,
// 		LeaderElection:          false,
// 		LeaderElectionID:        "73de1ad9.banzaicloud.io",
// 		LeaderElectionNamespace: "backyards-system",
// 	})
// 	if err != nil {
// 		logger.Error("unable to start test manager")
// 		logger.Error(err.Error())
// 		os.Exit(1)
// 	}

// 	ecrLogger := logur.WithField(logger, "controller", "ecr_token_refresh")
// 	ecr.Initialize(ecrLogger)

// 	impsLogger := logur.WithField(logger, "controller", "imagepullsecrets")
// 	impsReconciler := &ImagePullSecretReconciler{
// 		Client:                    mgr.GetClient(),
// 		Log:                       impsLogger,
// 		ErrorHandler:              errorHandler,
// 		ResourceReconciler:        reconciler.NewReconcilerWith(mgr.GetClient(), reconciler.WithLog(logrintegration.New(impsLogger))),
// 		Recorder:                  mgr.GetEventRecorderFor("imagepullsecrets-controller"),
// 		PeriodicReconcileInterval: 300 * time.Second,
// 	}

// 	if err = impsReconciler.SetupWithManager(mgr); err != nil {
// 		logger.Error("unable to create test controller")
// 		os.Exit(1)
// 	}

// 	//impsReconciler.impsMatchingNamespace()
// }
