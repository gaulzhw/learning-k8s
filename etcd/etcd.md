# etcd



## etcd 概念

### 术语

| 术语      | 描述                                    | 备注                                                     |
| --------- | --------------------------------------- | -------------------------------------------------------- |
| Raft      | Raft 算法，etcd 实现一致性的核心        | etcd 有 etcd-raft 模块                                   |
| Follower  | Raft 中的从属节点                       | 竞争 Leader 失败                                         |
| Leader    | Raft 中的领导协调节点，用于处理数据提交 | Leader 节点协调整个集群                                  |
| Candidate | 候选节点                                | 当 Follower 接收 Leader 节点的消息超时会转变为 Candidate |
| Node      | Raft 状态机的实例                       | Raft 中涉及多个节点                                      |
| Member    | etcd 实例，管理着对应的 Node 节点       | 可处理客户端请求                                         |
| Peer      | 同一个集群中的另一个 Member             | 其他成员                                                 |
| Cluster   | etcd 集群                               | 拥有多个 etcd Member                                     |
| Lease     | 租期                                    | 关键设置的租期，过期删除                                 |
| Watch     | 监测机制                                | 监控键值对的变化                                         |
| Term      | 任期                                    | 某个节点成为 Leader，到下一个竞选的时间                  |
| WAL       | 预写式日志                              | 用于持久化存储的日志格式                                 |
| Client    | 客户端                                  | 向 etcd 发起请求的客户端                                 |



### 应用场景

- 键值对存储
- 服务注册与发现
- 消息发布与订阅
- 分布式锁



### 核心架构

通过 etcdctl 客户端命令行操作和访问 etcd 中的数据

通过 HTTP API 接口直接访问 etcd



etcd 目前有 V2.x 和 V3.x 两个大版本，接口不一样、存储不一样，两个版本的数据相互隔离



- etcd server：对外接收和处理客户端的请求
- grpc server：etcd 与其他 etcd 节点之间的通信和信息同步
- MVCC：多版本控制，etcd 的存储模块，键值对的每一次操作行为都会被记录存储，数据底层存储在 BoltDB 数据库中
- WAL：预写式日志，etcd 中的数据提交前都会记录到日志
- snapshot：快照，以防 WAL 日志过多，用于存储某一时刻 etcd 的所有数据
- WAL + snapshot：etcd 可以有效地进行数据存储和节点故障恢复等操作
- raft



## References

https://www.bilibili.com/video/BV1344y1n77t?p=1

https://mp.weixin.qq.com/s/WA_EStxexM9pxDErPNXXcg