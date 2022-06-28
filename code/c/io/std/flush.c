#include <stdio.h>
#include <stdlib.h>

/*
 * 缓冲区的作用：大多数情况下是好事，合并系统调用
 *
 * 行缓冲：换行时候刷新，满了的时候刷新，强制刷新（标准输出是这样的，因为是终端设备）
 * 全缓冲：满了的时候刷新，强制刷新（默认，只要不是终端设备）
 * 无缓冲：如stderr，需要立即输出的内容
 *
 * setvbuf
 */

int main()
{
    // print是按照行缓存输出，如果去掉\n不会打印
    // printf("Before while()\n");

    printf("Before while()");
    fflush(NULL);

    while(1);

    printf("After while()\n");

    return 0;
}