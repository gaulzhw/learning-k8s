# OpenKruise



## install

### cluster with kind

```shell
kind create cluster --config=-<<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
- role: worker
EOF
```

### openkruise

```shell
helm repo add openkruise https://openkruise.github.io/charts/

# [Optional]
helm repo update

# Install the latest version.
helm install kruise openkruise/kruise --version 1.1.0
```

### uninstall

```shell
helm uninstall kruise
```



## workload

### CloneSet

```yaml
apiVersion: apps.kruise.io/v1alpha1
kind: CloneSet
metadata:
  name: sample
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sample
  template:
    metadata:
      labels:
        app: sample
    spec:
      containers:
      - name: nginx
        image: nginx:alpine
```



### Advanced StatefulSet



### Advanced DaemonSet



### BroadcastJob



### AdvancedCronJob



### SidecarSet



### WorkloadSpread



### UnitedDeployment



### ContainerRestart



### ImagePullJob



### ContainerLaunchPriority



### ResourceDistribution



### DeletionProtection



### PodUnavailableBudget
