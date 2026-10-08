# Lesson 089 — Replication and State Machines Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.

2. Test basic replication correctness: perform a sequence of writes against the leader, confirm each follower's local state exactly matches the leader's after all operations have propagated and been applied.

3. Simulate a follower crash and recovery: stop a follower process partway through a sequence of leader operations (so it misses some), then restart it and implement a simple catch-up mechanism (the follower requests all log entries after the last one it successfully applied, and replays them in order) — confirm the recovered follower's state correctly converge.

## Run and check

```powershell
Set-Location lessons/S2/089-replication-and-state-machines/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Replica state machine

Build a small replicated state-machine simulator with an append-only command log and deterministic application of committed entries.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
