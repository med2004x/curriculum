# Lesson 017 — Stacks, Queues, and Recursion

## Threads
**Primary:** C/Algorithms

**Supporting:** stack discipline; queue discipline; recursion and call frames

**Math companion:** recurrences and operation cost

**Time:** 3–4 hours. **Prerequisites:** Lesson 016.

## What you will be able to do

1. Explain the difference between LIFO and FIFO order.
2. Implement a stack and a queue using array or pointer-based data structures.
3. Explain how the call stack works during recursive function calls.
4. Describe the recurrence behind a simple recursive algorithm and interpret its cost.

## Retrieve first (no notes)

1. What is the invariant of a singly linked list?
2. Why does insertion at the head take O(1)?
3. Why is search in a list linear in the number of nodes?
4. What is a stale pointer?
5. What is meant by a cycle in a pointer chain?
6. What is the difference between an array and a pointer-chained list in terms of locality?

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

## Predict, then observe

1. Predict the output of a stack when operations are `push(1), push(2), pop()`.
2. Predict the output of a queue when operations are `enqueue(1), enqueue(2), dequeue()`.
3. Predict the recursive behavior of `factorial(5)` in terms of call depth.
4. Predict what happens if `n` is not reduced in a recursive function.
5. Predict the stack behavior if an expression is malformed and lacks a closing parenthesis.
6. Predict how the runtime handles nested function calls.

## Build from blank

1. Implement a stack using an array with `push`, `pop`, and `peek`.
2. Implement a queue using a ring buffer or pointer list.
3. Write a recursive function that computes factorial or traverses a directory tree.
4. Explain the function call stack in terms of the pushed frames.

## Failure drill

1. Cause a stack underflow by popping an empty stack.
2. Cause a queue underflow by dequeuing from an empty queue.
3. Write a recursive function where the parameter does not change and observe the failure mode.
4. Explain which invariant failed in each case.

## Transfer

Change exactly one condition: turn a stack into a queue using the same underlying array, or convert a recursive depth-first traversal into a breadth-first traversal. Explain why the discipline changes and what the new observable order becomes.

## Retrieval questions

1. What is LIFO?
2. What is FIFO?
3. Why is recursion naturally supported by the call stack?
4. What is a recurrence?
5. What is a stack underflow?
6. What is a queue overflow?

## Evidence to record

Create `lesson017/EVIDENCE.md` with exactly this structure:

```text
# Lesson 017 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Confusing stack order and queue order.
- Forgetting to check for empty/full conditions.
- Assuming recursion is always safe.
- Ignoring the stack depth of recursive functions.

## Precision checks

1. State the difference between stack and queue order in one sentence.
2. Describe the role of the function call stack in recursion.
3. Define a recurrence in one sentence.
4. Explain when recursion is a good fit and when it is risky.
5. State one case where a queue is more appropriate than a stack.

## University revision bridge

The academic model is the abstract data type, and the engineering model is the actual runtime behavior: functions are stored on a call stack, jobs are scheduled in a queue, and recursion expresses a smaller-subproblem structure with a clearly defined termination condition.

## Assigned material

- `07-data-structures-and-algorithms/lesson-03-stacks-queues/README.md`
- `07-data-structures-and-algorithms/lesson-16-recursion/README.md`
- `08-mathematics/step-02-discrete-math/lesson-05-recurrences-asymptotics/README.md`

## Mastery gate

You pass when you can explain LIFO vs FIFO, implement a small stack or queue from scratch, explain the role of the call stack in recursion, and state the recurrence behind a simple recursive process without notes.
