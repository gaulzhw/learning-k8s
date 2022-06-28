#include <stdio.h>
#include <stdlib.h>
#include <glob.h>
#include <syslog.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <fcntl.h>

#include <proto.h>

#include "medialib.h"
#include "mytbf.h"

#define PATHSIZE 1024

struct channel_context_st
{
    chnid_t chnid;
    char *desc;
    glob_t mp3glob;
    int pos;
    int fd;
    off_t offset;
    mytbf_t *tbuf;
};

static struct channel_context_st channel[MAX_CHANID + 1];

int mlib_getchnlist(struct mlib_listentry_st **result, int *resnum)
{
    int i, num = 0;
    char path[PATHSIZE];
    glob_t globres;
    struct mlib_listentry_st *ptr;
    struct channel_context_st *res;

    for (i = 0; i < MAX_CHANID + 1; i++)
    {
        channel[i].chnid = -1;
    }

    snprintf(path, PATHSIZE, "%s/*", server_conf.media_dir);

    if (glob(path, 0, NULL, &globres))
        return -1;

    ptr = malloc(sizeof(struct mlib_listentry_st) * globres.gl_pathc);
    if (ptr == NULL)
    {
        syslog(LOG_ERR, "malloc() error.");
        exit(1);
    }

    for (i = 0; i < globres.gl_pathc; i++)
    {
        // globres.gl_pathv[i] -> "/var/media/ch1"
        path2entry(globres.gl_pathv[i]);
        num++;
    }

    *result = realloc(ptr, sizeof(struct mlib_listentry_st) * num);
    if (*result == NULL)
        syslog(LOG_ERR, "realloc() failed.");

    *resnum = num;
    return 0;
}

int mlib_freechnlist(struct mlib_listentry_st *ptr)
{
    free(ptr);
    return 0;
}

static int open_next(chnid_t chnid)
{
    channel[chnid].pos++;

    if (channel[chnid].pos == channel[chnid].mp3glob.gl_pathc)
    {
        channel[chnid].pos = 0;
        return -1;
    }

    close(channel[chnid].fd);
    channel[chnid].fd = open(channel[chnid].mp3glob.gl_pathv[channel[chnid].pos], O_RDONLY);
    if (channel[chnid].fd < 0)
    {
        syslog(LOG_WARNING, "open(%s): %s", channel[chnid].mp3glob.gl_pathv[channel[chnid].pos], strerror(errno));
    }
    else
    {
        channel[chnid].offset = 0;
        return 0;
    }
}

ssize_t mlib_readchn(chnid_t chnid, void *buf, size_t size)
{
    int tbfsize, len;

    tbfsize = mytbf_fetchtoken(channel[chnid].tbf, size);

    while (1)
    {
        len = pread(channel[chnid].fd, buf, tbfsize, channel[chnid].offset);
        if (len < 0)
        {
            syslog(LOG_WARNING, "media file %s pread(): %s", channel[chnid].mp3glob.gl_pathv[channel[chnid].pos], strerror(errno));
            if (open_next(chnid) < 0)
                syslog(LOG_ERR, "channel %d: there is no successed open.", chnid);
        }
        else if (len == 0)
        {
            syslog(LOG_DEBUG, "media file %s is over.", channel[chnid].mp3glob.gl_pathv[channel[chnid].pos]);
            if (open_next(chnid) < 0)
                syslog(LOG_ERR, "channel %d: there is no successed open.", chnid);
        }
        else
        {
            channel[chnid].offset += len;
            break;
        }
    }

    if (tbfsize - len > 0)
        mytbf_returntoken(channel[chnid].tbf, tbfsize - len);

    return len;
}