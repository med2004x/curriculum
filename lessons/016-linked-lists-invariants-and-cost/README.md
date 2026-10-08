# Lesson 016 — Linked Lists, Invariants, and Cost

## Threads
**Primary:** C/Algorithms

**Supporting:** linked lists; pointer invariants; dynamic memory

**Math companion:** asymptotic cost and amortized reasoning

**Time:** 3–4 hours. **Prerequisites:** Lesson 015.

## What you will be able to do

1. Define the invariant of a singly linked list and explain what a valid node chain looks like.
2. Explain why insertion at the head is O(1) while searching for a value is O(n).
3. Describe the difference between list invariants and data structure performance.
4. Identify the common failure modes of linked lists: stale pointers, lost head references, and cycles.

## Retrieve first (no notes)

1. What is undefined behavior in C, in one sentence?
2. Why is a null pointer dangerous?
3. What does a sanitizer do?
4. What is the difference between a valid pointer and a dangling pointer?
5. Why can a malloc’d object become invalid?
6. What does “invariant” mean in data structures?

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

## Predict, then observe

1. Predict what happens if you remove the `next` assignment in `push_front`.
2. Predict what happens if the list has a cycle.
3. Predict which operation is cheaper: searching in an array or searching in a linked list.
4. Predict whether insertion at the head costs more or less than insertion at the tail in a singly linked list with only a head pointer.
5. Predict what happens when you free a node but leave a pointer to it in another list element.
6. Predict which invariant fails in a cycle.

## Build from blank

1. Implement a singly linked list node and list head.
2. Implement `push_front`, `search`, and `delete_value`.
3. Write one example where the list has a cycle and explain why traversal never ends.
4. Write one example where a freed node remains reachable and explain why that is unsafe.

## Failure drill

1. Create a list with five nodes and then remove the head.
2. Create a cycle by making the last node point to the second node.
3. Free a node while another node still points to it.
4. For each case, state which invariant was violated and what evidence indicates the violation.

## Transfer

Change exactly one condition: replace a singly linked list with a doubly linked list, or insert at the tail instead of the head. Explain which costs change and which invariants remain the same.

## Retrieval questions

1. What is the meaning of an invariant in a linked list?
2. Why is insertion at the head O(1)?
3. Why is searching in a singly linked list O(n)?
4. What is a cycle and why is it a bug?
5. Why do linked-list operations have poor cache locality compared with arrays?
6. What is the usual condition for a valid traversal to end?

## Evidence to record

Create `lesson016/EVIDENCE.md` with exactly this structure:

```text
# Lesson 016 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Forgetting to update the head pointer.
- Losing the tail pointer in a singly linked list.
- Forgetting to set `next` to `NULL` on the new node.
- Failing to consider cycles and stale pointers.
- Confusing list representation with cost of a single operation.

## Precision checks

1. State the invariant of a singly linked list in one sentence.
2. Explain why a cycle breaks the termination condition.
3. State the cost of `push_front` and `search` in big-O notation.
4. Explain the difference between a valid pointer and a dangling pointer.
5. Describe one case where arrays are better and one case where lists are better.

## University revision bridge

Academically, a linked list is a model for how a representation changes the algorithmic cost of operations. Practically, the list invariant is what keeps the pointer structure meaningful; without the invariant, every traversal becomes a speculative guess.

## Assigned material

- `07-data-structures-and-algorithms/lesson-02-arrays-linked-lists/README.md`
- `07-data-structures-and-algorithms/lesson-01-complexity/README.md`

## Mastery gate

You pass when you can state the list invariant, explain why insertion at the head is O(1), and identify a cycle or stale-pointer bug without notes. You should be able to tell why a linked list and an array have different trade-offs even when they store the same data.
