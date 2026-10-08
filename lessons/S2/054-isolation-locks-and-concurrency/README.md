# Lesson 054 — Isolation, Locks, and Concurrency

## Threads
**Primary:** Databases/OS

**Supporting:** anomalies; locking; isolation levels

**Math companion:** probability and interleavings

**Time:** 3–4 hours. **Prerequisites:** Lesson 053.

**Time:** 3–4 hours. **Prerequisites:** Lesson 053.

## What you will be able to do

1. Explain how PostgreSQL MVCC gives statements and transactions snapshots of row versions.
2. Compare the practical visibility guarantees of Read Committed, Repeatable Read, and Serializable.
3. Use row locks or atomic SQL updates to protect a concurrent state transition.
4. Handle blocking, deadlocks, and serialization failures as normal database outcomes.

## Retrieve first (no notes)

1. What does a B-tree index accelerate?
2. What is the difference between an index scan and a sequential scan?
3. Why might PostgreSQL choose not to use an available index?
4. What information does `EXPLAIN` provide that application timing alone does not?

## 1. Problem and model

**What problem this solves.** Concurrent transactions can read and update the same data. Isolation rules and locks determine which versions are visible, which operations wait, and which conflicting outcomes PostgreSQL rejects.

**The contract.** A transaction groups reads and writes, but the chosen isolation level defines visibility and anomaly behavior. Use a single atomic `UPDATE` when the invariant fits one statement; use `SELECT ... FOR UPDATE` when a multi-step decision must lock selected rows. Keep transactions short and handle database-detected conflicts.

**Worked example.** Two sessions adjust one account balance. Compare `UPDATE accounts SET balance = balance + $1 WHERE id = $2` with an application read-modify-write sequence based on a stale value. The single SQL update lets PostgreSQL serialize conflicting row updates without overwriting from an old application snapshot.

Self-check: identify the invariant, the rows it spans, the transaction isolation, and the lock/statement that protects the transition.

## 2. Mechanism and vocabulary

**Mechanism.** PostgreSQL uses MVCC row versions so readers and writers can often proceed without blocking each other. Row-level locks coordinate conflicting writes and explicit locking reads. Locks are generally held until transaction end; a long transaction therefore extends contention and can block other work.

**Second point.** Read Committed takes a new snapshot for each statement. Repeatable Read holds a transaction-level snapshot; PostgreSQL may still report serialization anomalies under this level. Serializable uses Serializable Snapshot Isolation and can abort a transaction that must be retried to preserve a serializable result.

`SELECT ... FOR UPDATE` locks selected rows until transaction completion. Acquire multiple row locks in a consistent order where possible. PostgreSQL can detect a deadlock and abort one participant; applications should roll back and decide whether a bounded retry is appropriate.

**Worked example.** To transfer funds, validate source and destination accounts, lock both rows in a stable ID order, verify sufficient funds, update both balances, and commit. If any step fails, roll back so only one side cannot be debited or credited alone.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```sql
BEGIN;
SELECT id, balance
FROM app.accounts
WHERE id IN (1, 2)
ORDER BY id
FOR UPDATE;
-- Validate funds and update both rows in this transaction.
COMMIT;
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

The stable `ORDER BY id` expresses one lock acquisition order for both accounts. The validation and updates must remain in the same transaction.

Self-check: what should happen if two transfers request the same pair of accounts in opposite directions?

## 4. Verification, failure, and evidence

**Verification.** Use two sessions to observe a conflicting row update or `FOR UPDATE` wait. Record isolation level, statement order, visible values, and commit/rollback outcome. Verify the application retries only errors classified as retryable and does not report success before commit.

**Failure mode.** Have two transactions lock account rows in opposite orders to create a deadlock. Then make both acquire rows by ascending ID and explain how that removes this cycle. Separately describe a serialization failure under Serializable and why the whole transaction, not only its final statement, must be retried.

**Completion criterion.** You can explain snapshot visibility, choose atomic updates or row locks for an invariant, interpret a blocked/deadlocked transaction, and handle serialization failures without partial state changes.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether a second `FOR UPDATE` on the same row proceeds immediately or waits.
2. Predict what a Read Committed statement sees after another transaction commits.
3. Predict what Repeatable Read preserves across statements in one transaction.
4. Predict what happens when Serializable detects a non-serializable execution.
5. Predict the partial state after a transfer transaction rolls back.
6. Predict how opposite row-lock order can deadlock.
7. Predict how ascending ID order changes the wait-for cycle.
8. State which errors the application may retry and why retries must be bounded.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement a balance adjustment as one atomic parameterized `UPDATE` and define its expected row-count behavior.

2. Design a two-account transfer transaction using stable row-lock ordering and rollback on any failed validation.

3. Compare the expected conflict behavior under Read Committed, Repeatable Read, and Serializable for a stated two-session schedule.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Construct the opposite-order row-lock deadlock and identify the wait-for cycle.

2. Apply ascending key order and explain why both transactions now request rows consistently.

3. Create a serialization conflict and document the correct whole-transaction retry boundary.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one transaction property: run the operation at a different isolation level or add a third account row. Explain which snapshots, locks, and retry outcomes may change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What visibility does MVCC provide to a transaction?
2. How does Read Committed snapshot behavior differ from Repeatable Read?
3. What does `SELECT ... FOR UPDATE` protect?
4. How does consistent lock ordering prevent a particular deadlock?
5. What is a serialization failure and what should be retried?
6. Why should a transaction be short?

## Evidence to record

Create `lesson054/EVIDENCE.md` with exactly this structure:

```text
# Lesson 054 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Confusing database row locks with Go mutexes or condition variables.
- Reporting success before transaction commit.
- Retrying only the final statement after a serialization failure.
- Locking multiple rows in inconsistent orders.
- Assuming Repeatable Read and Serializable have identical guarantees.

## Precision checks

1. State the snapshot guarantee for the chosen isolation level.
2. Explain the lock held by the transfer transaction and its lifetime.
3. Show how rollback preserves the two-account invariant.
4. Explain deadlock detection and the lock-order rule.
5. State the whole-transaction retry behavior for serialization errors.

Database isolation is a contract over concurrent histories; row locks and transaction retries must preserve the application invariant.
## University revision bridge

For the LIG track, connect transaction histories and isolation anomalies to the concrete MVCC snapshots, locks, and serialization checks PostgreSQL implements.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 054 — Isolation, Locks, and Concurrency without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your bounded buffer implementation is correct under concurrent producer/consumer load across multiple runs, you've reliably reproduced a two-lock deadlock and diagnosed its exact structure using a debugger rather than just observing a hang, and you've fixed it with lock ordering and confirmed — across many repeated runs — that the fix is structural, not just. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
