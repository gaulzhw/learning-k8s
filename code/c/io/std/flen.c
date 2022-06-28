#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

int main()
{
    FILE *fp;
    fp = fopen("tmp", "r");

    if (fp == NULL)
    {
        perror("fopen()");
        return 1;
    }

    fseek(fp, 0, SEEK_END);
    printf("%ld\n", ftell(fp));

    return 0;
}