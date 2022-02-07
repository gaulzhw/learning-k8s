package main

import (
	"github.com/gaulzhw/learning-k8s/sample/pkg/clientgo"
	"github.com/gaulzhw/learning-k8s/sample/pkg/clientgo/client"
)

func main() {
	// client
	client.InitRestClient()
	client.InitDynamicClient()
	client.InitClientSet()
	client.InitDiscoveryClient()

	// informer
	clientgo.StartInformer()

	// controller-manager
	clientgo.StartController()
}
