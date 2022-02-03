package main

import (
	"github.com/gaulzhw/learning-k8s/restful"
)

func main() {
	// go-restful
	//restful.StartContainer()
	restful.StartServer()

	// client
	//client.InitRestClient()
	//client.InitDynamicClient()
	//client.InitClientSet()
	//client.InitDiscoveryClient()

	// informer
	//clientgo.StartInformer()

	// controller-manager
	//clientgo.StartController()

	// AA
	//aggregateapiserver.Start()
}
