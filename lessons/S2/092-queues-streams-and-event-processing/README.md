# Lesson 092 — Queues, Streams, and Event Processing

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Process durable messages despite consumer crashes, retries, and redelivery while preserving the intended effect and progress position.

**The contract.** A broker may provide at-least-once delivery: unacknowledged messages can be redelivered. If the consumer's effect and offset/acknowledgment are not atomic, a crash between them can produce duplicates or lost effects. “Exactly once” is scoped to particular transactional boundaries; it is not a general end-to-end property across arbitrary external side effects.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before.

## 2. Mechanism and vocabulary

**Mechanism.** A consumer reads at an offset, performs its effect, then commits progress/acknowledges. If it crashes after the effect but before progress is durable, the message can be processed again. If it records progress before the effect and crashes, the effect can be lost.

**Second point.** To avoid duplicate database effects, store the event ID in a deduplication table and apply the domain change in the same database transaction. For external systems, use their idempotency facilities or an outbox/inbox protocol; a local checkpoint alone cannot make unrelated systems atomic.

**Worked example.** Confirm that if you kill the consumer process mid-processing (before it advances the checkpoint) and restart it, it correctly redelivers/reprocesses the message it was working on when killed, rather than skipping it.

## 3. Worked case from the reference material

Worked example: Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before.

## 4. Verification, failure, and evidence

**Verification.** Simulate a crash after effect/before acknowledgment and after acknowledgment/before effect. Record delivery count, effect count, committed offset, and recovery result. Then test duplicate event IDs with effect and dedupe record committed atomically.

**Failure mode.** Record the dedupe marker before applying the effect in separate commits, then simulate a crash between them. Explain why redelivery is suppressed even though the effect never happened. Repair by using one transaction for local DB effect plus dedupe state.

**Completion criterion.** You can identify duplicate/loss windows, make a consumer effect idempotent or transactionally deduplicated, and state the exact boundary within which an exactly-once claim applies.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
