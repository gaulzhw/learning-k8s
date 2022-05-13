package etcd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcdClient(t *testing.T) {
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(config)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
