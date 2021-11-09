# istio



## k8s cluster

```shell
kind create cluster --config=-<<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  image: kindest/node:v1.20.7
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
    EOF
```



## install istio in k8s

```shell
# download istio
curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.11.4 sh -

# generate yaml
./bin/istioctl manifest generate --set profile=default
```

调整部署的yaml文件 istio-1.11.4-default.yaml，发布到集群



## test

```yaml
# Gateway
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: test-gateway
spec:
  selector:
    istio: ingressgateway
  servers:
  - hosts:
    - test.com
    port:
      number: 80
      name: http
      protocol: HTTP
---
# VirtualService
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: test-vs
spec:
  gateways:
  - test-gateway
  hosts:
  - test.com
  http:
  - route:
    - destination:
        host: nginx-svc
        port:
          number: 80
```



## 参考文档

- https://blog.51cto.com/u_14625168/2474277
