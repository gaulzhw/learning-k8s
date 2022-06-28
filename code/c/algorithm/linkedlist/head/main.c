#include <stdio.h>
#include "list.h"

int main()
{
    list *l;
    int i;
    datatype arr[] = {12, 9, 23, 2, 34, 6, 45};

    l = list_create();
    if (l == NULL)
        return 1;

    for (i = 0; i < sizeof(arr) / sizeof(*arr); i++)
    {
        if (list_order_insert(l, &arr[i]))
            return 1;
    }

    list_display(l);

    int value = 12;
    list_delete(l, &value);

    list_display(l);

    list_destroy(l);
    return 0;
}