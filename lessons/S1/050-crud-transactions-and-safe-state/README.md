# Lesson 050 — CRUD, Transactions, and Safe State

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Persistent state changes must preserve domain invariants despite invalid input, concurrent requests, and partial failures.

**The contract.** Create validates required fields and reports conflicts; read distinguishes not-found from query failure; update changes only the intended row and reports whether it matched; delete obeys relationship constraints. A multi-row operation either commits all intended changes or rolls them back together.

**Worked example.** Update one task using `UPDATE app.tasks SET status = $1 WHERE id = $2 RETURNING id, status`. The placeholders keep values separate from SQL syntax. If no row is returned, report not-found rather than claiming an update occurred.

## 2. Mechanism and vocabulary

**Mechanism.** A transaction groups statements under a database transaction boundary. `COMMIT` makes its changes durable according to the database's transaction guarantees; `ROLLBACK` discards uncommitted changes. In Go's `database/sql`, use `BeginTx`, issue statements through the `sql.Tx`, and commit only after each operation succeeds.

**Second point.** Transactions provide atomic grouping, not automatic correctness: choose the right isolation/locking behavior for concurrent invariants and handle serialization, constraint, and commit errors. A successful earlier statement does not mean the whole transaction committed.

Never construct SQL by concatenating user-controlled text into the query. Parameterization protects values from being interpreted as SQL syntax; identifiers and sort directions generally require allow-listing because they cannot be bound as ordinary value parameters.

**Worked example.** Add one boundary case to the # 050 — CRUD, Transactions, and Safe State exercise.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```sql
BEGIN;
UPDATE app.tasks
SET status = 'done'
WHERE id = 42;
INSERT INTO app.task_events (task_id, event_type)
VALUES (42, 'completed');
COMMIT;
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

The update and event insert belong to one logical state transition. If the event insert fails, roll back rather than leaving the task marked done without its corresponding event.

## 4. Verification, failure, and evidence

**Verification.** For each CRUD operation, record input, expected row count or returned row, and database state before/after. Force a failure between the two transaction statements and confirm neither partial change remains. Check the commit result as well as individual statement results.

**Failure mode.** Cause the second statement in the transaction to violate a constraint. Confirm rollback leaves the task and event tables consistent. Separately update a nonexistent ID and verify the API reports not-found or zero affected rows.

**Completion criterion.** You can state each operation's success/failure contract, use parameters, and preserve a multi-table invariant across both success and failure.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
