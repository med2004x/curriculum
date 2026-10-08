# Lesson 074 — Problem Set

**Topic:** Graceful Shutdown and Lifecycles.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a server that receives shutdown cancellation and stops accepting new work.
2. Add readiness transition, worker cancellation/join, and a bounded shutdown deadline.
3. Write a test that observes worker completion using synchronization rather than an arbitrary sleep.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/074-graceful-shutdown-and-lifecycles/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Graceful shutdown

Build a service that stops admission, drains in-flight work, cancels workers, closes resources, and respects one bounded shutdown deadline.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
