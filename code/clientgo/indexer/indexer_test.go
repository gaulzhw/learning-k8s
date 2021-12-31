package indexer

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIndexer(t *testing.T) {
	index := newIndexer()
	index.Add(&corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "one", Annotations: map[string]string{"users": "annie,hoob"}}})
	index.Add(&corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "two", Annotations: map[string]string{"users": "hoob,onime"}}})
	index.Add(&corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "three", Annotations: map[string]string{"users": "onime,jonny"}}})
}
