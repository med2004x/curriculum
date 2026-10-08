// STUDENT PROBLEM SET — Lesson 089.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.
// 2. Extend it with one explicit boundary case: Test basic replication correctness: perform a sequence of writes against the leader, confirm each follower's local state exactly matches the leader's after all operations have propagated and been applied.
// 3. Rebuild the same behavior with one changed constraint: Simulate a follower crash and recovery: stop a follower process partway through a sequence of leader operations (so it misses some), then restart it and implement a simple catch-up mechanism (the follower requests all log entries after the last one it successfully applied, and replays them in order) — confirm the recovered follower's state correctly converge. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 089: Replication and State Machines
// Your task: define a deterministic state transition and reject invalid log order.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

type Event struct {
	Index int
	Value int
}

func apply(state int, e Event) int {
    // TODO: implement the assigned behavior.
    return 0
}

func main() {}
