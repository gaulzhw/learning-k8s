# k8s debug



## apiserver

- etcd

  ```shell
  docker run -d --name=etcd \
    -p 2379:2379 -p 2380:2380 \
    k8s.gcr.io/etcd:3.5.1-0 \
    etcd \
    --name etcd0 \
    --initial-advertise-peer-urls http://0.0.0.0:2380 \
    --listen-peer-urls http://0.0.0.0:2380 \
    --listen-client-urls http://0.0.0.0:2379 \
    --advertise-client-urls http://0.0.0.0:2379 \
    --initial-cluster-token etcd-cluster \
    --initial-cluster etcd0=http://0.0.0.0:2380 \
    --initial-cluster-state new
  ```

- 证书

  ```shell
  kind create cluster --name=host
  docker cp host-control-plane:/etc/kubernetes ~/kubernetes
  ```

- 启动

  ```shell
  kube-apiserver
  --etcd-servers=http://127.0.0.1:2379
  --secure-port=6443
  --client-ca-file=/etc/kubernetes/pki/ca.crt
  --tls-cert-file=/etc/kubernetes/pki/apiserver.crt
  --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
  --service-account-signing-key-file=/etc/kubernetes/pki/sa.key
  --service-account-key-file=/etc/kubernetes/pki/sa.pub
  --service-account-issuer=https://kubernetes.default.svc.cluster.local
  ```



## controller-manager

kubeconfig 指向 admin.conf



## scheduler

kubeconfig 指向 admin.conf

```yaml
apiVersion: kubescheduler.config.k8s.io/v1beta1
kind: KubeSchedulerConfiguration
clientConnection:
  acceptContentTypes: ""
  burst: 100
  contentType: application/vnd.kubernetes.protobuf
  kubeconfig: ~/.kube/config
  qps: 50
enableContentionProfiling: true
enableProfiling: true
healthzBindAddress: ""
leaderElection:
  leaderElect: true
  leaseDuration: 15s
  renewDeadline: 10s
  resourceLock: leases
  resourceName: kube-scheduler
  resourceNamespace: kube-system
  retryPeriod: 2s
metricsBindAddress: ""
parallelism: 16
percentageOfNodesToScore: 0
podInitialBackoffSeconds: 1
podMaxBackoffSeconds: 10
profiles:
- schedulerName: default-scheduler
  plugins:
    bind:
      enabled:
      - name: DefaultBinder
        weight: 0
      disabled:
      - name: '*'
    filter:
      enabled:
      - name: NodeUnschedulable
        weight: 0
      - name: NodeName
        weight: 0
      - name: TaintToleration
        weight: 0
      - name: NodeAffinity
        weight: 0
      - name: NodePorts
        weight: 0
      - name: NodeResourcesFit
        weight: 0
      - name: VolumeRestrictions
        weight: 0
      - name: EBSLimits
        weight: 0
      - name: GCEPDLimits
        weight: 0
      - name: NodeVolumeLimits
        weight: 0
      - name: AzureDiskLimits
        weight: 0
      - name: VolumeBinding
        weight: 0
      - name: VolumeZone
        weight: 0
      - name: PodTopologySpread
        weight: 0
      - name: InterPodAffinity
        weight: 0
      disabled:
      - name: '*'
    permit: {}
    postBind: {}
    postFilter:
      enabled:
      - name: DefaultPreemption
        weight: 0
      disabled:
      - name: '*'
    preBind:
      enabled:
      - name: VolumeBinding
        weight: 0
      disabled:
      - name: '*'
    preFilter:
      enabled:
      - name: NodeResourcesFit
        weight: 0
      - name: NodePorts
        weight: 0
      - name: PodTopologySpread
        weight: 0
      - name: InterPodAffinity
        weight: 0
      - name: VolumeBinding
        weight: 0
      disabled:
      - name: '*'
    preScore:
      enabled:
      - name: InterPodAffinity
        weight: 0
      - name: PodTopologySpread
        weight: 0
      - name: TaintToleration
        weight: 0
      disabled:
      - name: '*'
    queueSort:
      enabled:
      - name: PrioritySort
        weight: 0
      disabled:
      - name: '*'
    reserve:
      enabled:
      - name: VolumeBinding
        weight: 0
      disabled:
      - name: '*'
    score:
      enabled:
      - name: NodeResourcesBalancedAllocation
        weight: 1
      - name: ImageLocality
        weight: 1
      - name: InterPodAffinity
        weight: 1
      - name: NodeResourcesLeastAllocated
        weight: 1
      - name: NodeAffinity
        weight: 1
      - name: NodePreferAvoidPods
        weight: 10000
      - name: PodTopologySpread
        weight: 2
      - name: TaintToleration
        weight: 1
      disabled:
      - name: '*'
  pluginConfig:
  - args:
      apiVersion: kubescheduler.config.k8s.io/v1beta1
      kind: DefaultPreemptionArgs
      minCandidateNodesAbsolute: 100
      minCandidateNodesPercentage: 10
    name: DefaultPreemption
  - args:
      apiVersion: kubescheduler.config.k8s.io/v1beta1
      hardPodAffinityWeight: 1
      kind: InterPodAffinityArgs
    name: InterPodAffinity
  - args:
      apiVersion: kubescheduler.config.k8s.io/v1beta1
      kind: NodeAffinityArgs
    name: NodeAffinity
  - args:
      apiVersion: kubescheduler.config.k8s.io/v1beta1
      kind: NodeResourcesFitArgs
    name: NodeResourcesFit
  - args:
      apiVersion: kubescheduler.config.k8s.io/v1beta1
      kind: NodeResourcesLeastAllocatedArgs
      resources:
      - name: cpu
        weight: 1
      - name: memory
        weight: 1
    name: NodeResourcesLeastAllocated
  - args:
      apiVersion: kubescheduler.config.k8s.io/v1beta1
      defaultingType: System
      kind: PodTopologySpreadArgs
    name: PodTopologySpread
  - args:
      apiVersion: kubescheduler.config.k8s.io/v1beta1
      bindTimeoutSeconds: 600
      kind: VolumeBindingArgs
    name: VolumeBinding
```