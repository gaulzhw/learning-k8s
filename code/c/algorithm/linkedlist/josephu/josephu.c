#include <stdio.h>
#include <stdlib.h>

struct node_st
{
    int data;
    struct node_st *next;
};

struct node_st *jose_create(int num)
{
    struct node_st *list;
    struct node_st *newnode, *cur;
    int i = 1;

    list = malloc(sizeof(struct node_st));
    if (list == NULL)
        return NULL;

    list->data = i;
    list->next = list;
    i++;

    cur = list;
    for (; i <= num; i++)
    {
        newnode = malloc(sizeof(struct node_st));
        if (newnode == NULL)
            return NULL;

        newnode->data = i;
        newnode->next = list;
        cur->next = newnode;
        cur = newnode;
    }
    return list;
}

void jose_show(struct node_st *list)
{
    struct node_st *me;

    for (me = list; me->next != list; me = me->next)
        printf("%d", me->data);

    printf("%d\n", me->data);
}

void jose_kill(struct node_st **list, int n)
{
    struct node_st *cur = *list, *node;
    int i = 1;

    while (cur != cur->next)
    {
        while (i < n)
        {
            node = cur;
            cur = cur->next;
            i++;
        }

        node->next = cur->next;
        free(cur);

        cur = node->next;
        i = 1;
    }

    *list = cur;
}

int main()
{
    struct node_st *list = jose_create(8);
    jose_show(list);

    jose_kill(&list, 3);
    jose_show(list);

    return 0;
}