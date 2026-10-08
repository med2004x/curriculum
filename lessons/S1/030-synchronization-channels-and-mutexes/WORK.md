# Lesson 030 — Problem Set

**Topic:** Synchronization, Channels, and Mutexes.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a fixed-capacity ring buffer with `count`, head, and tail state protected by one mutex.

2. Coordinate empty/full transitions using two condition variables and predicate loops, or implement the same bounded flow with a buffered Go channel.

3. Run multiple producers and consumers over uniquely identified finite items; verify conservation and termination.

4. Build a separate two-lock example with opposite acquisition order, draw the wait-for graph, then enforce a global order.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/030-synchronization-channels-and-mutexes/code
go test -race exercise.go exercise_test.go
```

## Stretch problem

### Concurrent pipeline

Build a three-stage Go pipeline using channels and cancellation. Prove that every goroutine has an exit path and that no stage can block forever.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
