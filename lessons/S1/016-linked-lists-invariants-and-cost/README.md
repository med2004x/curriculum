# Lesson 016 — Linked Lists, Invariants, and Cost

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

A linked list stores values in nodes, where each node contains a value and a pointer to the next node. The key idea is that the structure is not stored in one contiguous block of memory. Instead, the nodes are scattered across the heap and linked by addresses.

This gives a flexible representation for insertion and deletion, but it also creates new responsibilities:

- every node must be reachable from the list head
- the final node must have `next == NULL`
- a node must not be inserted twice into the same list
- a freed node must not remain reachable from the list

A singly linked list is a good example of a data structure with a clear invariant: if you follow `next` pointers from the head, you eventually reach `NULL`, and each node is visited exactly once in that traversal.

## 2. Mechanism and vocabulary

The critical operations are:

- `push_front`: insert a node at the head in O(1)
- `search`: scan from the head to the target in O(n)
- `remove`: remove a node by pointer or by key, typically O(n) unless the node is known directly
- `append`: add to the tail in O(n) if you only keep a head pointer

The when/why:

- arrays are good when random access matters
- linked lists are good when insertion and deletion at the front matter
- linked lists are poor for cache locality because nodes are not contiguous

The performance terms are:

- O(1): constant time
- O(n): linear time
- O(n^2): quadratic time in an outer scan plus inner scan
- amortized: average cost over many operations

## 3. Worked case: head insertion and list invariant

A small list can be built as:

```c
struct node {
    int value;
    struct node *next;
};
```

Insertion at the head is simple:

```c
void push_front(struct node **head, int value) {
    struct node *n = malloc(sizeof(*n));
    if (!n) return;
    n->value = value;
    n->next = *head;
    *head = n;
}
```

The invariant is preserved because the new node becomes the first element, and its `next` points to the old head. The list still ends in `NULL` when the old tail was valid.

A search loop should look like:

```c
bool contains(struct node *head, int target) {
    for (struct node *p = head; p != NULL; p = p->next) {
        if (p->value == target) return true;
    }
    return false;
}
```

## 4. Verification, failure, and evidence

To verify the data structure, record:

- the initial list state
- the operation performed
- the before/after pointer chain
- the resulting value found or not found

Typical failure modes:

- stale pointer: after freeing a node, some other node still points to it
- cycle: a node points back to an earlier node, so traversal never reaches `NULL`
- lost head: you assign a new head but forget to preserve the old chain
- double-free: the same allocated node is freed twice

These are not “algorithmic mistakes” only; they are representation errors that violate the list invariant.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
