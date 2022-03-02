# linux性能优化实战



## CPU

### 1. 平均负载

单位时间内，系统处于**可运行状态**和**不可中断状态**的平均进程数，也就是平均活跃进程数，它和 CPU 使用率并没有直接关系。

- 可运行状态的进程，是指正在使用 CPU 或者正在等待 CPU 的进程，也就是我们常用 ps 命令看到的，处于 R 状态（Running 或 Runnable）的进程。
- 不可中断状态的进程，是正处于内核态关键流程中的进程，并且这些流程是不可打断的，比如最常见的是等待硬件设备的 I/O 响应，也就是我们在 ps 命令中看到的 D 状态（Uninterruptible Sleep，也称为 Disk Sleep）的进程。不可中断状态实际上是系统对进程和硬件设备的一种保护机制。



**平均负载最理想的情况是等于 CPU 个数**

```shell
grep 'model name' /proc/cpuinfo | wc -l

lscpu | grep 'CPU(s):'
```

**当平均负载高于 CPU 数量 70% 的时候**，就应该分析排查负载高的问题了。一旦负载过高，就可能导致进程响应变慢，进而影响服务的正常功能。但 70% 这个数字并不是绝对的，最推荐的方法，还是把系统的平均负载监控起来，然后根据更多的历史数据，判断负载的变化趋势。**当发现负载有明显升高趋势时，比如说负载翻倍了，再去做分析和调查**。



#### 案例

- stress 是一个 Linux 系统压力测试工具，用作异常进程模拟平均负载升高的场景。

- mpstat 是一个常用的多核 CPU 性能分析工具，用来实时查看每个 CPU 的性能指标，以及所有 CPU 的平均指标。
- pidstat 是一个常用的进程性能分析工具，用来实时查看进程的 CPU、内存、I/O 以及上下文切换等性能指标。



##### 场景一：CPU密集型进程

安装 stress、sysstat

```shell
apt install stress sysstat
```



stress 命令模拟一个 CPU 使用率 100% 的场景

```shell
$ stress -c 1 -t 600
```



运行 uptime 查看平均负载的变化情况

```shell
# -d 参数表示高亮显示变化的区域
$ watch -d uptime
..., load average: 1.00, 0.75, 0.39
```



mpstat 查看 CPU 使用率的变化情况

```shell
# -P ALL 表示监控所有CPU，后面数字5表示间隔5秒后输出一组数据
$ mpstat -P ALL 5
Linux 4.15.0 (ubuntu) 09/22/18 _x86_64_ (2 CPU)
13:30:06 CPU   %usr %nice %sys %iowait %irq %soft %steal %guest %gnice  %idle
13:30:11 all  50.05  0.00 0.00    0.00 0.00  0.00   0.00   0.00   0.00  49.95
13:30:11   0   0.00  0.00 0.00    0.00 0.00  0.00   0.00   0.00   0.00 100.00
13:30:11   1 100.00  0.00 0.00    0.00 0.00  0.00   0.00   0.00   0.00   0.00
```



pid 查看消耗 CPU 的进程

```shell
# 间隔5秒后输出一组数据
$ pidstat -u 5 1
13:37:07 UID  PID   %usr %system %guest %wait   %CPU CPU Command
13:37:12   0 2962 100.00    0.00   0.00  0.00 100.00   1  stress
```

stress 进程的 CPU 使用率为 100%。



##### 场景二：IO密集型进程

stress 模拟IO压力，不停地 sync

```shell
$ stress -i 1 -t 600
```



uptime 查看平均负载的变化情况

```shell
$ watch -d uptime
..., load average: 1.06, 0.58, 0.37
```



mpstat 查看 CPU 使用率的变化情况

```shell
# 显示所有CPU的指标，并在间隔5秒输出一组数据
$ mpstat -P ALL 5 1Linux 4.15.0 (ubuntu) 09/22/18 _x86_64_ (2 CPU)
13:41:28 CPU %usr %nice  %sys %iowait %irq %soft %steal %guest %gnice %idle
13:41:33 all 0.21  0.00 12.07   32.67 0.00  0.21   0.00   0.00   0.00 54.84
13:41:33   0 0.43  0.00 23.87   67.53 0.00  0.43   0.00   0.00   0.00  7.74
13:41:33   1 0.00  0.00  0.81    0.20 0.00  0.00   0.00   0.00   0.00 98.99
```

1 分钟的平均负载会慢慢增加到 1.06，其中一个 CPU 的系统 CPU 使用率升高到了 23.87，而 iowait 高达 67.53%。

这说明，平均负载的升高是由于 iowait 的升高。



pidstat 查看IO消耗高的进程

```shell
# 间隔5秒后输出一组数据，-u表示CPU指标
$ pidstat -u 5 1
Linux 4.15.0 (ubuntu) 09/22/18 _x86_64_ (2 CPU)
13:42:08 UID  PID %usr %system %guest %wait  %CPU CPU      Command
13:42:13   0  104 0.00    3.39   0.00  0.00  3.39   1 kworker/1:1H
13:42:13   0  109 0.00    0.40   0.00  0.00  0.40   0 kworker/0:1H
13:42:13   0 2997 2.00   35.53   0.00  3.99 37.52   1       stress
13:42:13   0 3057 0.00    0.40   0.00  0.00  0.40   0      pidstat
```



##### 场景三：大量进程

stress 模拟大量进程

```shell
$ stress -c 8 -t 600
```



uptime 查看平均负载的变化情况

```shell
$ uptime
..., load average: 7.97, 5.93, 3.02
```



pid 查看进程状态

```shell
# 间隔5秒后输出一组数据
$ pidstat -u 5 1
14:23:25 UID  PID  %usr %system %guest %wait  %CPU CPU Command
14:23:30   0 3190 25.00    0.00   0.00 74.80 25.00   0  stress
14:23:30   0 3191 25.00    0.00   0.00 75.20 25.00   0  stress
14:23:30   0 3192 25.00    0.00   0.00 74.80 25.00   1  stress
14:23:30   0 3193 25.00    0.00   0.00 75.00 25.00   1  stress
14:23:30   0 3194 24.80    0.00   0.00 74.60 24.80   0  stress
14:23:30   0 3195 24.80    0.00   0.00 75.00 24.80   0  stress
14:23:30   0 3196 24.80    0.00   0.00 74.60 24.80   1  stress
14:23:30   0 3197 24.80    0.00   0.00 74.80 24.80   1  stress
14:23:30   0 3200  0.00    0.20   0.00  0.20  0.20   0 pidstat
```

8 个进程在争抢 2 个 CPU，每个进程等待 CPU 的时间（也就是代码块中的 %wait 列）高达 75%。这些超出 CPU 计算能力的进程，最终导致 CPU 过载。



#### 小结

- 平均负载高有可能是 CPU 密集型进程导致的

- 平均负载高并不一定代表 CPU 使用率高，还有可能是 I/O 更繁忙了

- 当发现负载高的时候，你可以使用 mpstat、pidstat 等工具，辅助分析负载的来源



### 2. 上下文切换



## Memory



## IO



## Network



## References

https://time.geekbang.org/column/intro/100020901?tab=catalog
