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



## etcdctl

etcdctl 支持的命令大体上分为数据库操作和非数据库操作两类

### 常用命令

| 命令                   | 描述                                       |
| ---------------------- | ------------------------------------------ |
| alarm disarm           | 解除所有的报警                             |
| alarm list             | 列出所有的报警                             |
| auth disable           | 禁用 authentication                        |
| auth enable            | 启用 authentication                        |
| check datascale        | 对于给定服务实例，检查持有数据的存储使用率 |
| check perf             | 检查 etcd 集群的性能表现                   |
| compaction             | 压缩 etcd 中的事件历史                     |
| defrag                 | 整理给定 etcd 实例的存储碎片               |
| del                    | 移除指定范围 [key, range_end) 的键值对     |
| elect                  | 加入 leader 选举                           |
| endpoint hashkv        | 打印指定 etcd 实例的历史键值对 hash 信息   |
| endpoint health        | 打印指定 etcd 实例的健康信息               |
| endpoint status        | 打印指定 etcd 实例的状态信息               |
| get                    | 获取键值对                                 |
| help                   | 帮助命令                                   |
| lease grant            | 创建 leases                                |
| lease keep-alive       | 刷新 leases                                |
| lease list             | 罗列所有有效的 leases                      |
| lease revoke           | 撤销 leases                                |
| lease timetolive       | 获取 lease 信息                            |
| lock                   | 获取一个命名锁                             |
| make-mirror            | 指定一个 etcd 集群作为镜像集群             |
| member add             | 增加一个成员到集群                         |
| member list            | 列出集群的所有成员                         |
| member promote         | 提升集群中的一个 non-voting 成员           |
| member remove          | 移除集群中的成员                           |
| member update          | 更新集群中的成员信息                       |
| migrate                | 迁移 v2 存储中的键值对到 MVCC 存储         |
| move-leader            | 移除 etcd 集群的 leader 给另一个 etcd 成员 |
| put                    | 写入键值对                                 |
| role add               | 增加一个角色                               |
| role delete            | 删除一个角色                               |
| role get               | 获取某个角色的详细信息                     |
| role grant-permission  | 给某个角色授予 key                         |
| role list              | 罗列所有的角色                             |
| role revoke-permission | 撤销一个角色的 key                         |
| snapshot restore       | 恢复快照                                   |
| snapshot save          | 存储某一个 etcd 节点的快照文件至指定位置   |
| snapshot status        | 获取指定文件的后端快照文件状态             |
| txn                    | TXN 在一个事务内处理所有的请求             |
| user add               | 增加一个用户                               |
| user delete            | 删除某个用户                               |
| user get               | 获取某个用户的详细信息                     |
| user grant-role        | 将某个角色授予某个用户                     |
| user list              | 列出所有的用户                             |
| user passwd            | 更改某个用户的密码                         |
| user revoke-role       | 撤销某个用户的角色                         |
| version                | 输出 etcdctl 的版本                        |
| watch                  | 监测指定键或者前缀的事件流                 |

### options

| 选项                               | 描述                                                         |
| ---------------------------------- | ------------------------------------------------------------ |
| --cacert=""                        | 服务端使用 HTTPS 时，使用 CA 文件进行验证                    |
| --cert=""                          | HTTPS 下客户端使用的 SSL 证书文件                            |
| --command-timeout=5s               | 命令执行超时时间设置                                         |
| --debug[=false]                    | 输出 CURL 命令，显示执行命令时发起的请求日志                 |
| --dial-timeout=2s                  | 客户端连接超时时间                                           |
| -d, --discovery-srv=""             | 用于查询描述集群端点 SRV 记录的域名                          |
| --discovery-srv-name=""            | 使用 DNS 发现时，查询的服务名                                |
| --endpoints=[127.0.0.1:2379]       | gRPC 端点                                                    |
| -h, --help[=false]                 | etcdctl 帮助                                                 |
| --hex[=false]                      | 输出二进制字符串为十六进制编码的字符串                       |
| --insecure-discovery[=true]        | 接受集群成员中不安全的 SRV 记录                              |
| --insecure-skip-tls-verify[=false] | 跳过服务端证书认证                                           |
| --insecure-transport[=true]        | 客户端禁用安全传输                                           |
| --keepalive-time=2s                | 客户端连接的 keepalive 时间                                  |
| --keepalive-timeout=6s             | 客户端连接的 keepalive 的超时时间                            |
| --key=""                           | HTTPS 下客户端使用的 SSL 密钥文件                            |
| --password=""                      | 认证的密码，当该选项开启，--user 参数中不要包含密码          |
| --user=""                          | username[:password] 的形式                                   |
| -w, --write-out="simple"           | 输出内容的格式（Fields、Json、Protobuf、Simple、Talbe，其中 Simple 为原始信息；Json 为使用 Json 格式解码，易读性高） |



## References

https://www.bilibili.com/video/BV1344y1n77t?p=1

https://mp.weixin.qq.com/s/WA_EStxexM9pxDErPNXXcg