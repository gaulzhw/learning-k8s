#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
    pid_t pid;

    printf("[%d]: Begin!\n", getpid());
    fflush(NULL); // !!! 在执行fork前刷新缓冲区

    pid = fork();
    if (pid < 0)
    {
        perror("fork()");
        exit(1);
    }
    if (pid == 0) // child
    {
        printf("[%d]: Child is working!\n", getpid());
    }
    else // parent
    {
        printf("[%d]: Parent is working!\n", getpid());
    }

    printf("[%d]: End!\n", getpid());
    exit(0);
}