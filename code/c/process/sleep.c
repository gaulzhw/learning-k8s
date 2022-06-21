#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
    puts("Begin!");

    fflush(NULL);

    pid_t pid = fork();
    if (pid < 0)
    {
        perror("fork()");
        exit(1);
    }

    if (pid == 0)
    {
        execl("/bin/sleep", "sleep", "100", NULL);
        perror("execl()");
        exit(1);
    }

    wait(NULL);

    puts("End!");
    exit(0);
}