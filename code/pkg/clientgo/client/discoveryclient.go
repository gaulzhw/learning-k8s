package client

import (
	"errors"
	"path/filepath"

	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func newDiscoveryClient() (discovery.DiscoveryInterface, error) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		return nil, err
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, err
	}
	return discoveryClient, nil
}

func newFakeDiscoveryClient() (discovery.DiscoveryInterface, error) {
	clientset := newFakeClientSet()

	fakeDiscovery, ok := clientset.Discovery().(*fakediscovery.FakeDiscovery)
	if !ok {
		return nil, errors.New("convert error")
	}
	return fakeDiscovery, nil
}
