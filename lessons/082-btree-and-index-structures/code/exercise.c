#include <stddef.h>

typedef struct Node {
    int keys[8];
    size_t count;
    struct Node *child[9];
} Node;

static int insert(Node *root, int key) {
    /* TODO: preserve ordering and split behavior when a node becomes full. */
    (void)root;
    (void)key;
    return 0;
}

int main(void) {
    Node root = {0};
    return insert(&root, 42);
}
