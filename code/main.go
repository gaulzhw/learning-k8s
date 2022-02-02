package main

import (
	"github.com/gaulzhw/learning-k8s/aggregateapiserver"
)

func main() {
	//restful.Start()

	// informer
	//clientgo.Start()

	// controller-manager
	//clientgo.StartController()
	//clientgo.StartInformer()

	// AA
	aggregateapiserver.Start()
}
