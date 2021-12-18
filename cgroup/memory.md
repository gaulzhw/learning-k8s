# memory cgroup



## OOM

linux判断选择OOM kill的进程判断标准涉及两个条件：

1. 进程已经使用的物理内存页面数
2. 每个进程的OOM校准值oom_score_adj，在/proc文件系统中，每个进程都有一个/proc/[pid]/oom_score_adj的文件，可以在这个文件中输入-1000到1000之间的任意一个数值，调整进程被OOM kill的几率

```c
# oom_badness()

adj = (long)p->signal->oom_score_adj;

points = get_mm_rss(p->mm) + get_mm_counter(p->mm, MM_SWAPENTS) +mm_pgtables_bytes(p->mm) / PAGE_SIZE;

adj *= totalpages / 1000;
points += adj;
```

函数oom_badness()计算方式：

用系统总的可用页面数，乘以OOM校准值oom_score_adj，再加上进程已经使用的物理页面数，计算出来的值越大，那么这个进程被OOM Kill的几率也就越大



## 与OOM相关的三个参数

- memory.limit_in_bytes：限制进程组的可用内存的最大值
- memory.oom_control：控制OOM后是否kill进程，默认触发kill，设置为1可以不kill
- memory.usage_in_bytes：当前控制组所有进程实际使用的内存总和



如果memory.oom_control设置为1，那么容器中的进程在使用内存到达memroy.limit_in_bytes之后，不会被kill掉，但memalloc进程会被暂停申请内存，状态会因等待资源申请而变成TASK INTERRUPTABLE



## OOM相关日志

查看内核日志，使用journalctl -k或者直接查看日志文件/var/log/message

![oom-log](img/oom-log.png)

日志大致分为三部分：

1. 容器里每一个进程使用的内存页面数量，在rss列，rss是Resident Set Size的缩写，指的是进程真正在使用的物理内存页面数量
2. oom-kill，这一行列出了发生OOM的memory cgroup的控制组，从控制组的信息中知道OOM是在哪个容器发生的
3. killed process 7445(men_alloc)，显示了最终被OOM killer杀死的进程



问题分析：

1. 进程本身的确需要很大的内存，说明需要给memory.limit_in_bytes里的内存上限值设置小了，需要增大内存的上限值
2. 进程的代码有bug，会导致内存泄漏，进程内存使用到达了memory cgroup中的上限，需要具体去解决代码问题



## OOM结束进程信号

OOM killer是发送sigkill信号结束进程



## k8s与内存的关系

k8s中设置memory的request、limit

- request不修改memory cgroup里的参数，只是在scheduler里调度的时候做计算看是否可以分配内存

- limit设置memory.limit_in_bytes的值
