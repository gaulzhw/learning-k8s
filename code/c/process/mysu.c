#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

/*
 * ./mysu 0 cat /etc/shadow
 * 使用前需要先设置mysu权限为u+s
 * rws
 */

int main(int argc, char **argv)
{
    if (argc < 3)
    {
        fprintf(stderr, "Usage...\n");
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
        setuid(atoi(argv[1]));
        execvp(argv[2], argv+2);
        perror("execvp()");
        exit(1);
    }

    wait(NULL);

    exit(0);
}