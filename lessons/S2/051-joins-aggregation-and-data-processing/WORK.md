# Lesson 051 — Problem Set

**Topic:** Joins, Aggregation, and Data Processing.
**Your files:** `code/exercise.sql`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Query task/tag pairs by joining all three tables.

2. Add a left join so tasks without tags remain visible.

3. Aggregate tag counts per task and add a `HAVING` condition on the count.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Stretch problem

### Reporting query

Write a reporting query using joins and aggregation. Validate totals against a manually computed sample and explain duplicate-row effects.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
