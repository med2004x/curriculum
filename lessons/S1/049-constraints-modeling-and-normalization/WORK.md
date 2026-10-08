# Lesson 049 — Constraints, Modeling, and Normalization Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.sql`.

## Problems

1. Define task, tag, and task-tag association tables with primary and foreign keys.

2. Add `UNIQUE`, `NOT NULL`, and `CHECK` constraints that encode the domain rules.

3. Demonstrate invalid inserts and query the many-to-many relationship.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Optional challenge

### Constraint-driven model

Model a small business domain using keys, foreign keys, uniqueness, checks, and normalized relationships. Show a malformed state each constraint prevents.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
