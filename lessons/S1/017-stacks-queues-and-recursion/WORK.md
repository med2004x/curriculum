# Lesson 017 — Problem Set

**Topic:** Stacks, Queues, and Recursion.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a stack using an array with `push`, `pop`, and `peek`.
2. Implement a queue using a ring buffer or pointer list.
3. Write a recursive function that computes factorial or traverses a directory tree.
4. Explain the function call stack in terms of the pushed frames.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/017-stacks-queues-and-recursion/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Stack + queue toolkit

Implement a stack and queue API and use each to solve a small parsing or scheduling problem. Compare the two data-flow models.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
