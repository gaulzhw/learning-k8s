# vcluster

https://github.com/loft-sh/vcluster



## host cluster

```shell
kind create cluster
```



## requirements

- kubectl

- helm

- vcluster

  ```shell
  curl -s -L "https://github.com/loft-sh/vcluster/releases/latest" | sed -nE 's!.*"([^"]*vcluster-darwin-amd64)".*!https://github.com\1!p' | xargs -n 1 curl -L -o vcluster
  ```



## deploy vclusters

创建一个 virtual cluster `vcluster-1`，指定 namespace 为 `host-namespace-1`

```shell
vcluster create vcluster-1 -n host-namespace-1 --distro k8s
```

在创建出来的资源上，etcd 的 statefulset 使用 pvc 绑定存储，可以去掉 pvc 测试



## use

### connect to vcluster

- 终端1，vcluster-1 暴露 apiserver

  ```shell
  vcluster connect vcluster-1 -n host-namespace-1
  ```

- 终端2，连接 vcluster-1

  ```shell
  export KUBECONFIG=./kubeconfig.yaml
  
  kubectl get namespace
  
  kubectl create namespace demo-nginx
  kubectl create deployment nginx-deployment -n demo-nginx --image=nginx
  
  kubectl get pods -n demo-nginx
  ```

- 终端3，连接宿主集群

  ```shell
  export KUBECONFIG=
  
  kubectl get namespaces
  kubectl get deployments -n host-namespace-1
  kubectl get pods -n host-namespace-1
  ```
