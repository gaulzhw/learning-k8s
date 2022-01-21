package aggregateapiserver

import (
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"
)

// https://github.com/kubernetes-sigs/apiserver-runtime/blob/main/sample/cmd/apiserver/main.go
func Start() {
	logs.InitLogs()
	defer logs.FlushLogs()

	err := builder.APIServer.
		//WithOpenAPIDefinitions("sample", "v0.0.0", openapi.GetOpenAPIDefinitions).
		//WithResource(&v1alpha1.Flunder{}). // namespaced resource
		//WithResource(&v1alpha1.Fischer{}). // non-namespaced resource
		//WithResource(&v1alpha1.Fortune{}). // resource with custom rest.Storage implementation
		WithoutEtcd().
		//WithLocalDebugExtension().
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
