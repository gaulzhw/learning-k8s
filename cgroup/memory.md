# memory cgroup



## OOM

```c
# oom_badness()

adj = (long)p->signal->oom_score_adj;

points = get_mm_rss(p->mm) + get_mm_counter(p->mm, MM_SWAPENTS) +mm_pgtables_bytes(p->mm) / PAGE_SIZE;

adj *= totalpages / 1000;
points += adj;
```

用系统总的可用页面数，乘以OOM校准值oom_score_adj，再加上进程已经使用的物理页面数，计算出来的值越大，那么这个进程被OOM Kill的几率也就越大



## 与OOM相关的三个参数

- memory.limit_in_bytes：限制进程组的可用内存的最大值
- memory.oom_control：控制OOM后是否kill进程，默认触发kill，设置为1可以不kill
- memory.usage_in_bytes：当前控制组所有进程实际使用的内存总和



如果memory.oom_control设置为1，那么容器中的进程在使用内存到达memroy.limit_in_bytes之后，不会被kill掉，但memalloc进程会被暂停申请内存，状态会因等待资源申请而变成TASK INTERRUPTABLE

