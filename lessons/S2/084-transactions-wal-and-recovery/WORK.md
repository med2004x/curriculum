# Lesson 084 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 9. Security, Database Internals, and Distributed Systems  
**Primary:** DB Internals/OS  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. State the write-ahead rule.
2. What does commit acknowledgment require from WAL under the selected settings?
3. What is a checkpoint's role?
4. Why can data pages be written after commit?
5. What does crash recovery use WAL for?
6. How is a serialization failure different from a crash-recovery event?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict why WAL must be durable before a changed data page is written.
2. Predict what commit acknowledgment implies under the stated durability settings.
3. Predict what a checkpoint changes and what it does not guarantee.
4. Predict why data pages may lag behind committed WAL.
5. Predict which committed changes recovery must preserve after a crash.
6. Predict which database anomaly is a concurrency issue rather than a crash-recovery issue.
7. Predict how the application should respond to a serialization failure.
8. State which PostgreSQL-specific behavior needs version/setting context.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Explain the write-ahead logging rule and why log records must become durable before corresponding data pages.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Distinguish commit/WAL durability from checkpointing and data-page flushing.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Trace crash recovery at a high level using log sequence numbers, redo, and transaction status.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

Given a small relational dataset and a business question, write the query or transaction schedule on paper first. Then verify it against a deliberately chosen edge case that would break a naive solution.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### WAL + recovery simulation

Model a sequence of page changes and WAL records, simulate a crash at two points, and determine the recovered state under each point.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Take your step 2 fix (`REPEATABLE READ` preventing the lost update via a commit-time error) and confirm your *application code* actually needs to handle that error explicitly — rerun the scenario but have the "losing" transaction's code simply ignore the serialization error rather than catching it and retrying.

2. Confirm the increment is effectively lost anyway, not because the database allowed the anomaly, but because the application discarded the database's correct rejection instead of retrying the transaction.

3. Explain why choosing a stronger isolation level is necessary but not sufficient on its own — the database can correctly detect and reject the unsafe interleaving, but the application must still contain retry logic for the rejection to actually result in correct final behavior, rather than an error that's silently swallowed and produces the exact same lost up.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one durability setting or crash point in the model. Explain which acknowledged state must survive recovery and what storage assumptions affect the guarantee.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Model:** State the system model and assumptions before solving.
2. **Trace:** Work a small failure or state-transition example by hand.
3. **Trade-off:** Compare two designs and state the condition under which each wins.
4. **Evidence:** Identify the measurement, test, or trace that would distinguish the designs in practice.

## 11. Self-test

- [ ] State the write-ahead rule.
- [ ] What does commit acknowledgment require from WAL under the selected settings?
- [ ] What is a checkpoint's role?
- [ ] Why can data pages be written after commit?
- [ ] What does crash recovery use WAL for?
- [ ] How is a serialization failure different from a crash-recovery event?

Score yourself only after the attempt: **0 = cannot answer, 1 = answer with notes, 2 = answer from memory, 3 = answer and justify a new variation.**

## 12. Evidence and grading

Create `EVIDENCE.md` in this lesson directory. Record:

- the exact commands or test commands used;
- one before/after prediction;
- one failure reproduced deliberately and its diagnosis;
- the result of the coding challenge;
- one transfer variation;
- your self-test score.

### Gate rubric

| Dimension | Pass condition |
|---|---|
| Knowledge | Explain the core model without notes. |
| Procedure | Produce a working implementation or derivation from blank. |
| Debugging | Reproduce, classify, and repair one realistic failure. |
| Transfer | Solve a changed constraint without being handed the algorithm. |
| Evidence | Another person can reproduce the verification from your notes. |

Suggested pass threshold: all five dimensions pass. If one fails, review that dimension before moving to the next lesson.

## 13. Mastery gate

Use the lecture's mastery gate in [`README.md`](README.md) as the primary criterion. This work packet adds the stronger requirement that you can satisfy the gate **after** completing independent work, not by recognising the lecture examples.

For integrated assessment lessons, the gate is summative: treat the challenge as an exam/project submission and do not consult previous solutions while working.

## 14. Spaced review

- **Next day:** answer the self-test from memory and redo one exercise without looking.
- **One week:** solve the transfer problem again under a changed constraint.
- **Later module review:** revisit this lesson when the concept reappears in a later module and compare the old and new evidence.

## 15. External practice reference

- Use the lesson's assigned material and primary technical documentation when a precise semantic detail needs verification. The repository challenge remains the assessment artifact.
