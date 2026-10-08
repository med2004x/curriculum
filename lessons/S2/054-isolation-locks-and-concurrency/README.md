# Lesson 054 — Isolation, Locks, and Concurrency

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Concurrent transactions can read and update the same data. Isolation rules and locks determine which versions are visible, which operations wait, and which conflicting outcomes PostgreSQL rejects.

**The contract.** A transaction groups reads and writes, but the chosen isolation level defines visibility and anomaly behavior. Use a single atomic `UPDATE` when the invariant fits one statement; use `SELECT ... FOR UPDATE` when a multi-step decision must lock selected rows. Keep transactions short and handle database-detected conflicts.

**Worked example.** Two sessions adjust one account balance. Compare `UPDATE accounts SET balance = balance + $1 WHERE id = $2` with an application read-modify-write sequence based on a stale value. The single SQL update lets PostgreSQL serialize conflicting row updates without overwriting from an old application snapshot.

## 2. Mechanism and vocabulary

**Mechanism.** PostgreSQL uses MVCC row versions so readers and writers can often proceed without blocking each other. Row-level locks coordinate conflicting writes and explicit locking reads. Locks are generally held until transaction end; a long transaction therefore extends contention and can block other work.

**Second point.** Read Committed takes a new snapshot for each statement. Repeatable Read holds a transaction-level snapshot; PostgreSQL may still report serialization anomalies under this level. Serializable uses Serializable Snapshot Isolation and can abort a transaction that must be retried to preserve a serializable result.

`SELECT ... FOR UPDATE` locks selected rows until transaction completion. Acquire multiple row locks in a consistent order where possible. PostgreSQL can detect a deadlock and abort one participant; applications should roll back and decide whether a bounded retry is appropriate.

**Worked example.** To transfer funds, validate source and destination accounts, lock both rows in a stable ID order, verify sufficient funds, update both balances, and commit. If any step fails, roll back so only one side cannot be debited or credited alone.

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

## 4. Verification, failure, and evidence

**Verification.** Use two sessions to observe a conflicting row update or `FOR UPDATE` wait. Record isolation level, statement order, visible values, and commit/rollback outcome. Verify the application retries only errors classified as retryable and does not report success before commit.

**Failure mode.** Have two transactions lock account rows in opposite orders to create a deadlock. Then make both acquire rows by ascending ID and explain how that removes this cycle. Separately describe a serialization failure under Serializable and why the whole transaction, not only its final statement, must be retried.

**Completion criterion.** You can explain snapshot visibility, choose atomic updates or row locks for an invariant, interpret a blocked/deadlocked transaction, and handle serialization failures without partial state changes.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
