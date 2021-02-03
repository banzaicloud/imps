module github.com/banzaicloud/backyards/services/imps

go 1.15

require (
	emperror.dev/emperror v0.32.0
	github.com/banzaicloud/backyards v0.0.1
	github.com/banzaicloud/backyards/pkg/platform/log v0.0.0
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v12.0.0+incompatible
	logur.dev/integration/logr v0.3.2
	logur.dev/logur v0.17.0
	sigs.k8s.io/controller-runtime v0.6.2
)

replace (
	github.com/banzaicloud/backyards => ../..
	github.com/banzaicloud/backyards/pkg/common => ../../pkg/common
	github.com/banzaicloud/backyards/pkg/platform/log => ../../pkg/platform/log
	github.com/banzaicloud/backyards/pkg/prometheus => ../../pkg/prometheus
	github.com/banzaicloud/backyards/pkg/template => ../../pkg/template

	github.com/banzaicloud/backyards/services/imps/static => ./static
	github.com/banzaicloud/backyards/services/sre => ../../services/sre
	k8s.io/client-go => k8s.io/client-go v0.19.2
)
