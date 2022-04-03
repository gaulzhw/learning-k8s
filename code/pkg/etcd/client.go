package etcd

import (
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func NewEtcdClient() (*clientv3.Client, error) {
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}

	log.Println("success")
	return client, nil
}
