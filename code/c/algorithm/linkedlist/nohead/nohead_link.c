#include <stdio.h>
#include <stdlib.h>

#define NAMESIZE 32

struct score_st
{
    int id;
    char name[NAMESIZE];
    int math;
    int chinese;
};

struct node_st
{
    struct score_st data;
    struct node_st *next;
};

int list_insert(struct node_st **list, struct score_st *data)
{
    struct node_st *node;

    node = malloc(sizeof(struct node_st));
    if (node == NULL)
        return -1;

    node->data = *data;
    // node->next = NULL;

    node->next = *list;
    *list = node;

    return 0;
}

void list_show(struct node_st *list)
{
    struct node_st *cur;

    for (cur = list; cur != NULL; cur = cur->next)
        printf("%d %s %d %d\n", cur->data.id, cur->data.name, cur->data.math, cur->data.chinese);
}

int list_delete(struct node_st **list)
{
    struct node_st *cur;

    if (*list == NULL)
        return -1;

    cur = *list;
    *list = (*list)->next;

    free(cur);
    return 0;
}

int list_find(struct node_st *list, int id)
{
    struct node_st *cur;

    for (cur = list; cur != NULL; cur = cur->next)
        if (cur->data.id == id)
        {
            printf("%d %s %d %d\n", cur->data.id, cur->data.name, cur->data.math, cur->data.chinese);
            return 0;
        }

    return -1;
}

void list_destroy(struct node_st *list)
{
    struct node_st *cur;

    if (list == NULL)
        return;

    for (cur = list; cur != NULL; cur = list)
    {
        list = cur->next;
        free(cur);
    }
}

int main()
{
    struct node_st *list = NULL;
    struct score_st tmp;

    for (int i = 0; i < 7; i++)
    {
        tmp.id = i;
        snprintf(tmp.name, NAMESIZE, "stu%d", i);
        tmp.math = rand() % 100;
        tmp.chinese = rand() % 100;

        list_insert(&list, &tmp);
    }

    list_show(list);

    list_delete(&list);
    list_show(list);

    list_find(list, 3);

    list_destroy(list);
    return 0;
}