# Lesson 054 — Isolation, Locks, and Concurrency Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.sql`.

## Problems

1. Implement a balance adjustment as one atomic parameterized `UPDATE` and define its expected row-count behavior.

2. Design a two-account transfer transaction using stable row-lock ordering and rollback on any failed validation.

3. Compare the expected conflict behavior under Read Committed, Repeatable Read, and Serializable for a stated two-session schedule.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Optional challenge

### Isolation experiment

Open concurrent transactions and reproduce a chosen isolation phenomenon. Record the exact schedule and explain which guarantee prevented or allowed it.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
