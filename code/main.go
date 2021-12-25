package main

import (
	"github.com/gaulzhw/learning-k8s/etcd"
)

func main() {
	// clientgo.InitRestClient()
	// clientgo.InitClientSet()
	// clientgo.InitDynamicClient()
	// clientgo.InitDiscoveryClient()

	etcd.NewEtcdClient()
}
