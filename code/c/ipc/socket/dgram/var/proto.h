#ifndef PROTO_H__
#define PROTO_H__

#include <stdlib.h>

#define RCVPORT "1989"

// 512: udp推荐长度, 8: udp报头大小, 8: struct剩余字段大小
#define NAMEMAX (512-8-8)

struct msg_st
{
    uint32_t math;
    uint32_t chinese;
    uint8_t name[0];
} __attribute__((packed));

#endif