# Lesson 016 — Problem Set

**Topic:** Linked Lists, Invariants, and Cost.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a singly linked list node and list head.
2. Implement `push_front`, `search`, and `delete_value`.
3. Write one example where the list has a cycle and explain why traversal never ends.
4. Write one example where a freed node remains reachable and explain why that is unsafe.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/016-linked-lists-invariants-and-cost/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Invariant list

Implement a singly linked list with insertion and deletion. State the empty/single/multiple-node invariants and test each transition.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
