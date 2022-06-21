#include <stdio.h>
#include <stdlib.h>
#include "queue.h"

queue *qu_create(void)
{
    queue *q;
    q = malloc(sizeof(queue));
    if (q == NULL)
        return NULL;

    q->head = 0;
    q->tail = 0;
    return q;
}

int qu_isempty(queue *q)
{
    return q->head == q->tail;
}

int qu_enqueue(queue *q, datatype *data)
{
    if ((q->tail + 1) % MAXSIZE == q->head)
        return -1;

    q->tail = (q->tail + 1) % MAXSIZE;
    q->data[q->tail] = *data;
    return 0;
}

int qu_dequeue(queue *q, datatype *data)
{
    if (qu_isempty(q))
        return -1;

    q->head = (q->head + 1) % MAXSIZE;
    *data = q->data[q->head];
    return 0;
}

void qu_travel(queue *q)
{
    if (qu_isempty(q))
        return;

    int i = (q->head + 1) % MAXSIZE;
    while (i != q->tail)
    {
        printf("%d ", q->data[i]);
        i = (i + 1) % MAXSIZE;
    }
    printf("%d\n", q->data[i]);
}

void qu_clear(queue *q)
{
    q->head = q->tail;
}

void qu_destroy(queue *q)
{
    free(q);
}