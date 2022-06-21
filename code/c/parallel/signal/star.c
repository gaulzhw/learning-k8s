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
    // signal(SIGINT, SIG_IGN);
    signal(SIGINT, int_handler);

    for (int i = 0; i < 10; i++)
    {
        write(1, "*", 1);
        sleep(1);
    }

    exit(0);
}