# clusternet



# 环境

本地 kind 安装集群测试

https://github.com/clusternet/clusternet/blob/main/docs/tutorials/installing-clusternet-the-hard-way.md



# 部署

https://github.com/clusternet/clusternet/blob/main/docs/tutorials/deploying-applications-to-multiple-clusters.md

kubectl clusternet apply -f nginx-deployment.yaml

kubectl clusternet apply -f nginx-svc.yaml

kubectl clusternet apply -f subscription.yaml

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: foo

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
  namespace: foo
  labels:
    clusternet-app: multi-cluster-nginx
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80

---

apiVersion: v1
kind: Service
metadata:
  name: my-nginx-svc
  namespace: foo
  labels:
    app: nginx
spec:
  type: ClusterIP
  ports:
  - port: 80
    selector:
      app: nginx

---

### subscription
apiVersion: apps.clusternet.io/v1alpha1
kind: Subscription
metadata:
  name: app-demo
  namespace: default
spec:
  subscribers: # defines the clusters to be distributed to
  - clusterAffinity:
    matchLabels:
      clusters.clusternet.io/cluster-id: 7fa03738-5ae9-4805-b349-2d2119d15584
    feeds: # defines all the resources to be deployed with
    - apiVersion: v1
      kind: Namespace
      name: foo
    - apiVersion: v1
      kind: Service
      name: my-nginx-svc
      namespace: foo
    - apiVersion: apps/v1
      kind: Deployment
      name: my-nginx
      namespace: foo
```



# 结果

## 管理集群

部署的Namespace、Deployment、Service资源，都是通过Shadow转成Manifest的CRD



## 被管集群

Namespace、Deployment、Service均生成了