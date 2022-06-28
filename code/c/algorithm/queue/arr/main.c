#include <stdio.h>
#include "queue.h"

int main()
{
    queue *q;
    datatype arr[] = {2, 34, 15, 22};

    q = qu_create();
    if (q == NULL)
        return 1;

    for (int i = 0; i < sizeof(arr) / sizeof(*arr); i++)
        qu_enqueue(q, &arr[i]);

    qu_travel(q);

    datatype tmp;
    qu_dequeue(q, &tmp);
    printf("dqueue: %d\n", tmp);

    qu_travel(q);

    qu_destroy(q);
    return 0;
}