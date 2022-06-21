#include <stdio.h>
#include <stdlib.h>
#include "sqlist.h"

sqlist *sqlist_create()
{
    sqlist *list;
    list = malloc(sizeof(sqlist));
    if (list == NULL)
        return NULL;

    list->last = -1;
    return list;
}

/*
sqlist *sqlist_create(sqlist **ptr) {
    *ptr = malloc(sizeof(sqlist));
    if (*ptr == NULL)
        return;

    (*ptr)->last = -1;
    return;
}
*/

void sqlist_destroy(sqlist *list)
{
    free(list);
}

int sqlist_insert(sqlist *list, int i, datatype *data)
{
    int j;

    if (list->last == DATASIZE - 1)
        return -1; // full

    if (i < 0 || i > list->last + 1)
        return -2;

    for (j = list->last; i <= j; j--)
        list->data[j + 1] = list->data[j];

    list->data[i] = *data;
    list->last++;
    return 0;
}

int sqlist_delete(sqlist *list, int i)
{
    int j;
    if (i < 0 || i > list->last)
        return -1;

    for (j = i + 1; j <= list->last; j++)
        list->data[j - 1] = list->data[j];

    list->last--;
    return 0;
}

int sqlist_find(sqlist *list, datatype *data)
{
    int i;
    if (sqlist_isempty(list))
        return -1;

    for (i = 0; i < list->last; i++)
    {
        if (list->data[i] == *data)
            return i;
    }
    return -2;
}

bool sqlist_isempty(sqlist *list)
{
    if (list->last == -1)
        return true;

    return false;
}

int sqlist_setempty(sqlist *list)
{
    list->last = -1;
    return 0;
}

int sqlist_getnum(sqlist *list)
{
    return list->last + 1;
}

void sqlist_display(sqlist *list)
{
    int i;
    if (list->last == -1)
        return;

    for (i = 0; i <= list->last; i++)
        printf("%d ", list->data[i]);

    printf("\n");
    return;
}

int sqlist_union(sqlist *list1, sqlist *list2)
{
    int i;

    for (i = 0; i <= list2->last; i++)
    {
        if (sqlist_find(list1, &list2->data[i]) < 0)
            sqlist_insert(list1, 0, &list2->data[i]);
    }

    return 0;
}