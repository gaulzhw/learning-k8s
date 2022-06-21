#ifndef LLIST_H__
#define LLIST_H__

#define LLIST_FORWARD 1
#define LLIST_BACKWARD 2

typedef void llist_op(const void *);
typedef int llist_cmp(const void *, const void *);

struct llist_node_st
{
    struct llist_node_st *prev;
    struct llist_node_st *next;
    char data[0];
};

typedef struct llist_head
{
    int size;
    struct llist_node_st head;

    void (*travel)(struct llist_head *, llist_op *);
    int (*insert)(struct llist_head *, const void *, int); //int mode
    void *(*find)(struct llist_head *, const void *, llist_cmp *);
    int (*delete)(struct llist_head *, const void *, llist_cmp *);
    int (*fetch)(struct llist_head *, const void *, llist_cmp *, void *);
} LLIST;

LLIST *llist_create(int); //带头节点的空的双向链表
void llist_destroy(LLIST *);

#endif