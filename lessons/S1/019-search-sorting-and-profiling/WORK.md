# Lesson 019 — Searching, Sorting, and Measuring Work Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Implement a linear search and a binary search on an integer array.
2. Implement insertion sort and merge sort on a small integer array.
3. Measure the runtime of each algorithm over several input sizes.
4. Record the input sizes and observed timing pattern.

## Run and check

```powershell
Set-Location lessons/S1/019-search-sorting-and-profiling/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Search/sort benchmark

Implement one search and two sorting strategies. Measure operations on controlled inputs and reconcile observed results with asymptotic predictions.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
