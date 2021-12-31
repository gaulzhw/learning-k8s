package main

import (
	"github.com/gaulzhw/learning-k8s/informer"
)

func main() {
	// clientgo.InitRestClient()
	// clientgo.InitClientSet()
	// clientgo.InitDynamicClient()
	// clientgo.InitDiscoveryClient()

	//etcd.NewEtcdClient()

	informer.Start()
}
