# Lesson 052 — Problem Set

**Topic:** CTEs, Window Functions, and Query Composition.
**Your files:** `code/exercise.sql`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Write a query that ranks events newest-first within each task.

2. Put the ranking query in a CTE and select the latest event per task.

3. Add a tie-breaker and compare `ROW_NUMBER` with `RANK` when timestamps tie.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Stretch problem

### Windowed analysis

Use CTEs and window functions to produce a per-customer ranking and running total. Explain each intermediate relation before combining it.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
