# Lesson 084 — Transactions, WAL, and Recovery

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Databases must preserve committed changes across crashes without synchronously writing every modified data page at commit. WAL records make recovery possible by recording changes before data pages are persisted.

**The contract.** Write-ahead logging requires the log record describing a change to reach durable storage before the modified data page is written. A commit is acknowledged only after the commit record meets the configured durability policy. Checkpoints bound recovery work; they do not replace the log or mean every dirty page was flushed at that instant.

**Worked example.** Trace a transaction that updates a row: the storage engine emits WAL records, ensures the required log prefix is durable at commit, and may write the modified data page later. After a crash, recovery uses the log to redo changes required by committed transactions and restore a consistent database state.

## 2. Mechanism and vocabulary

**Mechanism.** WAL is append-oriented; log sequence numbers identify record positions. A checkpoint records a recovery boundary and flush progress. Recovery examines WAL from an appropriate starting point; exact redo/undo details vary by database architecture, so distinguish PostgreSQL's redo-oriented crash recovery from generic ARIES descriptions.

**Second point.** Isolation anomalies are related but separate from crash recovery. PostgreSQL Repeatable Read uses snapshot isolation and can reject concurrent updates, but write skew can still be possible; Serializable adds detection that may abort a transaction. Applications must handle and, when safe, retry serialization errors.

**Worked example.** Compare two-session outcomes for a row update at Read Committed and Repeatable Read, then construct a write-skew invariant and explain why Serializable may abort one transaction. Keep this anomaly exercise as a separate section from WAL recovery.

## 3. Worked case from the reference material

Worked example: Using PostgreSQL (from the postgresql-engineering track) with two separate `psql` sessions (simulating two concurrent transactions), reproduce a lost update at `READ COMMITTED` isolation (PostgreSQL's default): both sessions `BEGIN`, both `SELECT` the same counter value, both independently compute value+1, both `UPDATE` and `COMMIT` — confirm the counter onl.

## 4. Verification, failure, and evidence

**Verification.** State the configured durability assumptions, the commit/WAL relationship, checkpoint role, and recovery expectation. For the separate isolation exercise, record PostgreSQL version, isolation level, transaction order, final state, and exact serialization error if one occurs.

**Failure mode.** Explain what a system risks if it acknowledges commit before the required WAL is durable, then contrast that with data pages being written later. Separately ignore a serialization error in the isolation exercise and show why the application-level invariant may remain unsatisfied.

**Completion criterion.** You can state the WAL ordering rule, explain commit durability and checkpoint/recovery roles, and separately describe how PostgreSQL isolation errors must be handled by the application.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
