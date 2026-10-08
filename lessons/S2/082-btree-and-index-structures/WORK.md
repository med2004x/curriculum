# Lesson 082 — Problem Set

**Topic:** B+ Trees and Index Structures.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page).

2. Extend it with one explicit boundary case: Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.

3. Rebuild the same behavior with one changed constraint: Insert keys in increasing order (1, 2, 3, 4, 5, ...) and trace, at each insert, whether a split occurred and at which level.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/082-btree-and-index-structures/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### B+ tree lab

Implement a small B+ tree insertion/search model, trigger a split, and print the tree before and after. Explain the separator-key invariant.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
