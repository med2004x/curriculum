# Lesson 028 — Signals, Context, and Cancellation Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a worker that receives a context and exits when `ctx.Done()` closes.

2. Add a deadline and propagate the context through at least two function layers.

3. Connect `SIGTERM` to cancellation with `signal.NotifyContext` and give cleanup a bounded wait.

## Run and check

```powershell
Set-Location lessons/S1/028-signals-context-and-cancellation/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Cancellable worker

Implement a Go worker that consumes jobs until context cancellation, exits promptly, and reports how many jobs completed before shutdown.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
