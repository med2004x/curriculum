# Lesson 049 — Problem Set

**Topic:** Constraints, Modeling, and Normalization.
**Your files:** `code/exercise.sql`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define task, tag, and task-tag association tables with primary and foreign keys.

2. Add `UNIQUE`, `NOT NULL`, and `CHECK` constraints that encode the domain rules.

3. Demonstrate invalid inserts and query the many-to-many relationship.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Stretch problem

### Constraint-driven model

Model a small business domain using keys, foreign keys, uniqueness, checks, and normalized relationships. Show a malformed state each constraint prevents.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
