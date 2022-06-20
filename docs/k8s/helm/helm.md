# helm



## 仓库管理

```shell
# add
helm repo add [repo-name] [repo-url]

# remove
helm repo remove [repo-name]

# list
helm repo list

# update
helm repo update

# inspect chart
helm inspect chart [chart-name]
```



## 部署应用

```shell
# search
helm search repo [repo-name]

# install
helm install [name] [chart]

# template
helm template [chart]

# uninstall
helm uninstall [name]
```

