#ifndef QUEUE_H__
#define QUEUE_H__

#include "llist.h"

typedef LLIST QUEUE;

QUEUE *queue_create(int);

int queue_en(QUEUE *, const void *);

int queue_de(QUEUE *, void *);

void queue_destroy(QUEUE *);

#endif