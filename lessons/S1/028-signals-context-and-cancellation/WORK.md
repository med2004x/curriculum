# Lesson 028 — Problem Set

**Topic:** Signals, Context, and Cancellation.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a worker that receives a context and exits when `ctx.Done()` closes.

2. Add a deadline and propagate the context through at least two function layers.

3. Connect `SIGTERM` to cancellation with `signal.NotifyContext` and give cleanup a bounded wait.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/028-signals-context-and-cancellation/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Cancellable worker

Implement a Go worker that consumes jobs until context cancellation, exits promptly, and reports how many jobs completed before shutdown.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
