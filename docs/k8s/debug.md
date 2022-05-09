# k8s debug



## apiserver

- etcd

  ```shell
  docker run -d --name=etcd \
    -p 2379:2379 -p 2380:2380 \
    k8s.gcr.io/etcd:3.5.1-0 \
    etcd \
    --name etcd0 \
    --initial-advertise-peer-urls http://0.0.0.0:2380 \
    --listen-peer-urls http://0.0.0.0:2380 \
    --listen-client-urls http://0.0.0.0:2379 \
    --advertise-client-urls http://0.0.0.0:2379 \
    --initial-cluster-token etcd-cluster \
    --initial-cluster etcd0=http://0.0.0.0:2380 \
    --initial-cluster-state new
  ```

- 证书

  ```shell
  kind create cluster --name=host
  docker cp host-control-plane:/etc/kubernetes ~/kubernetes
  ```

- 启动

  ```shell
  kube-apiserver
  --etcd-servers=http://127.0.0.1:2379
  --secure-port=6443
  --client-ca-file=/etc/kubernetes/pki/ca.crt
  --tls-cert-file=/etc/kubernetes/pki/apiserver.crt
  --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
  --service-account-signing-key-file=/etc/kubernetes/pki/sa.key
  --service-account-key-file=/etc/kubernetes/pki/sa.pub
  --service-account-issuer=https://kubernetes.default.svc.cluster.local
  ```



## controller-manager

kubeconfig 指向 admin.conf



## scheduler

kubeconfig 指向 admin.conf