/*
 * STUDENT PROBLEM SET — Lesson 016.
 * Read the lecture, then solve these tasks before opening any worked examples:
 * 1. Implement a singly linked list node and list head.
 * 2. Implement `push_front`, `search`, and `delete_value`.
 * 3. Write one example where the list has a cycle and explain why traversal never ends.
 * 4. Write one example where a freed node remains reachable and explain why that is unsafe.
 * Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
 */
/* PROBLEM STARTER — Lesson 016: Linked Lists, Invariants, and Cost
 * Your task: preserve the list invariant while inserting at the tail. release every node exactly once.
 * Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
 */
#include <stddef.h>
#include <stdlib.h>

typedef struct Node {
    int value;
    struct Node *next;
} Node;

static Node *push_back(Node *head, int value) {
    /* TODO: implement the assigned behavior. */
    return 0;
}

static void free_list(Node *head) {
    /* TODO: implement the assigned behavior. */
}

int main(void) {
    Node *head = NULL;
    head = push_back(head, 1);
    head = push_back(head, 2);
    free_list(head);
    return 0;
}
