# Lesson 089 — Problem Set

**Topic:** Replication and State Machines.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.

2. Extend it with one explicit boundary case: Test basic replication correctness: perform a sequence of writes against the leader, confirm each follower's local state exactly matches the leader's after all operations have propagated and been applied.

3. Rebuild the same behavior with one changed constraint: Simulate a follower crash and recovery: stop a follower process partway through a sequence of leader operations (so it misses some), then restart it and implement a simple catch-up mechanism (the follower requests all log entries after the last one it successfully applied, and replays them in order) — confirm the recovered follower's state correctly converge.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/089-replication-and-state-machines/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Replica state machine

Build a small replicated state-machine simulator with an append-only command log and deterministic application of committed entries.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
