# Lesson 029 — Problem Set

**Topic:** Threads, Goroutines, and Data Races.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Write a Go program with several goroutines incrementing a shared counter without synchronization.

2. Add the race detector and explain the reported conflicting read/write operations.

3. Protect the counter with a mutex and make every read of the protected state use the same mutex.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/029-threads-goroutines-and-races/code
go test -race exercise.go exercise_test.go
```

## Stretch problem

### Race reproduction

Create a tiny concurrent Go counter that fails under `-race`, then redesign it using a synchronization strategy justified by the access pattern.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
