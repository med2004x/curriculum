# Lesson 054 — Problem Set

**Topic:** Isolation, Locks, and Concurrency.
**Your files:** `code/exercise.sql`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a balance adjustment as one atomic parameterized `UPDATE` and define its expected row-count behavior.

2. Design a two-account transfer transaction using stable row-lock ordering and rollback on any failed validation.

3. Compare the expected conflict behavior under Read Committed, Repeatable Read, and Serializable for a stated two-session schedule.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Stretch problem

### Isolation experiment

Open concurrent transactions and reproduce a chosen isolation phenomenon. Record the exact schedule and explain which guarantee prevented or allowed it.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
