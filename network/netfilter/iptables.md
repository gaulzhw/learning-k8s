# iptables



不同流向的数据包会经过不同的链：

- 到本机某进程的报文：PREROUTING --> INPUT
- 由本机转发的报文：PREROUTING --> FORWARD --> POSTROUTING
- 由本机的某进程发出报文：OUTPUT --> POSTROUTING



每个链上会有一些规则去过滤数据包进行操作，这些规则在大体上又可以分为4类，分别存在4张table中：

- filter表：负责过滤功能，防火墙；内核模块：iptables_filter
- nat表：network address translation，网络地址转换功能；内核模块：iptable_nat
- mangle表：拆解报文，做出修改，并重新封装 的功能；内核模块：iptable_mangle
- raw表：关闭nat表上启用的连接追踪机制；内核模块：iptable_raw



每条规则会通过一些条件去匹配数据包，比如：源IP，目的IP，端口，协议种类，出入网卡名称等等，匹配到数据包后还会指定target，也就是对这个数据包进行操作。操作包括：

- ACCEPT：允许数据包通过。
- DROP：直接丢弃数据包，不给任何回应信息，这时候客户端会感觉自己的请求泥牛入海了，过了超时时间才会有反应。
- REJECT：拒绝数据包通过，必要时会给数据发送端一个响应的信息，客户端刚请求就会收到拒绝的信息。
- SNAT：源地址转换，解决内网用户用同一个公网地址上网的问题。
- MASQUERADE：是SNAT的一种特殊形式，适用于动态的、临时会变的ip上。
- DNAT：目标地址转换。
- REDIRECT：在本机做端口映射。
- LOG：在内核日志中记录日志信息，然后将数据包传递给下一条规则，也就是说除了记录以外不对数据包做任何其他操作，仍然让下一条规则去匹配。



其中ACCEPT，DROP，REJECT，SNAT，MASQUERADE，REDIRECT，DNAT为终结操作，也就是匹配到它们之后不再继续匹配当前表中剩下的规则，如果数据包还存在的话就继续匹配后面剩的其他表中的规则。对于DROP和REJECT来讲，数据包经过它们后将不复存在，所以也就不再有后续的匹配了。其他的为非终结操作，操作完成后还会让数据包继续匹配当前表后面的规则。

>  举个例子：数据包A由是本机的某进程发出报文，经过OUTPUT --> POSTROUTING，在OUTPUT中首先依次进入raw，mangle，nat，filter这4个表，每个表都有自己的OUTPUT链，每条链都有若干规则，下面分析几种情况：
>
> 1. 如果A在nat表的OUTPUT的第一条规则中被REJECT或者DROP，那么A的处理到此为止；
> 2. 如果A在nat表的OUTPUT的第一条规则中被ACCEPT，那么它将进入filter表的OUTPUT链继续匹配规则；
> 3. 如果A在nat表的OUTPUT的第一条规则中被LOG，那么它将继续匹配nat表的第二条规则直到遇到终结操作。