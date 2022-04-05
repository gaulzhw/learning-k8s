package etcd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEtcdClient(t *testing.T) {
	client, err := NewEtcdClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
