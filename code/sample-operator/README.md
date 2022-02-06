# sample-operator



## build with kubebuilder

```shell
kubebuilder init --domain example.io --repo example.io/test

kubebuilder create api --group apps --version v1alpha1 --kind Application
```



## after editing API definition

```shell
make manifests
```

