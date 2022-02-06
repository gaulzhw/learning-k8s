package main

import (
	"github.com/gaulzhw/learning-k8s/sample/clientgo"
	"github.com/gaulzhw/learning-k8s/sample/clientgo/client"
	"github.com/gaulzhw/learning-k8s/sample/restful"
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
