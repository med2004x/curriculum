# Lesson 094 — Problem Set

**Topic:** Timeouts, Retries, Backpressure, and Recovery.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Draw a call chain and assign one end-to-end deadline across its hops.

2. Add bounded exponential retries with jitter only for a transient/idempotent operation.

3. Add a bounded queue or admission limit and define the overload response.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/094-timeouts-retries-backpressure-and-recovery/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Resilient client

Implement a client with deadlines, bounded retries, exponential backoff, and backpressure. Test slow, flaky, and permanently failing dependencies.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
