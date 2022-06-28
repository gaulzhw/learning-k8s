# Karmada



## install with kind

### 安装宿主集群

```shell
kind create cluster --name=hub --config=-<<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  image: kindest/node:v1.23.4
  extraPortMappings:
  - hostPort: 5443
    containerPort: 32443
EOF
```



### 容器内执行

`docker exec -it hub-control-plane bash`



1. install kubectl-karmada
2. install karmada cluster

```shell
# install kubectl-karmada
curl -sLo "kubectl-karmada-linux-amd64.tgz" "https://github.com/karmada-io/karmada/releases/download/v1.1.1/kubectl-karmada-linux-amd64.tgz"

tar -zxvf kubectl-karmada-linux-amd64.tgz
mv kubectl-karmada /usr/local/bin/kubectl-karmada
rm -rf LICENSE kubectl-karmada-linux-amd64.tgz

# install karmada cluster
kubectl karmada init
```



将 Karmada 集群的 kubeconfig 拷贝到宿主机

/etc/karmada/karmada-apiserver.config
