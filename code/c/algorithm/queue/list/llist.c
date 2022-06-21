#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "llist.h"

LLIST *llist_create(int initSize)
{
    LLIST *new;
    new = malloc(sizeof(LLIST));
    if (new == NULL)
        return NULL;

    new->size = initSize;
    new->head.prev = &new->head;
    new->head.next = &new->head;

    return new;
}

void llist_travel(LLIST *list, llist_op *op)
{
    struct llist_node_st *cur;

    for (cur = list->head.next; cur != &list->head; cur = cur->next)
        op(cur->data);
}

void llist_destroy(LLIST *list)
{
    struct llist_node_st *cur, *next;

    for (cur = list->head.next; cur != &list->head; cur = next)
    {
        next = cur->next;
        free(cur);
    }

    free(list);
}

int llist_insert(LLIST *list, const void *data, int mode)
{
    struct llist_node_st *newnode;

    newnode = malloc(sizeof(struct llist_node_st) + list->size);
    if (newnode == NULL)
        return -1;

    memcpy(newnode->data, data, list->size);

    if (mode == LLIST_FORWARD)
    {
        newnode->prev = &list->head;
        newnode->next = list->head.next;
    }
    else if (mode == LLIST_BACKWARD)
    {
        newnode->prev = list->head.prev;
        newnode->next = &list->head;
    }
    else
    {
        return -3;
    }

    newnode->prev->next = newnode;
    newnode->next->prev = newnode;
    return 0;
}

static struct llist_node_st *_find(LLIST *list, const void *key, llist_cmp *cmp)
{
    struct llist_node_st *cur;

    for (cur = list->head.next; cur != &list->head; cur = cur->next)
    {
        if (cmp(key, cur->data) == 0)
            break;
    }

    return cur;
}

void *llist_find(LLIST *list, const void *key, llist_cmp *cmp)
{
    struct llist_node_st *node = _find(list, key, cmp);

    if (node == &list->head)
        return NULL;

    return node->data;
}

int llist_delete(LLIST *list, const void *key, llist_cmp *cmp)
{
    struct llist_node_st *node = _find(list, key, cmp);
    if (node == &list->head)
        return -1;

    node->prev->next = node->next;
    node->next->prev = node->prev;
    free(node);
    return 0;
}

int llist_fetch(LLIST *list, const void *key, llist_cmp *cmp, void *data)
{
    struct llist_node_st *node = _find(list, key, cmp);
    if (node == &list->head)
        return -1;

    node->prev->next = node->next;
    node->next->prev = node->prev;
    if (data != NULL)
        memcpy(data, node->data, list->size);

    free(node);
    return 0;
}