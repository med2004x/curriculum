# Lesson 092 — Problem Set

**Topic:** Queues, Streams, and Event Processing.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before.

2. Extend it with one explicit boundary case: Confirm that if you kill the consumer process mid-processing (before it advances the checkpoint) and restart it, it correctly redelivers/reprocesses the message it was working on when killed, rather than skipping it.

3. Rebuild the same behavior with one changed constraint: Implement a naturally-idempotent operation as your consumer's processing logic (e.g. "set key K's value to V" in a key-value store) and confirm that redelivery/reprocessing (per step 1's crash-and-restart test) produces no observable problem — the final state is correct either way, since reapplying the same "set" operation is harmless by the operation's own.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/092-queues-streams-and-event-processing/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Event-processing service

Build a bounded event pipeline with producers, consumers, retries, dead-letter handling, and explicit at-least-once semantics.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
