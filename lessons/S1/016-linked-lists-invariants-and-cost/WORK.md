# Lesson 016 — Linked Lists, Invariants, and Cost Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Implement a singly linked list node and list head.
2. Implement `push_front`, `search`, and `delete_value`.
3. Write one example where the list has a cycle and explain why traversal never ends.
4. Write one example where a freed node remains reachable and explain why that is unsafe.

## Run and check

```powershell
Set-Location lessons/S1/016-linked-lists-invariants-and-cost/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Invariant list

Implement a singly linked list with insertion and deletion. State the empty/single/multiple-node invariants and test each transition.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
