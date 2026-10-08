# Lesson 017 — Stacks, Queues, and Recursion

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

A stack is a container with a single access point: the top. Operations are:

- `push`: add an item to the top
- `pop`: remove the item at the top
- `peek`: look at the top item without removing it

A queue has two ends:

- `enqueue`: add to the back
- `dequeue`: remove from the front

The order is different:

- stack: last in, first out (LIFO)
- queue: first in, first out (FIFO)

This difference matters because a stack models nested calls, undo operations, expression evaluation, and depth-first traversal. A queue models work scheduling, buffering, and breadth-first traversal.

## 2. Mechanism and vocabulary

The call stack is a stack used by the runtime. Each function call pushes a new frame containing local variables, return address, and stack state. When the function returns, that frame is popped. This is why recursion is naturally managed as a stack discipline.

Recursion is also the vocabulary behind recurrences. A recurrence defines a function in terms of smaller subproblems. For example:

```text
T(n) = T(n/2) + 1
```

This describes binary search: reduce the problem by half, do constant work, then recurse. In general, the recurrence and the structure of the algorithm are linked:

- stack-like work corresponds to depth-first exploration
- queue-like work corresponds to breadth-first exploration
- recursion depth must be bounded or the stack may overflow

## 3. Worked case: expression evaluation

A simple expression like:

```text
(3 + 4) * 2
```

can be evaluated with a stack by converting to postfix or by an explicit parser. In postfix form, the expression is: `3 4 + 2 *`.

The algorithm is:

1. read `3`, push it
2. read `4`, push it
3. read `+`, pop `4` and `3`, compute `7`, push `7`
4. read `2`, push it
5. read `*`, pop `2` and `7`, compute `14`, push it
6. result is `14`

This is a concrete example of a stack preserving the correct evaluation order.

## 4. Verification, failure, and evidence

A queue should preserve FIFO order. A stack should preserve LIFO order. A recursive function should not keep recursing until the stack overflows unless that is the intended termination condition.

The evidence you record should include:

- the exact push/pop or enqueue/dequeue sequence
- the expected order of values
- the actual observed order
- the explanation of the bug if the order differs

Common failures:

- using a queue for a stack operation and reversing the expected order
- allowing an empty queue to dequeue
- allowing a full stack to push
- recursive calls that do not reduce the problem size

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
