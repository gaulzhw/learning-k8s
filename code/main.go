package main

import (
	"github.com/gaulzhw/learning-k8s/clientgo"
)

func main() {
	// clientgo.InitRestClient()
	// clientgo.InitClientSet()
	// clientgo.InitDynamicClient()
	clientgo.InitDiscoveryClient()
}
