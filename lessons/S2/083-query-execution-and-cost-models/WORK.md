# Lesson 083 — Query Execution and Cost Models Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Implement the iterator interface (`Open()`, `Next() (row, hasMore)`, `Close()`, or the Go-idiomatic equivalent) and a `SeqScan` operator that iterates over an in-memory slice of rows (standing in for reading table pages, per Lesson 1) one at a time.

2. Implement a `Filter` operator that wraps any child iterator and a predicate function, calling the child's `Next()` and only returning rows where the predicate is true — confirm it correctly stops (returns `hasMore = false`) once the child is exhausted, not just once it happens to find no more matching rows in a finite lookahead.

3. Implement a `Project` operator that wraps a child iterator and a list of column indices/names to keep, transforming each row before passing it up.

## Run and check

```powershell
Set-Location lessons/S2/083-query-execution-and-cost-models/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Query cost model

Given a relation size and selectivity assumptions, estimate costs for sequential and indexed plans, then compare with `EXPLAIN` on a real query.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
