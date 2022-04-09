# kubeconfig



## 管理 kubeconfig

### RBAC

```shell
# 1. 创建 namespace
kubectl create ns well

# 2. 创建 ServiceAccount
apiVersion: v1
kind: ServiceAccount
metadata:
  name: well-sa
  namespace: well

# 3. 创建 Role
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: well-role
  namespace: well
rules:
- apiGroups: [""]
  resources: 
  - pods
  - deployments
  - configmaps
  - services
  verbs: 
  - get
  - list
  - watch
  - create
  - update
  - delete

# 4. 创建 RoleBinding
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: well-binding
  namespace: well
subjects:
- kind: ServiceAccount
  name: well-sa
roleRef:
  kind: Role
  name: well-role
  apiGroup: rbac.authorization.k8s.io
```



### kubeconfig

```yaml
apiVersion: v1
kind: Config
users:
- name: well
  user:
    token: <token>
clusters:
- cluster:
    certificate-authority-data: <certificate-authority-data>
    server: <api-server>
  name: well-cluster
contexts:
- context:
    cluster: well-cluster
    namespace: well
    user: well
  name: well-cluster
current-context: well-cluster
```

- 通过命令 `kubectl config view --flatten --minify` 可以拿到 certificate-authority-data 和 api-server 信息 。
- 通过命令 `kubectl describe sa well-sa -n well` 拿到 secret 的 key。
- 通过命令 `kubectl describe secret <key> -n well`  拿到 token 信息。



## Reference

https://cloud.tencent.com/developer/article/1847995
