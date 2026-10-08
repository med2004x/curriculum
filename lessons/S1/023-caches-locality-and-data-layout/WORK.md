# Lesson 023 — Caches, Locality, and Data Layout Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Write two correct summation loops over one initialized two-dimensional array, changing only loop order.

2. Compare multiple sizes, including a small working set and one substantially larger working set.

3. Record repeated timings and, when available, cache-miss counters; confirm both variants produce the same sum.

## Run and check

```powershell
Set-Location lessons/S1/023-caches-locality-and-data-layout/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Locality experiment

Compare two data layouts over a large dataset and benchmark traversal. Explain the result using cache lines and access locality rather than hand-waving.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
