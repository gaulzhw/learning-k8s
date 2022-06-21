#ifndef SQLIST_H__
#define SQLIST_H__

#include <stdbool.h>
#define DATASIZE 1024

typedef int datatype;

typedef struct node_st
{
    datatype data[DATASIZE];
    int last; // 表示当前已经赋值的索引位置，初始值为-1
} sqlist;

sqlist *sqlist_create();
// void sqlist_create(sqlist **);

void sqlist_destroy(sqlist *);

int sqlist_insert(sqlist *, int, datatype *);

int sqlist_delete(sqlist *, int);

int sqlist_find(sqlist *, datatype *);

bool sqlist_isempty(sqlist *);

int sqlist_setempty(sqlist *);

int sqlist_getnum(sqlist *);

void sqlist_display(sqlist *);

int sqlist_union(sqlist *, sqlist *);

#endif