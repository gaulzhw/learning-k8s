# OCM



## how to install

### 本地部署集群

hub 集群

```shell
# 部署hub集群，kind名叫hub
kind create cluster --name hub --image kindest/node:v1.20.7
```

worker 集群

```shell
# 部署worker集群，kind名叫cluster1
kind create cluster --name cluster1 --image kindest/node:v1.20.7
```



### 安装 clusteradm

hub、cluster1 两个 kind 容器内安装 clusteradm

```shell
docker exec -it hub-control-plane bash -c "curl -L -o clusteradm_linux_amd64.tar.gz https://github.com/open-cluster-management-io/clusteradm/releases/download/v0.1.0-alpha.7/clusteradm_linux_amd64.tar.gz && tar -zxvf clusteradm_linux_amd64.tar.gz"

docker exec -it cluster1-control-plane bash -c "curl -L -o clusteradm_linux_amd64.tar.gz https://github.com/open-cluster-management-io/clusteradm/releases/download/v0.1.0-alpha.7/clusteradm_linux_amd64.tar.gz && tar -zxvf clusteradm_linux_amd64.tar.gz"
```



### hub 集群部署

在 hub 集群的 kind 容器内初始化 cluster

```shell
docker exec -it hub-control-plane bash -c "./clusteradm init"
```

执行完会输出 clusteradm join ... 信息，记录下来用来给托管集群 join

hub-token 对应 open-cluster-management namespace 下的 serviceaccount/cluster-bootstrap 映射的 secret

等待 open-cluster-management、open-cluster-management-hub 两个 namespace 下的 pod 启动



### worker 集群部署

在 cluster1 集群的 kind 容器内将集群 join 到 hub 集群

```shell
docker exec -it cluster1-control-plane bash -c "./clusteradm join --hub-token <token_data> --hub-apiserver <hub-api> --cluster-name cluster1"
```

执行完会提示 Waiting for the management components to become ready...

cluster1 集群会多 open-cluster-management、open-cluster-management-agent、open-cluster-management-agent-addon 三个namespace，等待其中的 pod 运行

到 hub 集群查看 csr， `kubectl get csr --context=kind-hub` 会发现多了一个 cluster1 集群的 csr 请求，状态是 Pending

在 hub 集群的 kind 容器内执行命令： `docker exec -it hub-control-plane bash -c "./clusteradm accept --clusters cluster1"`

完成之后，hub 集群会有一个 cluster1 的 namespace，可以通过 `kubectl get managedcluster --context=kind-hub` 查看创建成功的托管集群



## test

### 部署 workload 测试

在 hub 集群内创建 ManifestWork 资源，指定 namespace 是 cluster1，表示在 cluster1 上创建 workload

```yaml
# kubectl apply -f xxx.yaml --context=kind-hub
apiVersion: work.open-cluster-management.io/v1
kind: ManifestWork
metadata:
  name: mw-01
  namespace: cluster1
spec:
  workload:
    manifests:
    - apiVersion: v1
      kind: Pod
      metadata:
        name: hello
        namespace: default
      spec:
        containers:
        - name: hello
          image: busybox
          command: ["sh", "-c", 'echo "Hello, Kubernetes!" && sleep 3600']
```