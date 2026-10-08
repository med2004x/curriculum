# Lesson 053 — Indexes, B-Trees, and EXPLAIN Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.sql`.

## Problems

1. Capture a baseline plan for the project/task query and record the relevant table/index context.

2. Add a composite B-tree index matching the filter and ordering pattern; capture the new plan.

3. Compare selective and nonselective variants and explain the planner's cost trade-off.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Optional challenge

### Index investigation

Create a realistic dataset, compare a query with and without an index, inspect `EXPLAIN`, and explain whether the planner decision matches your expectation.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
