package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateWithPrefix(t *testing.T) {
	client := newFakeClientSet()
	assert.NotNil(t, client)

	cm, err := client.CoreV1().ConfigMaps("kube-system").Create(context.TODO(), &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "test-",
		},
		Data: map[string]string{
			"test": "test",
		},
	}, metav1.CreateOptions{})
	assert.NoError(t, err)
	t.Logf("%+v", cm)
}
