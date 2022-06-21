#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>

static int ftype(const char *fname)
{
    struct stat result;
    if (stat(fname, &result) < 0)
    {
        perror("stat()");
        return -1;
    }

    if (S_ISREG(result.st_mode))
        return '-';
    else if (S_ISDIR(result.st_mode))
        return 'd';
    else if (S_ISSOCK(result.st_mode))
        return 's';
    else
        return '?';
}

int main(int argc, char **argv)
{
    if (argc < 2)
    {
        fprintf(stderr, "Usage ...\n");
        return 1;
    }

    printf("%c\n", ftype(argv[1]));

    return 0;
}