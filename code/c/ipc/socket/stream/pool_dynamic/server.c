#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <arpa/inet.h>
#include <signal.h>
#include <sys/mman.h>
#include <errno.h>
#include <time.h>

#include "proto.h"

#define MIN_SPARE_SERVERS 5
#define MAX_SPARE_SERVERS 10
#define MAX_CLIENTS 20

#define SIG_NOTIFY SIGUSR1

#define LINEBUFSIZE 1024

enum
{
    STATE_IDLE = 0,
    STATE_BUSY
};

struct server_st
{
    pid_t pid;
    int state;
    // int reuse;
};

static struct server_st *serverpool;
static int idle_count = 0, busy_count = 0;
static int sd;

static void server_job(int pos)
{
    pid_t ppid;
    struct sockaddr_in raddr;
    socklen_t raddr_len;
    int client_sd;
    char linebuf[LINEBUFSIZE];

    ppid = getppid();

    while (1)
    {
        serverpool[pos].state = STATE_IDLE;
        kill(ppid, SIG_NOTIFY);

        client_sd = accept(sd, (void *)&raddr, &raddr_len);
        if (client_sd < 0)
        {
            if (errno != EINTR || errno != EAGAIN)
            {
                perror("accept()");
                exit(1);
            }
        }

        serverpool[pos].state = STATE_BUSY;
        kill(ppid, SIG_NOTIFY);

        // inet_ntop(AF_INET, &raddr.sin_addr, ipstr, IPSTRSIZE);
        // printf("[%d]client: %s:%d\n", getpid(), ipstr, ntohs(raddr.sin_port));

        long long stamp = time(NULL);
        int len = snprintf(linebuf, LINEBUFSIZE, FMT_STAMP, stamp);
        send(client_sd, linebuf, len, 0);

        close(client_sd);
    }
}

static int add_one_server(void)
{
    int slot;
    pid_t pid;

    if (idle_count + busy_count >= MAX_CLIENTS)
        return -1;

    for (slot = 0; slot < MAX_CLIENTS; slot++)
        if (serverpool[slot].pid == -1)
            break;

    serverpool[slot].state = STATE_IDLE;
    pid = fork();
    if (pid < 0)
    {
        perror("fork()");
        exit(1);
    }
    if (pid == 0)
    {
        server_job(slot);
        exit(0);
    }

    // parent
    serverpool[slot].pid = pid;
    idle_count++;
    return 0;
}

static int del_one_server(void)
{
    int i;

    if (idle_count == 0)
        return -1;

    for (i = 0; i < MAX_CLIENTS; i++)
    {
        if (serverpool[i].pid != -1 && serverpool[i].state == STATE_IDLE)
        {
            kill(serverpool[i].pid, SIGTERM);
            serverpool[i].pid = -1;
            idle_count--;
        }
    }
    return 0;
}

static void usr1_handler(int s)
{
    return;
}

static void scan_pool(void)
{
    int idle = 0, busy = 0;

    for (int i = 0; i < MAX_CLIENTS; i++)
    {
        if (serverpool[i].pid == -1)
            continue;
        // sig 0 检测进程是否存在
        if (kill(serverpool[i].pid, 0))
        {
            serverpool[i].pid = -1;
            continue;
        }
        if (serverpool[i].state == STATE_IDLE)
            idle++;
        else if (serverpool[i].state == STATE_BUSY)
            busy++;
        else
        {
            fprintf(stderr, "unknown state...\n");
            // _exit(1);
            abort(); // 杀掉当前进程，并且获取core dump文件
        }
    }
    idle_count = idle;
    busy_count = busy;
}

int main()
{
    struct sigaction sa, osa;
    int val = 1;
    struct sockaddr_in laddr;
    sigset_t set, oset;

    sa.sa_handler = SIG_IGN;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = SA_NOCLDWAIT;
    sigaction(SIGCHLD, &sa, &osa);

    sa.sa_handler = usr1_handler;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = 0;
    sigaction(SIG_NOTIFY, &sa, &osa);

    sigemptyset(&set);
    sigaddset(&set, SIG_NOTIFY);
    sigprocmask(SIG_BLOCK, &set, &oset);

    serverpool = mmap(NULL, sizeof(struct server_st) * MAX_CLIENTS, PROT_READ | PROT_WRITE, MAP_SHARED | MAP_ANONYMOUS, -1, 0);
    if (serverpool == MAP_FAILED)
    {
        perror("mmap()");
        exit(1);
    }

    for (int i = 0; i < MAX_CLIENTS; i++)
        serverpool[i].pid = -1;

    sd = socket(AF_INET, SOCK_STREAM, 0);
    if (sd < 0)
    {
        perror("socket()");
        exit(1);
    }

    if (setsockopt(sd, SOL_SOCKET, SO_REUSEADDR, &val, sizeof(val)) < 0)
    {
        perror("setsockopt()");
        exit(1);
    }

    laddr.sin_family = AF_INET;
    laddr.sin_port = htons(atoi(SERVERPORT));
    inet_pton(AF_INET, "0.0.0.0", &laddr.sin_addr);
    if (bind(sd, (void *)&laddr, sizeof(laddr)) < 0)
    {
        perror("bind()");
        exit(1);
    }

    if (listen(sd, 100) < 0)
    {
        perror("listen()");
        exit(1);
    }

    for (int i = 0; i < MIN_SPARE_SERVERS; i++)
    {
        add_one_server();
    }

    while (1)
    {
        sigsuspend(&oset);

        scan_pool();

        // control the pool
        if (idle_count > MAX_CLIENTS)
        {
            for (int i = 0; i < (idle_count - MAX_SPARE_SERVERS); i++)
                del_one_server();
        }
        else if (idle_count < MIN_SPARE_SERVERS)
        {
            for (int i = 0; i < (MIN_SPARE_SERVERS - idle_count); i++)
                add_one_server();
        }

        // printf the pool
        for (int i = 0; i < MAX_CLIENTS; i++)
        {
            if (serverpool[i].pid == -1)
                putchar(' ');
            else if(serverpool[i].state == STATE_IDLE)
                putchar('.');
            else
                putchar('x');
        }
        putchar('\n');
    }

    sigprocmask(SIG_SETMASK, &oset, NULL);

    exit(0);
}