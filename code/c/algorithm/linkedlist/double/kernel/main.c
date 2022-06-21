#include <stdio.h>
#include <stdlib.h>
#include "list.h"

#define NAMESIZE 32

struct score_st
{
    int id;
    char name[NAMESIZE];
    int math;
    int chinese;
    struct list_head node;
};

static void print_s(struct score_st *d)
{
    printf("%d %s %d %d\n", d->id, d->name, d->math, d->chinese);
}

int main()
{
    struct score_st *data;
    struct list_head *cur;

    LIST_HEAD(head);

    for (int i = 0; i < 7; i++)
    {
        data = malloc(sizeof(struct score_st));
        if (data == NULL)
            return 1;
        
        data->id = i;
        snprintf(data->name, NAMESIZE, "stu%d", i);
        data->math = rand()%100;
        data->chinese = rand()%100;

        list_add(&data->node, &head);
    }

    __list_for_each(cur, &head)
    {
        data = list_entry(cur, struct score_st, node);
        print_s(data);
    }

    return 0;
}