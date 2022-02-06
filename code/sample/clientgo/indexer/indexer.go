package indexer

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

func userIndexFunc(obj interface{}) ([]string, error) {
	pod := obj.(*corev1.Pod)
	userString := pod.Annotations["users"]
	return strings.Split(userString, ","), nil
}

func newIndexer() cache.Indexer {
	return cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
		"byUser": userIndexFunc,
	})
}
