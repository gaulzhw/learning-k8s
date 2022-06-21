#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <arpa/inet.h>
#include <time.h>

#include "proto.h"

#define IPSTRSIZE 40
#define BUFSIZE 1024
#define PROCNUM 4

static void server_job(int sd)
{
    char buf[BUFSIZE];
    int len;

    len = sprintf(buf, FMT_STAMP, (long long)time(NULL));

    if (send(sd, buf, len, 0) < 0)
    {
        perror("send()");
        exit(1);
    }
}

static void server_loop(int sd)
{
    struct sockaddr_in raddr;
    socklen_t raddr_len;
    char ipstr[IPSTRSIZE];

    raddr_len = sizeof(raddr);

    while (1)
    {
        int newsd = accept(sd, (void *)&raddr, &raddr_len);
        if (newsd < 0)
        {
            perror("accept()");
            exit(1);
        }

        inet_ntop(AF_INET, &raddr.sin_addr, ipstr, IPSTRSIZE);
        printf("client: %s:%d\n", ipstr, ntohs(raddr.sin_port));

        server_job(newsd);
        close(newsd);
    }
}

int main()
{
    int sd;
    struct sockaddr_in laddr;

    sd = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    if (sd < 0)
    {
        perror("socket()");
        exit(1);
    }

    int val = 1;
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

    if (listen(sd, 200) < 0)
    {
        perror("listen()");
        exit(1);
    }

    for (int i = 0; i < PROCNUM; i++)
    {
        pid_t pid = fork();
        if (pid < 0)
        {
            perror("fork()");
            exit(1);
        }
        if (pid == 0)
        {
            server_loop(pid);
            exit(0);
        }
    }

    for (int i = 0; i < PROCNUM; i++)
        wait(NULL);
    close(sd);
    exit(0);
}