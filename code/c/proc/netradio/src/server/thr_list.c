#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <pthread.h>
#include <syslog.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>

#include <proto.h>

#include "thr_list.h"
#include "server_conf.h"

static pthread_t tid_list;
static int nr_list_ent;
static struct mlib_listentry_st *list_ent;

static void *thr_list(void *p)
{
    int totalsize, i, size, ret;
    struct msg_list_st *entlistp;
    struct msg_listentry_st *entryp;

    totalsize = sizeof(chnid_t);

    for (i = 0; i < nr_list_ent; i++)
    {
        totalsize += sizeof(struct msg_listentry_st) + strlen(list_ent[i].desc);
    }

    entlistp = malloc(totalsize);
    if (entlistp == NULL)
    {
        syslog(LOG_ERR, "malloc(): %s.", strerror(errno));
        exit(1);
    }

    entlistp->chnid = LIST_CHNID;
    entryp = entlistp->entry;

    for (i = 0; i < nr_list_ent; i++)
    {
        size = sizeof(struct msg_listentry_st) + strlen(list_ent[i].desc);

        entryp->chnid = list_ent[i].chnid;
        entryp->len = htons(size);
        strcpy((char *)entryp->desc, list_ent[i].desc);
        entryp = (void *)(((char *)entryp) + size);
    }

    while (1)
    {
        ret = sendto(serversd, entlistp, totalsize, 0, (void *)&sndaddr, sizeof(sndaddr));
        if (ret < 0)
            syslog(LOG_WARNING, "sendto(serversd, entlistp...): %s", strerror(errno));
        else
            syslog(LOG_DEBUG, "send(serversd, entlistp...): success");

        sleep(1);
    }
}