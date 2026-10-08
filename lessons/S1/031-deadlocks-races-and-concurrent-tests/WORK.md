# Lesson 031 — Problem Set

**Topic:** Deadlocks, Races, and Concurrent Tests.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Write the two-thread/two-lock schedule and draw its wait-for cycle.

2. Change the implementation to acquire locks in one global order and explain which Coffman condition is broken.

3. Design a concurrency test with explicit barriers/channels to control the relevant schedule and a timeout only as a bounded failure guard.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/031-deadlocks-races-and-concurrent-tests/code
go test -race exercise.go exercise_test.go
```

## Stretch problem

### Deadlock lab

Construct a two-lock deadlock with controlled timing, detect it, then fix it using a lock-order discipline or different ownership model.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
