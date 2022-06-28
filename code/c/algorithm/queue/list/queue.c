#include "queue.h"

QUEUE *queue_create(int size)
{
    return llist_create(size);
}

int queue_en(QUEUE *q, const void *data)
{
    return llist_insert(q, data, LLIST_BACKWARD);
}

static int always_match(const void *p1, const void *p2)
{
    return 0;
}

int queue_de(QUEUE *q, void *data)
{
    return llist_fetch(q, (void *)0, always_match, data);
}

void queue_destroy(QUEUE *q)
{
    llist_destroy(q);
}