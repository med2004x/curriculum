# Code map — Lesson 082 — B+ Trees and Index Structures

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### B+ tree lab

Implement a small B+ tree insertion/search model, trigger a split, and print the tree before and after. Explain the separator-key invariant.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Deliberately implement `rangeSearch` *without* the leaf-chain optimization — instead, for every key you want in the range, perform a fresh, independent tree search from the root.

2. Measure and compare the total number of node visits for a range query spanning, say, 20 keys using this naive re-search approach versus your step 3 leaf-chain approach.

3. Report the actual difference in node-visit count, and explain why the leaf-chain approach's advantage grows with the size of the range being queried — a range of 2 keys barely matters, but a range of thousands does, which is exactly the shape of query B+ tree indexes are specifically optimized to make cheap.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint in the main case for # 082 — B+ Trees and Index Structures: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculatio

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
