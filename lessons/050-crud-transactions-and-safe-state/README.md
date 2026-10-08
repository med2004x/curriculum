# Lesson 050 — CRUD, Transactions, and Safe State

## Threads
**Primary:** Databases/Go

**Supporting:** CRUD; transaction boundaries; safe updates

**Math companion:** logic and state transitions

**Time:** 3–4 hours. **Prerequisites:** Lesson 049.

**Time:** 3–4 hours. **Prerequisites:** Lesson 049.

## What you will be able to do

1. Implement create, read, update, and delete operations with explicit not-found and validation behavior.
2. Use parameterized SQL rather than string concatenation for user-controlled values.
3. Group multi-statement state changes in a transaction and handle rollback/commit errors.
4. Verify the before/after database state and distinguish a no-op from a successful mutation.

## Retrieve first (no notes)

1. What does a primary key identify?
2. How does a foreign key prevent orphan relationships?
3. How can a junction table represent many-to-many data?
4. Which constraints can reject invalid state even if an application has a bug?

## 1. Problem and model

**What problem this solves.** Persistent state changes must preserve domain invariants despite invalid input, concurrent requests, and partial failures.

**The contract.** Create validates required fields and reports conflicts; read distinguishes not-found from query failure; update changes only the intended row and reports whether it matched; delete obeys relationship constraints. A multi-row operation either commits all intended changes or rolls them back together.

**Worked example.** Update one task using `UPDATE app.tasks SET status = $1 WHERE id = $2 RETURNING id, status`. The placeholders keep values separate from SQL syntax. If no row is returned, report not-found rather than claiming an update occurred.

Self-check: what is the observable distinction between one updated row, no matching row, and a database error?

## 2. Mechanism and vocabulary

**Mechanism.** A transaction groups statements under a database transaction boundary. `COMMIT` makes its changes durable according to the database's transaction guarantees; `ROLLBACK` discards uncommitted changes. In Go's `database/sql`, use `BeginTx`, issue statements through the `sql.Tx`, and commit only after each operation succeeds.

**Second point.** Transactions provide atomic grouping, not automatic correctness: choose the right isolation/locking behavior for concurrent invariants and handle serialization, constraint, and commit errors. A successful earlier statement does not mean the whole transaction committed.

Never construct SQL by concatenating user-controlled text into the query. Parameterization protects values from being interpreted as SQL syntax; identifiers and sort directions generally require allow-listing because they cannot be bound as ordinary value parameters.

**Worked example.** Add one boundary case to the # 050 — CRUD, Transactions, and Safe State exercise.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: which rows should be visible after a rollback? Which error can occur at commit time?

## 4. Verification, failure, and evidence

**Verification.** For each CRUD operation, record input, expected row count or returned row, and database state before/after. Force a failure between the two transaction statements and confirm neither partial change remains. Check the commit result as well as individual statement results.

**Failure mode.** Cause the second statement in the transaction to violate a constraint. Confirm rollback leaves the task and event tables consistent. Separately update a nonexistent ID and verify the API reports not-found or zero affected rows.

**Completion criterion.** You can state each operation's success/failure contract, use parameters, and preserve a multi-table invariant across both success and failure.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the effect of updating an ID that does not exist.
2. Predict the final state if the event insert fails before commit.
3. Predict whether an uncommitted update should be treated as completed by the API.
4. Predict how a constraint violation is reported.
5. Predict why interpolating a title into SQL is unsafe.
6. Predict which writes must use the same transaction to preserve the domain invariant.
7. Predict how `RETURNING` can distinguish changed data from no matching row.
8. State the database state expected after both commit and rollback paths.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement create/read/update/delete functions with parameterized SQL and explicit error handling.

2. Add a transaction that updates a task and inserts its event record atomically.

3. Add tests or a written state trace for success, not-found, constraint failure, rollback, and commit failure.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one failure point: fail the update, fail the event insert, or fail validation before the transaction begins. Explain which state changes are allowed and whether rollback is required.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What should each CRUD operation return for success and not-found?
2. Why should SQL values be parameterized?
3. What does a transaction make atomic?
4. What must happen when a statement fails before commit?
5. Why check the result of `Commit`?
6. How can `RETURNING` or affected-row counts distinguish no match from success?

## Evidence to record

Create `lesson050/EVIDENCE.md` with exactly this structure:

```text
# Lesson 050 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Concatenating user input into SQL text.
- Returning success before the transaction commits.
- Ignoring the commit error.
- Splitting one invariant-preserving operation across separate transactions.
- Treating zero affected rows as a database failure without defining not-found behavior.

## Precision checks

1. State the success and failure contract for each CRUD operation.
2. Identify every write that belongs in one transaction.
3. Explain the state after rollback and after commit.
4. Show that user-controlled values are parameterized.
5. Distinguish not-found, validation, constraint, query, and commit errors.

Safe state changes require both database constraints and transaction boundaries; neither replaces explicit API error contracts.
## University revision bridge

For the LIG track, connect CRUD behavior to relational constraints and transaction state transitions, then define how the Go API surfaces each database outcome.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 050 — CRUD, Transactions, and Safe State without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can reproduce the main behavior and diagnose one failure from evidence. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
