# Lesson 084 — Transactions, WAL, and Recovery

## Threads
**Primary:** DB Internals/OS

**Supporting:** logging; crash recovery; durability

**Math companion:** probability and failure models

**Time:** 3–4 hours. **Prerequisites:** Lesson 083.

**Time:** 3–4 hours. **Prerequisites:** Lesson 083.

## What you will be able to do

1. Explain the write-ahead logging rule and why log records must become durable before corresponding data pages.
2. Distinguish commit/WAL durability from checkpointing and data-page flushing.
3. Trace crash recovery at a high level using log sequence numbers, redo, and transaction status.
4. Explain why application retry handling is still required after serialization or connection failures.

## Retrieve first (no notes)

1. What does a query plan estimate?
2. How does actual row count differ from an estimate?
3. Why can a nonselective predicate favor a sequential scan?
4. What does a B+ tree leaf chain support?

## 1. Problem and model

**What problem this solves.** Databases must preserve committed changes across crashes without synchronously writing every modified data page at commit. WAL records make recovery possible by recording changes before data pages are persisted.

**The contract.** Write-ahead logging requires the log record describing a change to reach durable storage before the modified data page is written. A commit is acknowledged only after the commit record meets the configured durability policy. Checkpoints bound recovery work; they do not replace the log or mean every dirty page was flushed at that instant.

**Worked example.** Trace a transaction that updates a row: the storage engine emits WAL records, ensures the required log prefix is durable at commit, and may write the modified data page later. After a crash, recovery uses the log to redo changes required by committed transactions and restore a consistent database state.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** WAL is append-oriented; log sequence numbers identify record positions. A checkpoint records a recovery boundary and flush progress. Recovery examines WAL from an appropriate starting point; exact redo/undo details vary by database architecture, so distinguish PostgreSQL's redo-oriented crash recovery from generic ARIES descriptions.

**Second point.** Isolation anomalies are related but separate from crash recovery. PostgreSQL Repeatable Read uses snapshot isolation and can reject concurrent updates, but write skew can still be possible; Serializable adds detection that may abort a transaction. Applications must handle and, when safe, retry serialization errors.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Compare two-session outcomes for a row update at Read Committed and Repeatable Read, then construct a write-skew invariant and explain why Serializable may abort one transaction. Keep this anomaly exercise as a separate section from WAL recovery.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Using PostgreSQL (from the postgresql-engineering track) with two separate `psql` sessions (simulating two concurrent transactions), reproduce a lost update at `READ COMMITTED` isolation (PostgreSQL's default): both sessions `BEGIN`, both `SELECT` the same counter value, both independently compute value+1, both `UPDATE` and `COMMIT` — confirm the counter onl.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** State the configured durability assumptions, the commit/WAL relationship, checkpoint role, and recovery expectation. For the separate isolation exercise, record PostgreSQL version, isolation level, transaction order, final state, and exact serialization error if one occurs.

**Failure mode.** Explain what a system risks if it acknowledges commit before the required WAL is durable, then contrast that with data pages being written later. Separately ignore a serialization error in the isolation exercise and show why the application-level invariant may remain unsatisfied.

**Completion criterion.** You can state the WAL ordering rule, explain commit durability and checkpoint/recovery roles, and separately describe how PostgreSQL isolation errors must be handled by the application.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict why WAL must be durable before a changed data page is written.
2. Predict what commit acknowledgment implies under the stated durability settings.
3. Predict what a checkpoint changes and what it does not guarantee.
4. Predict why data pages may lag behind committed WAL.
5. Predict which committed changes recovery must preserve after a crash.
6. Predict which database anomaly is a concurrency issue rather than a crash-recovery issue.
7. Predict how the application should respond to a serialization failure.
8. State which PostgreSQL-specific behavior needs version/setting context.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Draw the order of a WAL record, durable log flush, commit record, and later data-page write.

2. Explain how checkpointing bounds recovery work without replacing WAL.

3. Separately trace the prescribed two-session isolation anomaly and record the exact PostgreSQL result/error.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take your step 2 fix (`REPEATABLE READ` preventing the lost update via a commit-time error) and confirm your *application code* actually needs to handle that error explicitly — rerun the scenario but have the "losing" transaction's code simply ignore the serialization error rather than catching it and retrying.

2. Confirm the increment is effectively lost anyway, not because the database allowed the anomaly, but because the application discarded the database's correct rejection instead of retrying the transaction.

3. Explain why choosing a stronger isolation level is necessary but not sufficient on its own — the database can correctly detect and reject the unsafe interleaving, but the application must still contain retry logic for the rejection to actually result in correct final behavior, rather than an error that's silently swallowed and produces the exact same lost up.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one durability setting or crash point in the model. Explain which acknowledged state must survive recovery and what storage assumptions affect the guarantee.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. State the write-ahead rule.
2. What does commit acknowledgment require from WAL under the selected settings?
3. What is a checkpoint's role?
4. Why can data pages be written after commit?
5. What does crash recovery use WAL for?
6. How is a serialization failure different from a crash-recovery event?

## Evidence to record

Create `lesson084/EVIDENCE.md` with exactly this structure:

```text
# Lesson 084 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 084 — Transactions, WAL, and Recovery.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 084 — Transactions, WAL, and Recovery without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've directly reproduced a lost update at `READ COMMITTED` and confirmed `REPEATABLE READ` prevents it (with an explicit commit-time error, not silent success), you've directly reproduced write skew at `REPEATABLE READ` and confirmed `SERIALIZABLE` prevents it, and you can explain — using the failure drill — why choosing the correct isolation level alone i. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
