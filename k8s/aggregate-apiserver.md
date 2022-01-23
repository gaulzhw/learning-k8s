# aggregate-apiserver



## Demo

1. download apiserver-build for MAC，https://github.com/kubernetes-sigs/apiserver-builder-alpha/releases/download/v1.18.0/apiserver-builder-alpha-v1.18.0-darwin-amd64.tar.gz

2. 初始化工程，apiserver-boot init repo --domain demo.com
3. 创建api资源，apiserver-boot create group version resource --group demo --version v1alpha1 --kind Bird
4. 本地运行apiserver和controller，apiserver-boot run local



## 相关代码

- from apiserver-builder to kube-builder，https://github.com/kubernetes-sigs/apiserver-builder-alpha

- from controller-runtime to apiserver-runtime，https://github.com/kubernetes-sigs/apiserver-runtime



## References

https://pkg.go.dev/sigs.k8s.io/apiserver-runtime/pkg/builder

https://github.com/duyanghao/kubernetes-reading-notes/blob/master/core/api-server/extension/aggregated-apiserver.md

https://hackerain.me/2020/10/08/kubernetes/kube-apiserver-extensions.html

https://developer.aliyun.com/article/610357