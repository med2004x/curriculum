# Lesson 051 — Joins, Aggregation, and Data Processing Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.sql`.

## Problems

1. Query task/tag pairs by joining all three tables.

2. Add a left join so tasks without tags remain visible.

3. Aggregate tag counts per task and add a `HAVING` condition on the count.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Optional challenge

### Reporting query

Write a reporting query using joins and aggregation. Validate totals against a manually computed sample and explain duplicate-row effects.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
