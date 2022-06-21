#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

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
    struct score_st tmp;

    STACK *st = stack_create(sizeof(struct score_st));
    if (st == NULL)
        return -1;

    for (int i = 0; i < 7; i++)
    {
        tmp.id = i;
        snprintf(tmp.name, NAMESIZE, "stu%d", i);
        tmp.math = rand() % 100;
        tmp.chinese = rand() % 100;

        if (stack_push(st, &tmp))
            return 1;
    }

    while (1)
    {
        int ret = stack_pop(st, &tmp);
        if (ret != 0)
            break;
        print_s(&tmp);
    }

    stack_destroy(st);
    return 0;
}