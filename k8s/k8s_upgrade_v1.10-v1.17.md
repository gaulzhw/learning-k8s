# 集群升级



## 背景

集群版本升级，v1.10->v1.17



## 升级难点

升级过程中不能因为集群组件自身逻辑变化导致业务重启

- 当前运行版本较低，但是运行容器数量却很多，其中部分仍然是单副本运行，为了不影响业务运行，需要尽可能避免容器重启，这无疑是升级中最大的难点，而在 v1.10 版本和 v1.17 版本之间，kubelet 关于容器 Hash 值计算方式发生了变化，也就是说一旦升级必然会触发 kubelet 重新启动容器。

- 社区推荐的方式是基于偏差策略的升级以保证高可用集群升级同时不会因为 API resources 版本差异导致 kube-apiserve 和 kubelet 等组件出现兼容性错误，这就要求每次升级组件版本不能有 2 个 Final Release 以上的偏差，比如直接从 v1.11 升级至 v1.13 是不推荐的。

- 升级过程中由于新特性的引入，API 兼容性可能引发旧版本集群的配置不生效，为整个集群埋下稳定性隐患。要求在升级前尽可能的熟悉升级版本间的 ChangeLog，排查出可能带来潜在隐患的新特性。



## 升级方案

主流的应用升级方式有两种，分别是原地升级和替换升级。

替换升级：

- Kubernetes 替换升级是先准备一个高版本集群，对低版本集群通过逐个节点排干、删除最后加入新集群的方式将低版本集群内节点逐步轮换升级到新版本。
- 替换升级的优点是原子性更强，逐步升级各个节点，升级过程不存在中间态，对业务安全更有保障；缺点是集群升级工作量较大，排干操作对 Pod 重启敏感度高的应用、有状态应用、单副本应用等都不友好。

原地升级：

- Kubernetes原地升级是对节点上服务如kube-controller-manager、kubelet等组件按照一定顺序批量更新，从节点角色维度批量管理组件版本。
- 原地升级的优点是自动化操作便捷，并且通过适当的修改能够很好的保证容器的生命周期连续性；缺点是集群升级中组件升级顺序很重要，升级中存在中间态，并且一个组件重启失败可能影响后续其他组件升级，原子性差。



## 原地升级的困难

考虑业务对重启的容忍度，来权衡选择升级方案。



### 跨版本升级

社区在设计API时遵循向上或向下兼容的原则，也遵循社区的偏差策略。

API groups弃用、启用时，对于alpha版本会立即生效，对于beta版本将会继续支持3个版本，超过对应版本将导致API resource version不兼容。

https://kubernetes.io/releases/version-skew-policy/

对于各个版本升级的影响，可以分析change log，梳理和充分测试，确认升级版本之间是否存在影响业务运行和集群管理操作的API兼容性问题。对于API类型的废弃，可以通过配置apiserver中相应参数来启动继续使用，保证环境业务继续正常运行。



### 避免容器重启

版本升级会导致大量容器被重建，重启原因从升级后的kubelet组件日志看到"Container definition changed"，报错位于pkg/kubelet/kuberuntime_manage.go文件computePodActions方法，该方法用来计算pod的spec哈希值是否发生变化，如果变化则返回true，告知kubelet syncPod方法触发pod内容器重建或者pod重建。

v1.10、v1.17版本在计算容器hash时使用的数据不同，而且高版本kubelet中对容器的结构也增加了新属性，通过go-spew库计算出结果自然不一致，进一步向上传递使得syncPod方法触发容器重建。



https://github.com/cloudusers/UpdateKubeletVersionIgnoreContainerRestart

本地创建一个记录旧集群版本信息和启动时间的配置文件，kubelet代码中维护一个cache读取配置文件，在每个syncPod周期中，当kubelet发现自身version高于cache中记录的oldVersion，并且容器启动时间早于当前kubelet启动时间，则会跳过容器hash值计算。升级后的集群内运行定时任务探测pod的containerSpec是否与高版本计算方式计算得到的hash结果全部一致，如果是则可以删除掉本地配置文件，syncPod逻辑恢复到与社区一致。



### pod非预期驱逐问题

v1.13版本引入的TaintBasedEvictions特性用于更细粒度的管理 Pod 的驱逐条件。

在v1.13基于条件版本之前，驱逐是基于NodeController的统一时间驱逐，节点NotReady超过默认5分钟后，节点上的Pod才会被驱逐

在v1.16默认开启TaintBasedEvictions后，节点NotReady的驱逐将会根据每个Pod自身配置的TolerationSeconds来差异化的处理。

旧版本集群创建的Pod默认没有设置TolerationSeconds，一旦升级完毕TaintBasedEvictions被开启，节点变成NotReady后5秒就会驱逐节点上的Pod。对于短暂的网络波动、kubelet重启等情况都会影响集群中业务的稳定性。

TaintBasedEvictions对应的控制器是按照Pod定义中的tolerationSeconds决定Pod的驱逐时间，也就是说只要正确设置Pod中的tolerationSeconds就可以避免出现Pod的非预期驱逐。

```yaml
tolerations:
- effect: NoExecute
  key: node.kubernetes.io/not-ready
  operator: Exists
  tolerationSeconds: 300
- effect: NoExecute
  key: node.kubernetes.io/unreachable
  operator: Exists
  tolerationSeconds: 300
```

v1.16默认开启DefaultTolerationSeconds准入控制器基于k8s-apiserver输入参数default-not-ready-toleration-seconds和default-unreachable-toleration-seconds为Pod设置默认的容忍度，以容忍notready:NoExecute和unreachable:NoExecute污点。新建Pod在请求发送后会经过DefaultTolerationSeconds准入控制器给pod加上默认的tolerations。

但是这个逻辑如何对集群中已经创建的Pod生效呢？准入控制器除了支持create操作，update操作也会更新pod定义触发DefaultTolerationSeconds插件去设置tolerations。

通过给集群中已经运行的Pod打label就可以达成目的。



### pod MatchNodeSelector

写一个脚本去实时同步节点上非Running状态的pod和发生重启的容器，来判断升级时pod是否发生非预期的驱逐以及是否存在pod内容器批量重启。

发现有pod被标记为MatchNodeSelector状态，该节点上业务容器停止了。



kubelet中的错误日志

```
predicate.go:132] Predicate failed on Pod: nginx-7dd9db975d-j578s_default(e3b79017-0b15-11ec-9cd4-000c29c4fa15), for reason: Predicate MatchNodeSelector failed
kubelet_pods.go:1125] Killing unwanted pod "nginx-7dd9db975d-j578s"
```

Pod变成MatchNodeSelector状态是因为kubelet重启时对节点上Pod做准入检查时无法找到节点满足要求的节点标签，pod状态就会被设置为Failed状态，而Reason被设置为MatchNodeSelector。在kubectl命令获取时，printer做了相应转换直接显示了Reason，因此看到Pod状态是MatchNodeSelector。

通过给节点加上标签，可以让Pod重新调度回来，然后删除掉MatchNodeSelector状态的Pod即可。



升级前写脚本检查节点上Pod定义中使用的NodeSelector属性节点是否都有对应的Label。



### 无法访问apiserver

集群升级后，突然有节点变成NotReady，通过重启kubelet节点恢复。

https://github.com/kubernetes/kubernetes/issues/87615#issuecomment-803517109

分析出错原因发现kubelet日志中出现了大量use of closed network connection报错，问题起因是kubelet默认连接采用HTTP/2.0长连接，在构建client到server的连接时使用的golang net/http2包中存在bug (https://github.com/golang/go/issues/34978)，在http连接池中仍然能获取到broken的连接，导致kubelet无法正常与apiserver通信。

Golang通过增加http2连接健康检查规避这个问题，Golang v1.15.11版本彻底修复。

https://github.com/kubernetes/kubernetes/pull/100376



### TCP连接数问题

集群每个节点kubelet都有近10个长连接与apiserver通信，v1.10版本只有1个长连接。这种TCP连接数增加无疑会对LB造成压力，随着节点增多，一旦LB被拖垮，kubelet无法上报心跳，节点会变成NotReady，紧接着将会有大量pod被驱逐。



https://github.com/kubernetes/kubernetes/pull/95427

增加了判断逻辑导致kubelet获取client时不再从cache中获取缓存的长连接。

transport的主要功能其实就是缓存长连接，用于大量http请求场景下的连接服用，减少发送请求时TCP (TLS)连接建立的时间损耗。

```go
// 为 clientConfig 设置 Dial属性,因此 kubelet 构建 clinet 时会新建 transport
func updateDialer(clientConfig *restclient.Config) (func(), error) {
    if clientConfig.Transport != nil || clientConfig.Dial != nil {
        return nil, fmt.Errorf("there is already a transport or dialer configured")
    }
    d := connrotation.NewDialer((&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext)
    clientConfig.Dial = d.DialContext
    return d.CloseAll, nil
}
```

该PR中对transport自定义RoundTripper的接口，一旦tlsConfig对象中有Dial或者Proxy属性，则不使用cache中的连接而新建连接。



https://github.com/kubernetes/kubernetes/pull/105490

通过重构client-go的接口实现对自定义RESTClient的TCP连接复用。



## 无损升级操作

跨版本升级最大的风险是升级前后对象定义不一致，可能导致升级后的组件无法解析保存在etcd数据库中的对象；也可能是升级存在中间态，kubelet还未升级而控制平面组件升级，存在上报状态异常，最坏的情况是节点上pod被驱逐。需要在升级前考虑并测试验证。



为了保证升级时能及时处理未覆盖到的特殊情况，需要备份etcd数据库，并在升级期间停止控制器和调度器，避免非预期的控制逻辑发生。实际应该是停止controller manager中的部分控制器，不过需要修改代码编译临时controller manager，增加了升级流程步骤和管理复杂度，可以直接停掉全局控制器简化。



检查新老版本服务的配置项区别。



升级步骤

- 备份集群（二进制，配置文件，etcd 数据库等）
- 灰度升级部分节点，验证二进制和配置文件正确性
- 提前分发升级的二进制文件
- 停止控制器、调度器和告警
- 更新控制平面服务配置文件，升级组件
- 更新计算节点服务配置文件，升级组件
- 为节点打 Label 触发 Pod 增加 tolerations 属性
- 打开控制器和调度器，启用告警
- 集群业务点检，确认集群正常



升级过程中建议节点并发数不要太高，因为大量节点 kubelet 同时重启上报信息，对 kube-apiserver 前面使用的 LB 带来冲击，特别情况下可能节点心跳上报失败，节点状态会在 NotReady 与 Ready 状态间跳动。



## References

https://mp.weixin.qq.com/s/fAp9bq7hhFDjn8MvP5mNgw