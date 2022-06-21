#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "llist.h"

#define NAMESIZE 32

struct score_st
{
    int id;
    char name[NAMESIZE];
    int math;
    int chinese;
};

static void print_s(const void *record)
{
    const struct score_st *r = record;
    printf("%d %s %d %d\n", r->id, r->name, r->math, r->chinese);
}

static int id_cmp(const void *key, const void *record)
{
    const int *k = key;
    const struct score_st *r = record;

    return (*k - r->id);
}

static int name_cmp(const void *key, const void *record)
{
    const char *k = key;
    const struct score_st *r = record;

    return strcmp(k, r->name);
}

int main()
{
    LLIST *handler;
    struct score_st tmp;
    int ret;

    handler = llist_create(sizeof(struct score_st));
    if (handler == NULL)
        return 1;

    for (int i = 0; i < 7; i++)
    {
        tmp.id = i;
        snprintf(tmp.name, NAMESIZE, "stu%d", i);
        tmp.math = rand() % 100;
        tmp.chinese = rand() % 100;
        ret = llist_insert(handler, &tmp, LLIST_FORWARD);
        if (ret)
            return 2;
    }
    llist_travel(handler, print_s);

    char *del_name = "stu6";
    ret = llist_delete(handler, del_name, name_cmp);
    if (ret)
        printf("llist_delete failed!\n");

#if 0
    int id = 3;
    struct score *data = llist_find(handler, &id, id_cmp);
    if (data == NULL)
        printf("Can not find!\n");
    else
        print_s(data);
#endif

    llist_destroy(handler);
    return 0;
}