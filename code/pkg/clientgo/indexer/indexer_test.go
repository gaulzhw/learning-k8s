package indexer

import (
	"strings"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func TestIndexer(t *testing.T) {
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{
		"byUser": func(obj interface{}) ([]string, error) {
			pod := obj.(*corev1.Pod)
			userString := pod.Annotations["users"]
			return strings.Split(userString, ","), nil
		},
	})

	index.Add(&corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "one", Annotations: map[string]string{"users": "annie,hoob"}}})
	index.Add(&corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "two", Annotations: map[string]string{"users": "hoob,onime"}}})
	index.Add(&corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "three", Annotations: map[string]string{"users": "onime,jonny"}}})
}
