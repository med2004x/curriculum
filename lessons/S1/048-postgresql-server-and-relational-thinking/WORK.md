# Lesson 048 — Problem Set

**Topic:** PostgreSQL, psql, and Relational Thinking.
**Your files:** `code/exercise.sql`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Connect with `psql` and inspect connection identity, database, schema, tables, and roles.

2. Query `app.tasks` for selected columns and filter rows with a `WHERE` predicate.

3. Compare duplicates, NULL predicates, and result ordering with and without the relevant SQL clauses.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Stretch problem

### Relational schema

Design a PostgreSQL schema for a small service. Create it with psql, inspect it, insert representative rows, and prove one relational query is correct.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
