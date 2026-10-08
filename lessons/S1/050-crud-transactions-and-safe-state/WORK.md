# Lesson 050 — CRUD, Transactions, and Safe State Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.sql`.

## Problems

1. Implement create/read/update/delete functions with parameterized SQL and explicit error handling.

2. Add a transaction that updates a task and inserts its event record atomically.

3. Add tests or a written state trace for success, not-found, constraint failure, rollback, and commit failure.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Optional challenge

### Transactional CRUD

Build CRUD operations whose state changes are transactionally safe. Demonstrate one rollback and one integrity violation without leaving partial state.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
