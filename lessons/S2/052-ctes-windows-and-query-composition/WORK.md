# Lesson 052 — CTEs, Window Functions, and Query Composition Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.sql`.

## Problems

1. Write a query that ranks events newest-first within each task.

2. Put the ranking query in a CTE and select the latest event per task.

3. Add a tie-breaker and compare `ROW_NUMBER` with `RANK` when timestamps tie.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Optional challenge

### Windowed analysis

Use CTEs and window functions to produce a per-customer ranking and running total. Explain each intermediate relation before combining it.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
