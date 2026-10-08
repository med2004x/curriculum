# Lesson 029 — Threads, Goroutines, and Data Races Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Write a Go program with several goroutines incrementing a shared counter without synchronization.

2. Add the race detector and explain the reported conflicting read/write operations.

3. Protect the counter with a mutex and make every read of the protected state use the same mutex.

## Run and check

```powershell
Set-Location lessons/S1/029-threads-goroutines-and-races/code
go test -race exercise.go exercise_test.go
```

## Optional challenge

### Race reproduction

Create a tiny concurrent Go counter that fails under `-race`, then redesign it using a synchronization strategy justified by the access pattern.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
