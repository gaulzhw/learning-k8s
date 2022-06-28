#ifndef LIST_H__
#define LIST_H__

#include <stdbool.h>

typedef int datatype;

// 带头节点的单链表
typedef struct node_st {
    datatype data;
    struct node_st *next;
} list;

list *list_create();
void list_display(list *);
int list_insert_at(list *, int, datatype *);
int list_order_insert(list *, datatype *);
int list_delete_at(list *, int, datatype *);
int list_delete(list *, datatype *);
bool list_isempty(list *);
void list_destroy(list *);

#endif