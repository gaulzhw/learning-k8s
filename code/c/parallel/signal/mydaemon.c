#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <syslog.h>
#include <string.h>
#include <errno.h>
#include <signal.h>

#define FNAME "/tmp/out"

static FILE *fp;

static int daemonize(void)
{
    int fd;

    pid_t pid = fork();
    if (pid < 0)
        return -1;

    if (pid > 0)
        exit(0);

    fd = open("/dev/null", O_RDWR);
    if (fd < 0)
        return -1;

    dup2(fd, 0);
    dup2(fd, 1);
    dup2(fd, 2);

    if (fd > 2)
        close(fd);

    setsid();

    chdir("/");
    umask(0); // 如果不产生文件，可以关掉umask

    return 0;
}

// 参数表示信号值 SIGINT SIGQUIT SIGTERM
static void daemon_exit(int s)
{
    fclose(fp);
    closelog();
    exit(0);
}

int main()
{
    struct sigaction sa;

    sa.__sigaction_u.__sa_handler = daemon_exit;
    sigemptyset(&sa.sa_mask);
    sigaddset(&sa.sa_mask, SIGQUIT);
    sigaddset(&sa.sa_mask, SIGTERM);
    sigaddset(&sa.sa_mask, SIGINT);
    sa.sa_flags = 0;

    sigaction(SIGINT, &sa, NULL);
    sigaction(SIGQUIT, &sa, NULL);
    sigaction(SIGTERM, &sa, NULL);

    // signal(SIGINT, daemon_exit);
    // signal(SIGQUIT, daemon_exit);
    // signal(SIGTERM, daemon_exit);

    openlog("mydaemon", LOG_PID, LOG_DAEMON);

    if (daemonize())
    {
        syslog(LOG_ERR, "daemonize() failed");
        exit(1);
    }
    else
    {
        syslog(LOG_INFO, "daemonize() successded");
    }

    fp = fopen(FNAME, "w");
    if (fp == NULL)
    {
        syslog(LOG_ERR, "fopen: %s", strerror(errno));
        exit(1);
    }

    syslog(LOG_INFO, "%s was opend.", FNAME);

    for (int i = 0;; i++)
    {
        fprintf(fp, "%d", i);
        fflush(fp);
        syslog(LOG_DEBUG, "%d is printed.", i);
        sleep(1);
    }

    exit(0);
}