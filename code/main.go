package main

import (
	"github.com/gaulzhw/learning-k8s/aggregateapiserver"
)

func main() {
	// informer
	//clientgo.Start()

	// controller-manager
	//clientgo.StartController()
	//clientgo.StartInformer()

	// AA
	aggregateapiserver.Start()
}
