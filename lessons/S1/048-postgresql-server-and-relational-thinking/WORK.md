# Lesson 048 — PostgreSQL, psql, and Relational Thinking Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.sql`.

## Problems

1. Connect with `psql` and inspect connection identity, database, schema, tables, and roles.

2. Query `app.tasks` for selected columns and filter rows with a `WHERE` predicate.

3. Compare duplicates, NULL predicates, and result ordering with and without the relevant SQL clauses.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Optional challenge

### Relational schema

Design a PostgreSQL schema for a small service. Create it with psql, inspect it, insert representative rows, and prove one relational query is correct.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
