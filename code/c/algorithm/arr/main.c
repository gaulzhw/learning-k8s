#include <stdio.h>
#include "sqlist.h"

int main()
{
    sqlist *list = NULL, *list1 = NULL;
    datatype arr[] = {12, 23, 34, 45, 56};
    datatype arr1[] = {89, 90, 78, 67, 56, 45};
    int i;

    list = sqlist_create();
    if (list == NULL)
    {
        fprintf(stderr, "sqlist_create() failed!\n");
        return 1;
    }

    list1 = sqlist_create();
    if (list1 == NULL)
    {
        fprintf(stderr, "sqllist_create() failed!\n");
        return 1;
    }

    // printf("%d\n", __LINE__);

    for (i = 0; i < sizeof(arr) / sizeof(*arr); i++)
    {
        if (sqlist_insert(list, 0, &arr[i]) != 0)
        {
            fprintf(stderr, "insert list error\n");
            return 2;
        }
    }

    for (i = 0; i < sizeof(arr1) / sizeof(*arr1); i++)
    {
        if (sqlist_insert(list1, 0, &arr1[i]) != 0)
        {
            fprintf(stderr, "insert list error\n");
            return 2;
        }
    }

#if 0
    sqlist_display(list);
    sqlist_delete(list, 1);
    sqlist_display(list);
#endif

    sqlist_union(list, list1);
    sqlist_display(list);

    sqlist_destroy(list);
    sqlist_destroy(list1);
    return 0;
}