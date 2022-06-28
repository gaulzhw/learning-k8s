#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>

#define BUFSIZE 1024

int main()
{
    int pd[2];
    pid_t pid;
    char buf[BUFSIZE];
    int len;

    if (pipe(pd) < 0)
    {
        perror("pipe()");
        exit(1);
    }

    pid = fork();
    if (pid < 0)
    {
        perror("fork()");
        exit(1);
    }
    if (pid == 0)
    {
        close(pd[1]);
        dup2(pd[0], 0);
        close(pd[0]);

        int fd = open("/dev/null", O_RDWR);
        dup2(fd, 1);
        dup2(fd, 2);

        execl("/usr/local/bin/mpg123", "mpg123", "-", NULL);
        perror("execl()"); // unlikely
        exit(1);
    }

    close(pd[0]);

    // 父进程从网上收数据，往管道中写
    close(pd[1]);
    wait(NULL);
    exit(0);
}