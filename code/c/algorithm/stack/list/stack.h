#ifndef STACK_H__
#define STACK_H__

#include "llist.h"

typedef LLIST STACK;

STACK *stack_create(int);

int stack_push(STACK *, const void *);

int stack_pop(STACK *, void *);

void stack_destroy(STACK *);

#endif