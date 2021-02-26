module github.com/banzaicloud/imps

go 1.15

require (
	emperror.dev/emperror v0.33.0
	emperror.dev/errors v0.8.0
	emperror.dev/handler/logur v0.5.0
	github.com/aws/aws-sdk-go-v2 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.1.1
	github.com/banzaicloud/operator-tools v0.16.1
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	logur.dev/adapter/logrus v0.5.0
	logur.dev/integration/logr v0.3.2
	logur.dev/logur v0.17.0
	sigs.k8s.io/controller-runtime v0.6.2
)

replace github.com/banzaicloud/imps/static => ./static
