# Lesson 092 — Queues, Streams, and Event Processing Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before.

2. Confirm that if you kill the consumer process mid-processing (before it advances the checkpoint) and restart it, it correctly redelivers/reprocesses the message it was working on when killed, rather than skipping it.

3. Implement a naturally-idempotent operation as your consumer's processing logic (e.g. "set key K's value to V" in a key-value store) and confirm that redelivery/reprocessing (per step 1's crash-and-restart test) produces no observable problem — the final state is correct either way, since reapplying the same "set" operation is harmless by the operation's own.

## Run and check

```powershell
Set-Location lessons/S2/092-queues-streams-and-event-processing/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Event-processing service

Build a bounded event pipeline with producers, consumers, retries, dead-letter handling, and explicit at-least-once semantics.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
