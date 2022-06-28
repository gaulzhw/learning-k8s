# Learning C

## 静态库
libxx.a: xx指代库名

gcc -c yyy.c -> 生成.o文件

ar -cr libxx.a yyy.o



发布到

/usr/local/include: .h头文件声明

/usr/local/lib: 二进制实现 libxx.a

gcc -L/usr/local/lib -o main main.c -lxx

-l 参数必须在最右，有依赖



## 动态库

libxx.so: xx指代库名

gcc -shared -fpic -o libxx.so yyy.c

发布到

/usr/local/include

/usr/local/lib

在/etc/ld.so.conf中添加路径

执行/sbin/ldconfig 重读 /etc/ld.so.conf

gcc -I/usr/local/include -L/usr/local/lib -o ... -lxx

ldd - print shared library dependencies



非root用户发布

cp xx.so ~/lib

export LD_LIBRARY_PATH=~/lib



## tree

深度（层数）

度

叶子节点

孩子节点

兄弟节点

堂兄弟节点

二叉树

满二叉树

完全二叉树

存储: 顺序 链式

遍历: 按行，先序，中序，后序



## I/O

input & output, 一切实现的基础

### stdio

标准IO: FILE类型贯穿始终

- fopen
- fclose
- fgetc
- fputc
- fgets
- fputs
- fread
- fwrite
- printf
- scanf
- fseek
- ftell
- rewind
- fflush



### sysio

系统调用IO(文件IO)，文件描述符是在文件IO中贯穿始终的类型

文件描述符：整形数，数组下标，文件描述符优先使用当前可用范围内最小的

文件IO操作：open, close, read, write, lseek



文件IO与标准IO的区别：响应速度、吞吐量

标准IO与文件IO不能混用

转换：fileno, fdopen



IO的效率

文件共享: 多个任务共同操作一个文件或者协同完成任务

补充函数: truncate ftruncate

原子操作

程序中的重定向: dup, dup2

同步: sync, fsync, fdatasync

fcntl: 文件描述符所变的魔术几乎都来源于该函数

ioctl: 设备相关的内容

/dev/fd: 虚目录，显示的是当前进程的文件描述符信息



## 文件系统

类ls的实现，myls -a -i -l -n

### 目录和文件

- 获取文件属性
  - stat: 通过文件路径获取属性，面对符号链接文件时获取的是所指向的目标文件的属性
  - fstat: 通过文件描述符获取属性
  - lstat: 面对符号链接文件时获取的是符号链接文件的属性

- 文件访问权限: st_mode是一个16位的位图，用于表示文件类型、文件访问权限、特殊权限位

- umask: 防止产生权限过松的文件

- 文件权限的更改/管理: chmod fchmod

- 粘住位: t位

- 文件系统: FAT, UFS
  文件系统: 文件或数据的存储和管理

- 硬链接，符号链接
  硬链接与目录项是同义词，且建立硬链接有限制，不能给分区建立，不能给目录建立
  符号链接可跨分区，可以给目录建立
  link
  unlink
  remove (rm)
  rename (mv)

- utime: 可更改文件最后读的时间和最后修改的时间

- 目录的创建和销毁
  mkdir
  rmdir

- 更改当前工作路径
  chdir
  fchdir
  getcwd (pwd)

- 分析目录/读取目录内容
  glob: 解析模式/通配符

  opendir
  closedir
  readdir
  rewinddir
  seekdir
  telldir



### 系统数据文件和信息

/etc/passwd
  getpwuid
  getpwnam

/etc/group
  getgrgid
  getgrgrnam

/etc/shadow
  getspnam
  crypt
  getpass

时间戳: time_t  char *  struct tm
  time
  gmtime
  localtime
  mktime
  strftime



### 进程环境

main函数
  int main(int argc, char *argv[])

进程的终止
- 正常终止
  1. 从main函数返回
  2. 调用exit
  3. 调用_exit或_Exit
  4. 最后一个线程从其启动例程返回
  5. 最后一个线程调用pthread_exit
- 异常终止
  1. 调用abort
  2. 接到一个信号并终止
  3. 最后一个线程对其取消请求作出响应
- atexit 钩子函数

命令行参数的分析
  getopt: -
  getopt_long: --

环境变量

C程序的存储空间布局
  pmap: 查看进程内存空间分配

库
- 动态库
- 静态库
- 手工装载库
  dlopen
  dlclose
  dlerror
  dlsym

函数跳转
  区别于goto(只能跳转到本函数内)
  setjmp: 设置跳转点
  longjmp: 跳转到跳转点

资源的获取及控制
  getrlimit
  setrlimit



## 进程

进程标识符 pid
  类型: pid_t
  命令: ps
  进程号是顺次向下使用
  getpid
  getppid

进程的产生
  fork
    注意理解关键字: duplicating，意味着拷贝、克隆、一模一样等含义
    fork后父子进程的区别:
    1. fork的返回值不一样
    2. pid不同
    3. ppid不同
    4. 未决信号和文件锁不继承
    5. 资源利用量清0
    init进程: 1号进程，所有进程的祖先进程
    调度器的调度策略来决定哪个进程先运行
    fflush的重要性
  vfork

进程的消亡及释放资源
  wait
  waitpid
  waitid
  wait3
  wait4

exec函数族
  execl
  execlp
  execle
  execv
  execvp
  注意：fflush

用户权限及组权限: u+s、g+s
  getuid
  geteuid
  getgid
  getegid
  setuid
  setgid
  setreuid
  setregid
  seteuid
  setegid

解释器文件
  unix讲求的是机制而不是策略，永远只告诉能做什么而不是怎么做的

system
  理解: fork+exec+wait

进程会计
  acct

进程时间
  times

守护进程
  会话 session，标识sid
  终端
  setsid
  getpgrp
  getpgid
  setpgid
  单实例守护进程: 锁文件 /var/run/[name.pid]
  启动脚本文件: /etc/rc*

系统日志
  /var/lib
  主日志messages
  syslogd服务

  openlog
  syslog
  closelog



## 并发
同步
异步

异步事件的处理: 查询法、通知法

coredump    ulimit -c

### 信号
信号的概念
  信号是软件中断
  信号的响应依赖于中断

signal
  void (*signal(int signum, void (*func)(int)))(int)
  信号会打断阻塞的系统调用

信号的不可靠
  信号的行为不可靠

可重入函数
  解决信号不可靠
  所有的系统调用都是可重入的，一部分库函数也是可重入的，如: memcpy

信号的响应过程
  信号从收到到响应有一个不可避免的延迟
  如何忽略掉一个信号
  标准信号为什么要丢失
  标准信号的响应没有严格的顺序
  不能从信号处理函数中随意的往外跳  sigsetjmp、siglongjmp

常用函数
  kill
  raise
  alarm
  pause
  abort
  system
  sleep、nanosleep、usleep
  select

信号集
  信号集类型: sigset_t
  sigemptyset
  sigfillset
  sigaddset
  sigdelset
  sigismember

信号屏蔽字/pending集的处理
  sigprocmask
  sigpending

扩展
  sigsuspend
  sigaction -> signal
  setitimer

实时信号

### 线程

线程的概念
  一个正在运行的函数
  posix线程是一套标准，而不是实现
  openmp线程
  线程标识: pthread_t
  pthread_equal
  pthread_self

线程的创建
  pthread_create
  线程的调度取决于调度器策略

线程的终止
  3种方式: 
    线程从启动例程返回，返回值就是线程的退出码
    线程可以被同一进程中的其他线程取消
    线程调用pthread_exit函数
  pthread_join --> wait

栈清理
  pthread_cleanup_push
  pthread_cleanup_pop

线程的取消选项
  pthread_cancel
  取消有2种状态: 允许、不允许
  允许取消分为: 异步cancel、推迟cancel(默认)->推迟至cance点再响应
  cancel点: POSIX定义的cancel点，都是可能引发阻塞的系统调用
  pthread_setcanceltype: 设置取消方式
  pthread_setcancelstate: 设置是否允许取消
  pthread_testcancel: 本函数什么都不做，就是一个取消点

线程分离
  pthread_detach

线程同步
  互斥量
    pthread_mutex_t
    pthread_mutex_init
    pthread_mutex_destroy
    pthread_mutex_lock
    pthread_mutex_trylock
    pthread_mutex_unlock
    pthread_once
  条件变量
    pthread_cond_t
    pthread_cond_init
    pthread_cond_destroy
    pthread_cond_broadcast
    pthread_cond_signal
    pthread_cond_wait
    pthread_cond_timedwait
  信号量
  读写锁: 读锁->共享锁，写锁->互斥锁

线程属性
  pthread_attr_init
  pthread_attr_destroy
  pthread_attr_setxxx
  pthread_attr_getxxx
  见man pthread_attr_init的see also

线程同步的属性
  互斥量属性
    pthread_mutexattr_init
    pthread_mutexattr_destroy
    pthread_mutexattr_getpshared
    pthread_mutexattr_setpshared
    clone
    pthread_mutexattr_gettype
    pthread_mutexattr_settype
  条件变量属性
    pthread_condattr_init
    pthread_condattr_destroy
  读写锁属性

重入
  多线程中的IO

线程与信号
  pthread_sigmask
  sigwait
  pthread_kill

线程与fork

openmp -> www.OpenMP.org



## 高级IO

非阻塞IO -- 阻塞IO
补充: 有限状态机编程

非阻塞IO
  简单流程: 自然流程是结构化的
  复杂流程: 自然流程不是结构化的

IO多路转接
  select
  poll
  epoll
    epoll_create
    epoll_ctl
    epoll_wait

其他读写函数
  readv
  writev

存储映射IO
  mmap
  munmap

文件锁
  文件锁作用在inode上
  fcntl
  lockf
  flock



## 进程间通信

管道
  内核提供，单工，自同步机制
  匿名管道
    pipe
  命名管道: 文件类型为p
    mkfifo

XSI -> SysV
  主动端: 先发包的一方
  被动端: 先收包的一方(先运行)
  命令: ipcs、ipcrm
  key: ftok

  Message Queues
    ulimit -q: POSIX message queues 约束消息的大小
    msgget
    msgop: msgsnd、msgrcv
    msgctl
  Semaphore Arrays
    semget
    semop
    semctl
  Shared Memory
    shmget
    shmop: shmat、shmdt
    shmctl

网络套接字
  跨主机的传输要注意的问题
  1. 字节序
    大端: 低地址处放高字节
    小端: 低地址处放低字节
    主机字节序: host
    网络字节序: network
    解决方法 _ to _ _: htons、htonl、ntohs、ntohl
  2. 对齐
    解决方法 不对齐
  3. 类型长度
    解决方法 int32_t、uint32_t、int64_t、int8_t、uint8_t

  socket

  报式套接字
    socket
    bind
    sendto
    rcvfrom
    inet_pton
    inet_ntop
    setsockopt
    getsockopt
    多点通讯: 广播(全网广播、子网广播)、多播/组播
  被动端: 先运行
    1. 取得socket
    2. 给socket取得地址
    3. 收/发消息
    4. 关闭socket
  主动端
    1. 取得socket
    2. 给socket取得地址(可省略)
    3. 发/收消息
    4. 关闭socket

  流式套接字
  C端(主动端)
    1. 获取socket
    2. 给socket取得地址(可省略)
    3. 发送连接
    4. 收/发消息
    5. 关闭
  S端
    1. 获取socket
    2. 给socket取得地址
    3. 将socket置为监听模式
    4. 接受连接
    5. 收/发消息
    6. 关闭
