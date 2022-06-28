#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <unistd.h>

static void int_handler(int s)
{
    write(1, "!", 1);
}

int main()
{
    sigset_t set, oset, saveset;

    signal(SIGINT, int_handler);
    sigemptyset(&set);
    sigaddset(&set, SIGINT);

    sigprocmask(SIG_UNBLOCK, &set, &saveset);

    for (int j = 0; j < 1000; j++)
    {
        // sigprocmask(SIG_BLOCK, &set, NULL);
        sigprocmask(SIG_BLOCK, &set, &oset);

        for (int i = 0; i < 5; i++)
        {
            write(1, "*", 1);
            sleep(1);
        }

        write(1, "\n", 1);
        // sigprocmask(SIG_UNBLOCK, &set, NULL);
        sigprocmask(SIG_SETMASK, &oset, NULL);
    }

    sigprocmask(SIG_SETMASK, &saveset, NULL);
    exit(0);
}