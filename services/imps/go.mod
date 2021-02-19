module github.com/banzaicloud/backyards/services/imps

go 1.15

require (
	emperror.dev/emperror v0.33.0
	emperror.dev/errors v0.8.0
	github.com/aws/aws-sdk-go-v2 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.1.1
	github.com/banzaicloud/backyards/pkg/common v0.0.0
	github.com/banzaicloud/operator-tools v0.16.1
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v12.0.0+incompatible
	logur.dev/integration/logr v0.3.2
	logur.dev/logur v0.17.0
	sigs.k8s.io/controller-runtime v0.6.2
)

replace (
	github.com/banzaicloud/backyards/pkg/common => ../../pkg/common
	github.com/banzaicloud/backyards/services/health => ../../services/health

	github.com/banzaicloud/backyards/services/imps/static => ./static
	github.com/banzaicloud/backyards/services/sre => ../../services/sre
	k8s.io/client-go => k8s.io/client-go v0.19.2
)
