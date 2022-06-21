#include <stdio.h>
#include <stdlib.h>
#include <sys/ipc.h>
#include <sys/types.h>
#include <sys/sem.h>
#include <unistd.h>
#include <errno.h>

#define PROCNUM 20
#define FNAME "/tmp/out"
#define LINESIZE 1024

static int semid;

static void P(void)
{
    struct sembuf op;
    op.sem_num = 0;
    op.sem_op = -1;
    op.sem_flg = 0;

    while (semop(semid, &op, 1) < 0)
    {
        if (errno != EINTR || errno != EAGAIN)
        {
            perror("semop()");
            exit(1);
        }
    }
}

static void V(void)
{
    struct sembuf op;
    op.sem_num = 0;
    op.sem_op = 1;
    op.sem_flg = 0;

    if (semop(semid, &op, 1) < 0)
    {
        perror("semop()");
        exit(1);
    }
}

static void func_add(void)
{
    FILE *fp;
    char linebuf[LINESIZE];

    fp = fopen(FNAME, "r+");
    if (fp == NULL)
    {
        perror("fopen()");
        exit(1);
    }

    P();
    fgets(linebuf, LINESIZE, fp);
    fseek(fp, 0, SEEK_SET);
    fprintf(fp, "%d\n", atoi(linebuf) + 1);
    fflush(fp);

    V();
    fclose(fp);

    return;
}

int main()
{
    pid_t pid;

    semid = semget(IPC_PRIVATE, 1, 0600);
    if (semid < 0)
    {
        perror("semget()");
        exit(1);
    }

    if (semctl(semid, 0, SETVAL, 1) < 0)
    {
        perror("semctl()");
        exit(1);
    }

    for (int i = 0; i < PROCNUM; i++)
    {
        pid = fork();
        if (pid < 0)
        {
            perror("fork()");
            exit(1);
        }
        if (pid == 0)
        {
            func_add();
            exit(0);
        }
    }

    for (int i = 0; i < PROCNUM; i++)
        wait(NULL);

    semctl(semid, 0, IPC_RMID);
    exit(0);
}