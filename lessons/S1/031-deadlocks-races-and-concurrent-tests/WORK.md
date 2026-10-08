# Lesson 031 — Deadlocks, Races, and Concurrent Tests Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Write the two-thread/two-lock schedule and draw its wait-for cycle.

2. Change the implementation to acquire locks in one global order and explain which Coffman condition is broken.

3. Design a concurrency test with explicit barriers/channels to control the relevant schedule and a timeout only as a bounded failure guard.

## Run and check

```powershell
Set-Location lessons/S1/031-deadlocks-races-and-concurrent-tests/code
go test -race exercise.go exercise_test.go
```

## Optional challenge

### Deadlock lab

Construct a two-lock deadlock with controlled timing, detect it, then fix it using a lock-order discipline or different ownership model.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
