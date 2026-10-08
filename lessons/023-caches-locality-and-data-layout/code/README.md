# Code map — Lesson 023 — Caches, Locality, and Data Layout

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Locality experiment

Compare two data layouts over a large dataset and benchmark traversal. Explain the result using cache lines and access locality rather than hand-waving.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Shrink the array size until row-major and column-major run at the same speed.

2. Find (approximately, by binary search on size) the array size where the effect disappears, and explain why in terms of the array now fitting inside L2 or L3 cache regardless of access order — locality only matters when the working set exceeds cache capacity.

3. Create a misleading use of `volatile` and explain why it does not establish a safe communication protocol between threads.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint: use a one-dimensional array with a configurable stride, or change element size while keeping total bytes approximately constant. Predict the locality effect before measuring it.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic dr

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
