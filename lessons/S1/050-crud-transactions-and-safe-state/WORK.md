# Lesson 050 — Problem Set

**Topic:** CRUD, Transactions, and Safe State.
**Your files:** `code/exercise.sql`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement create/read/update/delete functions with parameterized SQL and explicit error handling.

2. Add a transaction that updates a task and inserts its event record atomically.

3. Add tests or a written state trace for success, not-found, constraint failure, rollback, and commit failure.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Stretch problem

### Transactional CRUD

Build CRUD operations whose state changes are transactionally safe. Demonstrate one rollback and one integrity violation without leaving partial state.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
