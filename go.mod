module github.com/banzaicloud/imps

go 1.16

require (
	emperror.dev/emperror v0.33.0
	emperror.dev/errors v0.8.0
	emperror.dev/handler/logur v0.5.0
	github.com/aws/aws-sdk-go-v2 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.1.1
	github.com/banzaicloud/imps/api v0.3.3
	github.com/banzaicloud/operator-tools v0.24.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
	logur.dev/adapter/logrus v0.5.0
	logur.dev/integration/logr v0.3.2
	logur.dev/logur v0.17.0
	sigs.k8s.io/controller-runtime v0.9.5
)

replace (
	github.com/banzaicloud/imps/api => ./api
	github.com/banzaicloud/imps/static => ./static
)
