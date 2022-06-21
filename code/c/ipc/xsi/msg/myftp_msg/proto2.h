#ifndef PROTO_H__
#define PROTO_H__

#define KEYPATH "/etc/services"
#define KEYPROJ 'a'

#define PATHMAX 1024
#define DATAMAX 1024

enum
{
    MSG_PATH = 1,
    MSG_DATA,
    MSG_EOT
};

struct msg_path_st
{
    long mtype;         // must be MSG_PATH
    char path[PATHMAX]; // ASCII 带尾0的串
};

struct msg_s2c_st
{
    long mtype; // must be MSG_DATA or MSG_EOT
    int datalen;
    // datalen > 0: data
    // datalen == 0: eot
    char data[DATAMAX];
};

#endif