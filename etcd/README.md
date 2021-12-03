# ETCD

注：不特别强调，都是使用ETCDCTL_API=3

etcd 3.4+默认使用API=3



## 操作容器的etcd

```shell
kubectl exec -it <etcd-pod> -n kube-system -- \
	etcdctl \
  --cacert=/etc/kubernetes/pki/etcd/ca.crt \
  --cert=/etc/kubernetes/pki/etcd/server.crt \
  --key=/etc/kubernetes/pki/etcd/server.key \
  --endpoints=https://127.0.0.1:2379 \
  member list
```



## 实验

创建非安全集群，便于调试

```shell
docker run -d -p 2379:2379 -p 2380:2380 --name=etcd k8s.gcr.io/etcd:3.5.0-0 -- \
	etcd \
	--name infra0 \
	--initial-advertise-peer-urls http://0.0.0.0:2380 \
	--listen-peer-urls http://0.0.0.0:2380 \
	--listen-client-urls http://0.0.0.0:2379 \
	--advertise-client-urls http://0.0.0.0:2379 \
	--initial-cluster-token etcd-cluster-1 \
	--initial-cluster infra0=http://0.0.0.0:2380 \
	--initial-cluster-state new
```



## 常见运维场景、操作

- 数据备份、同步

```shell
etcdctl snapshot save BackupFile.db
etcdctl snapshot restore BackupFile.db -data-dir ETCDDir
```

- 碎片整理
- 存储扩容
- 节点变更
  - 节点迁移、替换
  - 节点增加
  - 节点移除
  - 强制性重启集群



## 优化方案



## trouble shooting