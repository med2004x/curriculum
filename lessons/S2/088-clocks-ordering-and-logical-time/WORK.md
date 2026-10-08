# Lesson 088 — Problem Set

**Topic:** Clocks, Ordering, and Logical Time.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement a Lamport clock as a small class/struct with a `tick()` method (increment for a local event, return new value) and a `receive(remoteTimestamp)` method (implementing the max-plus-one rule from Learn).

2. Extend it with one explicit boundary case: Write a unit test confirming the basic increment behavior in isolation.

3. Rebuild the same behavior with one changed constraint: Simulate 3 independent "processes" (just 3 instances of your Lamport clock in one program, or 3 actual goroutines/threads communicating via channels/sockets) that occasionally send messages to each other carrying their current Lamport timestamp, with the receiver updating its own clock per the receive rule.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/088-clocks-ordering-and-logical-time/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Logical ordering

Implement Lamport-style logical timestamps or another explicit ordering mechanism and demonstrate how concurrent events become comparable.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
