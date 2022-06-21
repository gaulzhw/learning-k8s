#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <unistd.h>

// mac没有实时信号kill -l
#define MYRTSIG (SIGRTMIN+6)

static void mysig_handler(int s)
{
    write(1, "!", 1);
}

int main()
{
    sigset_t set;

    signal(MYRTSIG, mysig_handler);
    sigemptyset(&set);
    sigaddset(&set, MYRTSIG);

    for (int j = 0; j < 1000; j++)
    {
        sigprocmask(SIG_BLOCK, &set, NULL);

        for (int i = 0; i < 5; i++)
        {
            write(1, "*", 1);
            sleep(1);
        }

        write(1, "\n", 1);
        sigsuspend(&set);
    }
    exit(0);
}