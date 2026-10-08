/*
 * STUDENT PROBLEM SET — Lesson 082.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Build the smallest useful version: Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page).
 * 2. Extend it with one explicit boundary case: Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.
 * 3. Rebuild the same behavior with one changed constraint: Insert keys in increasing order (1, 2, 3, 4, 5, ...) and trace, at each insert, whether a split occurred and at which level. Do not open the reference implementation until your own attempt has been saved.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 082: B+ Trees and Index Structures
 * Your task: preserve ordering and split behavior when a node becomes full.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stddef.h>

typedef struct Node {
    int keys[8];
    size_t count;
    struct Node *child[9];
} Node;

static int insert(Node *root, int key) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

int main(void) {
    Node root = {0};
    return insert(&root, 42);
}
