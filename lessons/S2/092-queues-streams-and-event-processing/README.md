# Lesson 092 — Queues, Streams, and Event Processing

## Threads
**Primary:** Distributed/Go

**Supporting:** delivery semantics; consumers; offsets

**Math companion:** random variables and distributions

**Time:** 3–4 hours. **Prerequisites:** Lesson 091.

**Time:** 3–4 hours. **Prerequisites:** Lesson 091.

## What you will be able to do

1. Distinguish queue delivery, consumer processing, acknowledgment, and checkpoint state.
2. Explain why at-least-once delivery can repeat a message after a crash.
3. Make a consumer effect idempotent or atomically coordinate the effect with deduplication state.
4. Distinguish consumer groups, partitions, offsets, and ordering scope.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 091, # Lesson 091 — Partitioning and Sharding.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Process durable messages despite consumer crashes, retries, and redelivery while preserving the intended effect and progress position.

**The contract.** A broker may provide at-least-once delivery: unacknowledged messages can be redelivered. If the consumer's effect and offset/acknowledgment are not atomic, a crash between them can produce duplicates or lost effects. “Exactly once” is scoped to particular transactional boundaries; it is not a general end-to-end property across arbitrary external side effects.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** A consumer reads at an offset, performs its effect, then commits progress/acknowledges. If it crashes after the effect but before progress is durable, the message can be processed again. If it records progress before the effect and crashes, the effect can be lost.

**Second point.** To avoid duplicate database effects, store the event ID in a deduplication table and apply the domain change in the same database transaction. For external systems, use their idempotency facilities or an outbox/inbox protocol; a local checkpoint alone cannot make unrelated systems atomic.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Confirm that if you kill the consumer process mid-processing (before it advances the checkpoint) and restart it, it correctly redelivers/reprocesses the message it was working on when killed, rather than skipping it.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Simulate a crash after effect/before acknowledgment and after acknowledgment/before effect. Record delivery count, effect count, committed offset, and recovery result. Then test duplicate event IDs with effect and dedupe record committed atomically.

**Failure mode.** Record the dedupe marker before applying the effect in separate commits, then simulate a crash between them. Explain why redelivery is suppressed even though the effect never happened. Repair by using one transaction for local DB effect plus dedupe state.

**Completion criterion.** You can identify duplicate/loss windows, make a consumer effect idempotent or transactionally deduplicated, and state the exact boundary within which an exactly-once claim applies.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict what happens if a consumer crashes after effect but before acknowledgment.
2. Predict what happens if it commits the offset before performing the effect.
3. Predict how an idempotent `set` differs from a non-idempotent `append` on redelivery.
4. Predict what atomic DB deduplication prevents.
5. Predict the ordering guarantee within one partition.
6. Predict whether two consumers in one consumer group process the same partition simultaneously.
7. Distinguish broker delivery semantics from application effect semantics.
8. State the exact scope of any exactly-once claim.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before.

2. Extend it with one explicit boundary case: Confirm that if you kill the consumer process mid-processing (before it advances the checkpoint) and restart it, it correctly redelivers/reprocesses the message it was working on when killed, rather than skipping it.

3. Rebuild the same behavior with one changed constraint: Implement a naturally-idempotent operation as your consumer's processing logic (e.g. "set key K's value to V" in a key-value store) and confirm that redelivery/reprocessing (per step 1's crash-and-restart test) produces no observable problem — the final state is correct either way, since reapplying the same "set" operation is harmless by the operation's own.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. In your step 4 deduplication mechanism, deliberately introduce the same ordering bug distributed database-internals Lesson 6's WAL failure drill covered: record the "already processed" marker *before* actually applying the operation's effect (reversed from the correct order), then simulate a crash occurring between recording the marker and actually applying.

2. Confirm that after restart, the message is now considered "already processed" (so it will never be redelivered or reprocessed) even though its actual effect was never applied — a silently *lost* operation, which is arguably worse than the duplicate this lesson has otherwise been preventing, since a duplicate is at least visible/detectable while a silently sk.

3. Explain why the correct ordering (apply the effect, then durably record it as processed — mirroring database-internals Lesson 6's "write-ahead" ordering principle, just applied in the opposite temporal direction here: record only after the effect is safely applied, not before) is what avoids this specific new failure mode.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one effect or acknowledgment boundary: use an idempotent upsert instead of an append, or commit the consumer offset with a different transaction boundary. Explain which duplicate/loss window changes.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is at-least-once delivery?
2. What crash window can cause duplicate processing?
3. What crash window can cause a lost effect?
4. When can an effect and dedupe record be atomic?
5. What ordering does a partition provide?
6. Why is exactly-once delivery not automatically end-to-end exactly-once processing?

## Evidence to record

Create `lesson092/EVIDENCE.md` with exactly this structure:

```text
# Lesson 092 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 092 — Queues, Streams, and Event Processing.
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

You pass when you can explain the central model of # 092 — Queues, Streams, and Event Processing without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've directly demonstrated at-least-once delivery redelivering a message after a simulated consumer crash, you've shown the concrete difference in outcome between a naturally-idempotent and a non-idempotent operation under that same redelivery, and you've implemented and verified explicit deduplication fixing the non-idempotent case — while also demonstrat. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
