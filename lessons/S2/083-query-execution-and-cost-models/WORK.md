# Lesson 083 — Problem Set

**Topic:** Query Execution and Cost Models.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement the iterator interface (`Open()`, `Next() (row, hasMore)`, `Close()`, or the Go-idiomatic equivalent) and a `SeqScan` operator that iterates over an in-memory slice of rows (standing in for reading table pages, per Lesson 1) one at a time.

2. Extend it with one explicit boundary case: Implement a `Filter` operator that wraps any child iterator and a predicate function, calling the child's `Next()` and only returning rows where the predicate is true — confirm it correctly stops (returns `hasMore = false`) once the child is exhausted, not just once it happens to find no more matching rows in a finite lookahead.

3. Rebuild the same behavior with one changed constraint: Implement a `Project` operator that wraps a child iterator and a list of column indices/names to keep, transforming each row before passing it up.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/083-query-execution-and-cost-models/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Query cost model

Given a relation size and selectivity assumptions, estimate costs for sequential and indexed plans, then compare with `EXPLAIN` on a real query.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
