package main

import (
	"github.com/gaulzhw/learning-k8s/clientgo"
)

func main() {
	// informer
	//clientgo.Start()

	// controller-manager
	clientgo.StartController()
}
