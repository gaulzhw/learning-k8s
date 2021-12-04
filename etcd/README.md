# ETCD

注：不特别强调，都是使用ETCDCTL_API=3

etcd 3.4+默认使用API=3



创建非安全集群，便于调试

```shell
docker run -d -p 2379:2379 -p 2380:2380 --name=etcd k8s.gcr.io/etcd:3.5.0-0 \
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



## etcd命令

- 查看节点状态，etcdctl endpoint status [-w=json 显示更详细的信息]
- 查看member信息，etcdctl member list
- member的操作，etcdctl member add/remove
- 手动变更leader，etcdctl move-leader



## 常见运维场景、操作

- 集群搭建

  ```shell
  docker run -d --network=host --name=etcd1 k8s.gcr.io/etcd:3.5.0-0 \
    etcd \
    --name infra1 \
    --initial-advertise-peer-urls http://127.0.0.1:12380 \
    --listen-peer-urls http://127.0.0.1:12380 \
    --listen-client-urls http://127.0.0.1:12379 \
    --advertise-client-urls http://127.0.0.1:12379 \
    --initial-cluster-token etcd-cluster-1 \
    --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 \
    --initial-cluster-state new
  
  docker run -d --network=host --name=etcd1 k8s.gcr.io/etcd:3.5.0-0 \
    etcd \
    --name infra2 \
    --initial-advertise-peer-urls http://127.0.0.1:12380 \
    --listen-peer-urls http://127.0.0.1:12380 \
    --listen-client-urls http://127.0.0.1:12379 \
    --advertise-client-urls http://127.0.0.1:12379 \
    --initial-cluster-token etcd-cluster-1 \
    --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 \
    --initial-cluster-state new
  
  docker run -d --network=host --name=etcd1 k8s.gcr.io/etcd:3.5.0-0 \
    etcd \
    --name infra3 \
    --initial-advertise-peer-urls http://127.0.0.1:12380 \
    --listen-peer-urls http://127.0.0.1:12380 \
    --listen-client-urls http://127.0.0.1:12379 \
    --advertise-client-urls http://127.0.0.1:12379 \
    --initial-cluster-token etcd-cluster-1 \
    --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380 \
    --initial-cluster-state new
  ```

- 数据备份、同步

  ```shell
  etcdctl snapshot save BackupFile.db
  etcdctl snapshot restore BackupFile.db -data-dir ETCDDir
  ```

- 碎片整理

  ```shell
  # 显示配额
  etcdctl -w=table endpoint status
  # 查看警告
  etcdctl alarm list
  
  # 获取etcd数据的revision
  etcdctl endpoint status -w=json | egrep -o '"revision":[0-9]*' | egrep -o '[0-9].*'
  # 压缩旧版本数据
  etcdctl compact $rev
  # 执行碎片整理
  etcdctl defrag
  
  # 解除告警
  etcdctl alarm disarm
  ```

- 存储扩容

- 节点变更

  - 节点迁移、替换

    操作同节点增加，只是需要同步数据文件夹，同时修改member的url `etcdctl member update [member-id] --peer-urls=xxx`

  - 节点增加

    ```shell
    etcdctl member add infra4 --peer-urls=http://127.0.0.1:42380
    
    docker run -d --network=host --name=etcd1 k8s.gcr.io/etcd:3.5.0-0 \
      etcd \
      --name infra4 \
      --initial-advertise-peer-urls http://127.0.0.1:42380 \
      --listen-peer-urls http://127.0.0.1:42380 \
      --listen-client-urls http://127.0.0.1:42379 \
      --advertise-client-urls http://127.0.0.1:42379 \
      --initial-cluster-token etcd-cluster-1 \
      --initial-cluster infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380,infra4=http://127.0.0.1:42380 \
      --initial-cluster-state existing
    ```

  - 节点移除

    ```shell
    etcdctl member list
    etcdctl member remove [member-id]
    
    docker rm -f etcd4 # 容器会结束
    ```

  - 强制性重启集群

    当集群超过半数的节点都失效时，就需要通过手动的方式，强制性让某个节点以自己为Leader，利用原有数据启动一个新集群

    创建完成后需要检查member list，确认删除memeber还是修改member ip

    ```shell
    docker run -d --network=host --name=etcd1 k8s.gcr.io/etcd:3.5.0-0 \
      etcd \
      --name infra1 \
      --initial-advertise-peer-urls http://127.0.0.1:12380 \
      --listen-peer-urls http://127.0.0.1:12380 \
      --listen-client-urls http://127.0.0.1:12379 \
      --advertise-client-urls http://127.0.0.1:12379 \
      --initial-cluster-token etcd-cluster-1 \
      --initial-cluster infra1=http://127.0.0.1:12380 \
      --initial-cluster-state new
      --force-new-cluster # 强制启动
    ```



## 优化方案



## trouble shooting