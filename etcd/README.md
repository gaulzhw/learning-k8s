# ETCD



## 操作容器的etcd

```shell
kubectl exec -it <etcd-pod> -n kube-system -- etcdctl \
  --cacert=/etc/kubernetes/pki/etcd/ca.crt \
  --cert=/etc/kubernetes/pki/etcd/server.crt \
  --key=/etc/kubernetes/pki/etcd/server.key \
  --endpoints=https://127.0.0.1:2379 \
  member list
```



## 常见运维场景、操作