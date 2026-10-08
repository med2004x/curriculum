// STUDENT PROBLEM SET — Lesson 092.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before.
// 2. Extend it with one explicit boundary case: Confirm that if you kill the consumer process mid-processing (before it advances the checkpoint) and restart it, it correctly redelivers/reprocesses the message it was working on when killed, rather than skipping it.
// 3. Rebuild the same behavior with one changed constraint: Implement a naturally-idempotent operation as your consumer's processing logic (e.g. "set key K's value to V" in a key-value store) and confirm that redelivery/reprocessing (per step 1's crash-and-restart test) produces no observable problem — the final state is correct either way, since reapplying the same "set" operation is harmless by the operation's own. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 092: Queues, Streams, and Event Processing
// Your task: append without violating ordering or ownership assumptions.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

type Message struct {
	Topic string
	Value string
}

func publish(topics map[string][]Message, msg Message) {
    // TODO: implement the assigned behavior.
}

func main() {}
