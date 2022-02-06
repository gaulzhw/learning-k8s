module github.com/gaulzhw/learning-k8s/hack/tools

go 1.17

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/drone/envsubst/v2 v2.0.0-20210730161058-179042472c46
	github.com/hashicorp/go-multierror v1.1.1
	github.com/joelanford/go-apidiff v0.1.0
	github.com/onsi/ginkgo v1.16.5
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/net v0.0.0-20211118161319-6a13c67c3ce4
	golang.org/x/tools v0.1.8-0.20211029000441-d6a9af8af023
	gotest.tools/gotestsum v1.6.4
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	k8s.io/code-generator v0.23.0
	k8s.io/klog/v2 v2.30.0
	sigs.k8s.io/cluster-api v0.0.0-00010101000000-000000000000
	sigs.k8s.io/cluster-api/test v0.0.0-00010101000000-000000000000
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/controller-runtime/tools/setup-envtest v0.0.0-20211110210527-619e6b92dab9
	sigs.k8s.io/controller-tools v0.8.0
	sigs.k8s.io/kubebuilder/docs/book/utils v0.0.0-20211028165026-57688c578b5d
	sigs.k8s.io/promo-tools/v3 v3.3.0-beta.3
)
