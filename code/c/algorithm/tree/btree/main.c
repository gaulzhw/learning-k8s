#include <stdio.h>
#include <stdlib.h>

#define NAMESIZE 32

struct score_st
{
    int id;
    char name[NAMESIZE];
    int math;
    int chinese;
};

struct node_st
{
    struct score_st data;
    struct node_st *l, *r;
};

int insert(struct node_st **root, struct score_st *data)
{
    struct node_st *node;

    if (*root == NULL)
    {
        node = malloc(sizeof(*node));
        if (node == NULL)
            return -1;

        node->data = *data;
        node->l = NULL;
        node->r = NULL;

        *root = node;
        return 0;
    }

    if (data->id <= (*root)->data.id)
        return insert(&(*root)->l, data);

    return insert(&(*root)->r, data);
}

struct score_st *find(struct node_st *root, int id)
{
    if (root == NULL)
        return NULL;

    if (id == root->data.id)
        return &root->data;

    if (id < root->data.id)
        return find(root->l, id);

    return find(root->r, id);
}

int main()
{
    int arr[] = {1, 2, 3, 7, 6, 5, 9, 8, 4};
    struct node_st *tree = NULL;
    struct score_st tmp;

    for (int i = 0; i < sizeof(arr) / sizeof(*arr); i++)
    {
        tmp.id = arr[i];
        snprintf(tmp.name, NAMESIZE, "stu%d", arr[i]);
        tmp.math = rand() % 100;
        tmp.chinese = rand() % 100;

        insert(&tree, &tmp);
    }

    int tmpid = 2;
    struct score_st *datap = find(tree, tmpid);
    if (datap == NULL)
        printf("Can not find the id %d\n", tmpid);
    else
        printf("%d %s %d %d\n", datap->id, datap->name, datap->math, datap->chinese);

    return 0;
}