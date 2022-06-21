#ifndef PROTO_H__
#define PROTO_H__

#include <site_type.h>

#define DEFAULT_MGROUP "224.2.2.2"
#define DEFAULT_RCVPORT "1989"

#define CHNNR 100

#define LIST_CHNID 0
#define MIN_CHANID 1
#define MAX_CHANID (MIN_CHANID + CHANNR - 1)

#define MSG_CHANNEL_MAX (65536 - 20 - 8)
#define MAX_DATA (MSG_CHANNEL_MAX - sizeof(chnid_t))

#define MSG_LIST_MAX (65536 - 20 - 8)
#define MAX_ENTRY (MSG_LIST_MAX - sizeof(chnid_t))

struct msg_channel_st
{
    chnid_t chnid; // must between [MIN_CHANID, MAX_CHANID]
    uint8_t data[1];
} __attribute__((packed));

struct msg_listentry_st
{
    chnid_t chnid;
    uint16_t len;
    uint8_t desc[1];
} __attribute__((packed));

struct msg_list_st
{
    chnid_t chnid; // must be LIST_CHNID
    struct msg_listentry_st entry[1];
} __attribute__((packed));

#endif