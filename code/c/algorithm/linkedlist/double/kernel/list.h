#ifndef LINUX_LIST_H__
#define LINUX_LIST_H__

#define LIST_HEAD_INIT(name) \
    {                        \
        &(name), &(name)     \
    }

#define LIST_HEAD(name) struct list_head name = LIST_HEAD_INIT(name)

#define __list_for_each(pos, head) for (pos = (head)->next; pos != (head); pos = pos->next)

#define offsetof(TYPE, MEMBER) ((size_t) &((TYPE *)0)->MEMBER)

#define container_of(ptr, type, member) ({ \
    const typeof( ((type *)0)->member ) *__mptr = (ptr); \
    (type *)( (char *)__mptr - offsetof(type, member) );})

#define list_entry(ptr, type, member) container_of(ptr, type, member)

struct list_head
{
    struct list_head *prev;
    struct list_head *next;
};

static inline void __list_add(struct list_head *list, struct list_head *prev, struct list_head *next)
{
    next->prev = list;
    list->next = next;
    list->prev = prev;
    prev->next = list;
}

static inline void list_add(struct list_head *list, struct list_head *head)
{
    __list_add(list, head, head->next);
}

static inline void list_add_tail(struct list_head *list, struct list_head *head)
{
    __list_add(list, head->prev, head);
}

#endif