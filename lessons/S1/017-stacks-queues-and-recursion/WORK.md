# Lesson 017 — Stacks, Queues, and Recursion Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Implement a stack using an array with `push`, `pop`, and `peek`.
2. Implement a queue using a ring buffer or pointer list.
3. Write a recursive function that computes factorial or traverses a directory tree.
4. Explain the function call stack in terms of the pushed frames.

## Run and check

```powershell
Set-Location lessons/S1/017-stacks-queues-and-recursion/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Stack + queue toolkit

Implement a stack and queue API and use each to solve a small parsing or scheduling problem. Compare the two data-flow models.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
