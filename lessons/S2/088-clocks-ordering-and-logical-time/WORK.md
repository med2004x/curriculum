# Lesson 088 — Clocks, Ordering, and Logical Time Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a Lamport clock as a small class/struct with a `tick()` method (increment for a local event, return new value) and a `receive(remoteTimestamp)` method (implementing the max-plus-one rule from Learn).

2. Write a unit test confirming the basic increment behavior in isolation.

3. Simulate 3 independent "processes" (just 3 instances of your Lamport clock in one program, or 3 actual goroutines/threads communicating via channels/sockets) that occasionally send messages to each other carrying their current Lamport timestamp, with the receiver updating its own clock per the receive rule.

## Run and check

```powershell
Set-Location lessons/S2/088-clocks-ordering-and-logical-time/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Logical ordering

Implement Lamport-style logical timestamps or another explicit ordering mechanism and demonstrate how concurrent events become comparable.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
