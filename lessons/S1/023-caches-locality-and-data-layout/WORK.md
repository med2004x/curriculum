# Lesson 023 — Problem Set

**Topic:** Caches, Locality, and Data Layout.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Write two correct summation loops over one initialized two-dimensional array, changing only loop order.

2. Compare multiple sizes, including a small working set and one substantially larger working set.

3. Record repeated timings and, when available, cache-miss counters; confirm both variants produce the same sum.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/023-caches-locality-and-data-layout/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Locality experiment

Compare two data layouts over a large dataset and benchmark traversal. Explain the result using cache lines and access locality rather than hand-waving.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
