# Lesson 030 — Synchronization, Channels, and Mutexes Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a fixed-capacity ring buffer with `count`, head, and tail state protected by one mutex.

2. Coordinate empty/full transitions using two condition variables and predicate loops, or implement the same bounded flow with a buffered Go channel.

3. Run multiple producers and consumers over uniquely identified finite items; verify conservation and termination.

4. Build a separate two-lock example with opposite acquisition order, draw the wait-for graph, then enforce a global order.

## Run and check

```powershell
Set-Location lessons/S1/030-synchronization-channels-and-mutexes/code
go test -race exercise.go exercise_test.go
```

## Optional challenge

### Concurrent pipeline

Build a three-stage Go pipeline using channels and cancellation. Prove that every goroutine has an exit path and that no stage can block forever.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
