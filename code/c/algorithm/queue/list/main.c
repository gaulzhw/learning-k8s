#include <stdio.h>
#include <stdlib.h>
#include "queue.h"

#define NAMESIZE 32

struct score_st
{
    int id;
    char name[NAMESIZE];
    int math;
    int chinese;
};

static void print_s(void *record)
{
    struct score_st *r = record;
    printf("%d %s %d %d\n", r->id, r->name, r->math, r->chinese);
}

int main()
{
    QUEUE *q;
    struct score_st tmp;

    q = queue_create(sizeof(struct score_st));
    if (q == NULL)
        return 1;

    for (int i = 0; i < 7; i++)
    {
        tmp.id = i;
        snprintf(tmp.name, NAMESIZE, "stu%d", i);
        tmp.math = rand() % 100;
        tmp.chinese = rand() % 100;

        if (queue_en(q, &tmp) != 0)
            break;
    }

    while(1)
    {
        if (queue_de(q, &tmp) != 0)
            break;
        
        print_s(&tmp);
    }

    queue_destroy(q);
    return 0;
}