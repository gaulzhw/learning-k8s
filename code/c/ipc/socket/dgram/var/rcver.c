#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <arpa/inet.h>
#include <unistd.h>

#include "proto.h"

#define IPSTRSIZE 40

int main()
{
    int sd;
    int size;
    struct sockaddr_in laddr, raddr;
    struct msg_st *rbufp;
    socklen_t raddr_len = sizeof(raddr); // !!!
    char ipstr[IPSTRSIZE];

    size = sizeof(struct msg_st) + NAMEMAX;
    rbufp = malloc(size);
    if (rbufp == NULL)
    {
        perror("malloc()");
        exit(1);
    }

    sd = socket(AF_INET, SOCK_DGRAM, 0 /*IPPROTO_UDP*/);
    if (sd < 0)
    {
        perror("socket()");
        exit(1);
    }

    laddr.sin_family = AF_INET;
    laddr.sin_port = htons(atoi(RCVPORT));
    inet_pton(AF_INET, "0.0.0.0", &laddr.sin_addr);

    if (bind(sd, (void *)&laddr, sizeof(laddr)) < 0)
    {
        perror("bind()");
        exit(1);
    }

    while (1)
    {
        recvfrom(sd, rbufp, size, 0, (void *)&raddr, &raddr_len);

        inet_ntop(AF_INET, &raddr.sin_addr, ipstr, IPSTRSIZE);
        printf("--- MESSAGE FROM %s:%d ---\n", ipstr, ntohs(raddr.sin_port));
        printf("NAME = %s\n", rbufp->name);
        printf("MATH = %d\n", ntohl(rbufp->math));
        printf("CHINESE = %d\n", ntohl(rbufp->chinese));
    }

    close(sd);

    exit(0);
}