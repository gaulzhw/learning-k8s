package main

import (
	"flag"
	"os"

	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"

	"github.com/gaulzhw/learning-k8s/sample-apiserver/pkg/server"
)

// https://github.com/kubernetes-sigs/apiserver-runtime/blob/main/sample/cmd/apiserver/main.go
func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	stopCh := genericapiserver.SetupSignalHandler()
	options := server.NewAPIServerOptions(os.Stdout, os.Stderr)
	cmd := server.NewCommandStartAPIServer(options, stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
