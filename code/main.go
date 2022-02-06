package main

import (
	"github.com/gaulzhw/learning-k8s/pkg/clientgo"
	"github.com/gaulzhw/learning-k8s/pkg/clientgo/client"
	"github.com/gaulzhw/learning-k8s/pkg/restful"
)

func main() {
	// go-restful
	restful.StartContainer()
	restful.StartServer()

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
