# Lesson 019 — Problem Set

**Topic:** Searching, Sorting, and Measuring Work.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a linear search and a binary search on an integer array.
2. Implement insertion sort and merge sort on a small integer array.
3. Measure the runtime of each algorithm over several input sizes.
4. Record the input sizes and observed timing pattern.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/019-search-sorting-and-profiling/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Search/sort benchmark

Implement one search and two sorting strategies. Measure operations on controlled inputs and reconcile observed results with asymptotic predictions.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
