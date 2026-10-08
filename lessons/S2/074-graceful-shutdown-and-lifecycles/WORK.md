# Lesson 074 — Graceful Shutdown and Lifecycles Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a server that receives shutdown cancellation and stops accepting new work.
2. Add readiness transition, worker cancellation/join, and a bounded shutdown deadline.
3. Write a test that observes worker completion using synchronization rather than an arbitrary sleep.

## Run and check

```powershell
Set-Location lessons/S2/074-graceful-shutdown-and-lifecycles/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Graceful shutdown

Build a service that stops admission, drains in-flight work, cancels workers, closes resources, and respects one bounded shutdown deadline.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
