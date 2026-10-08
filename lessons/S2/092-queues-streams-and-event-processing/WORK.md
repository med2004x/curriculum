# Lesson 092 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 9. Security, Database Internals, and Distributed Systems  
**Primary:** Distributed/Go  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What is at-least-once delivery?
2. What crash window can cause duplicate processing?
3. What crash window can cause a lost effect?
4. When can an effect and dedupe record be atomic?
5. What ordering does a partition provide?
6. Why is exactly-once delivery not automatically end-to-end exactly-once processing?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict what happens if a consumer crashes after effect but before acknowledgment.
2. Predict what happens if it commits the offset before performing the effect.
3. Predict how an idempotent `set` differs from a non-idempotent `append` on redelivery.
4. Predict what atomic DB deduplication prevents.
5. Predict the ordering guarantee within one partition.
6. Predict whether two consumers in one consumer group process the same partition simultaneously.
7. Distinguish broker delivery semantics from application effect semantics.
8. State the exact scope of any exactly-once claim.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Distinguish queue delivery, consumer processing, acknowledgment, and checkpoint state.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Explain why at-least-once delivery can repeat a message after a crash.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Make a consumer effect idempotent or atomically coordinate the effect with deduplication state.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

Assume messages can be delayed, duplicated, or lost. State the guarantees you need, then design the minimum protocol state required to preserve them.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### Event-processing service

Build a bounded event pipeline with producers, consumers, retries, dead-letter handling, and explicit at-least-once semantics.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. In your step 4 deduplication mechanism, deliberately introduce the same ordering bug distributed database-internals Lesson 6's WAL failure drill covered: record the "already processed" marker *before* actually applying the operation's effect (reversed from the correct order), then simulate a crash occurring between recording the marker and actually applying.

2. Confirm that after restart, the message is now considered "already processed" (so it will never be redelivered or reprocessed) even though its actual effect was never applied — a silently *lost* operation, which is arguably worse than the duplicate this lesson has otherwise been preventing, since a duplicate is at least visible/detectable while a silently sk.

3. Explain why the correct ordering (apply the effect, then durably record it as processed — mirroring database-internals Lesson 6's "write-ahead" ordering principle, just applied in the opposite temporal direction here: record only after the effect is safely applied, not before) is what avoids this specific new failure mode.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one effect or acknowledgment boundary: use an idempotent upsert instead of an append, or commit the consumer offset with a different transaction boundary. Explain which duplicate/loss window changes.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Model:** State the system model and assumptions before solving.
2. **Trace:** Work a small failure or state-transition example by hand.
3. **Trade-off:** Compare two designs and state the condition under which each wins.
4. **Evidence:** Identify the measurement, test, or trace that would distinguish the designs in practice.

## 11. Self-test

- [ ] What is at-least-once delivery?
- [ ] What crash window can cause duplicate processing?
- [ ] What crash window can cause a lost effect?
- [ ] When can an effect and dedupe record be atomic?
- [ ] What ordering does a partition provide?
- [ ] Why is exactly-once delivery not automatically end-to-end exactly-once processing?

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
