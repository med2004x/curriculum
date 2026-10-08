# Lesson 082 — B+ Trees and Index Structures Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page).

2. Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.

3. Insert keys in increasing order (1, 2, 3, 4, 5, ...) and trace, at each insert, whether a split occurred and at which level.

## Run and check

```powershell
Set-Location lessons/S2/082-btree-and-index-structures/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### B+ tree lab

Implement a small B+ tree insertion/search model, trigger a split, and print the tree before and after. Explain the separator-key invariant.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
