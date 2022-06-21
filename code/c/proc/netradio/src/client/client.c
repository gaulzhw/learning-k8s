#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <net/if.h>
#include <arpa/inet.h>
#include <getopt.h>
#include <errno.h>

#include <proto.h>
#include "client.h"

struct client_conf_st client_conf = {
    .rcvport = DEFAULT_RCVPORT,
    .mgroup = DEFAULT_MGROUP,
    .player_cmd = DEFAULT_PLAYERCMD};

static void printhelp(void)
{
    printf("-P --port   指定接收端口\n"
           "-M --mgroup 指定多播组\n"
           "-p --player 指定播放器\n"
           "-H --help   显示帮助\n");
}

static ssize_t writen(int fd, const char *buf, size_t len)
{
    int ret, pos = 0;

    while (1)
    {
        ret = write(fd, buf + pos, len);
        if (ret < 0)
        {
            if (errno == EINTR)
                continue;
            perror("write()");
            return -1;
        }
        len -= ret;
        pos += ret;
    }

    return pos;
}

int main(int argc, char *argv[])
{
    int index = 0;
    int c;
    int sd;
    int pd[2];
    int len;
    int chosenid;

    struct ip_mreqn mreq;
    struct sockaddr_in laddr, serveraddr, raddr;
    socklen_t serveraddr_len, raddr_len;

    struct option argarr[] = {
        {"port", 1, NULL, 'P'},
        {"mgroup", 1, NULL, 'M'},
        {"player", 1, NULL, 'p'},
        {"help", 0, NULL, 'H'},
        {NULL, 0, NULL, 0},
    };

    /*
     * 初始化
     * 级别: 默认值、配置文件、环境变量、命令行参数
     */

    while (1)
    {
        c = getopt_long(argc, argv, "P:M:p:H", argarr, &index);
        if (c < 0)
            break;

        switch (c)
        {
        case 'P':
            client_conf.rcvport = optarg;
            break;
        case 'M':
            client_conf.mgroup = optarg;
            break;
        case 'p':
            client_conf.player_cmd = optarg;
            break;
        case 'H':
            printhelp();
            exit(0);
            break;
        default:
            abort();
            break;
        }
    }

    sd = socket(AF_INET, SOCK_DGRAM, 0);
    if (sd < 0)
    {
        perror("socket()");
        exit(1);
    }

    inet_pton(AF_INET, client_conf.mgroup, &mreq.imr_multiaddr);
    inet_pton(AF_INET, "0.0.0.0", &mreq.imr_address);
    mreq.imr_ifindex = if_nametoindex("eth0");

    if (setsockopt(sd, IPPROTO_IP, IP_ADD_MEMBERSHIP, &mreq, sizeof(mreq)) < 0)
    {
        perror("setsockopt()");
        exit(1);
    }

    int val = 1;
    if (setsockopt(sd, IPPROTO_IP, IP_MULTICAST_LOOP, &val, sizeof(val)) < 0)
    {
        perror("setsockopt()");
        exit(1);
    }

    laddr.sin_family = AF_INET;
    laddr.sin_port = htons(atoi(client_conf.rcvport));
    inet_pton(AF_INET, "0.0.0.0", &laddr.sin_addr);

    if (bind(sd, (void *)&laddr, sizeof(laddr)) < 0)
    {
        perror("bind()");
        exit(1);
    }

    if (pipe(pd) < 0)
    {
        perror("pipe()");
        exit(1);
    }

    pid_t pid = fork();
    if (pid < 0)
    {
        perror("fork()");
        exit(1);
    }
    if (pid == 0)
    {
        // child: 调用解码器
        close(sd);
        close(pd[1]);
        dup2(pd[0], 0);
        if (pd[0] > 0)
            close(pd[0]);

        execl("/bin/sh", "sh", "-c", client_conf.player_cmd, NULL);
        perror("execl()");
        exit(1);
    }

    // parent: 从网络收包，发送给子进程

    // 收节目单
    struct msg_list_st *msg_list;

    msg_list = malloc(MSG_LIST_MAX);
    if (msg_list == NULL)
    {
        perror("malloc()");
        exit(1);
    }

    while (1)
    {
        len = recvfrom(sd, msg_list, MSG_LIST_MAX, 0, (void *)&serveraddr, &serveraddr_len);
        if (len < sizeof(struct msg_list_st))
        {
            fprintf(stderr, "message is too small.\n");
            continue;
        }
        if (msg_list->chnid != LIST_CHNID)
        {
            fprintf(stderr, "chnid is not match.\n");
            continue;
        }
        break;
    }

    // 打印节目单，选择频道
    struct msg_listentry_st *pos;

    for (pos = msg_list->entry; (char *)pos < ((char *)msg_list + len); pos = (void *)(((char *)pos) + ntohs(pos->len)))
    {
        fprintf(stdout, "channel %d : %s\n", pos->chnid, pos->desc);
    }

    free(msg_list);
    while (1)
    {
        int ret = scanf("%d", &chosenid);
        if (ret != 1)
            exit(1);
    }

    // 收频道包，发送给子进程
    struct msg_channel_st *msg_channel;

    msg_channel = malloc(MSG_CHANNEL_MAX);
    if (msg_channel == NULL)
    {
        perror("malloc()");
        exit(1);
    }

    while (1)
    {
        len = recvfrom(sd, msg_channel, MSG_CHANNEL_MAX, 0, (void *)&raddr, &raddr_len);
        if (raddr.sin_addr.s_addr != serveraddr.sin_addr.s_addr || raddr.sin_port != serveraddr.sin_port)
        {
            fprintf(stderr, "Ignore: address or port not match.\n");
            continue;
        }
        if (len < sizeof(struct msg_channel_st))
        {
            fprintf(stderr, "Ignore: message too small.\n");
            continue;
        }
        if (msg_channel->chnid == chosenid)
        {
            fprintf(stdout, "accepted msg: %d received.\n", msg_channel->chnid);
            if (writen(pd[1], (void *)msg_channel->data, len - sizeof(chnid_t)) < 0)
                exit(1);
        }
    }

    free(msg_channel);
    close(sd);

    exit(0);
}