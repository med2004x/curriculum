#include <stddef.h>
#include <stdlib.h>

typedef struct Node {
    int value;
    struct Node *next;
} Node;

static Node *push_back(Node *head, int value) {
    /* TODO: preserve the list invariant while inserting at the tail. */
    (void)value;
    return head;
}

static void free_list(Node *head) {
    /* TODO: release every node exactly once. */
    (void)head;
}

int main(void) {
    Node *head = NULL;
    head = push_back(head, 1);
    head = push_back(head, 2);
    free_list(head);
    return 0;
}
