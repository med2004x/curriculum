# Lesson 094 — Timeouts, Retries, Backpressure, and Recovery Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Draw a call chain and assign one end-to-end deadline across its hops.

2. Add bounded exponential retries with jitter only for a transient/idempotent operation.

3. Add a bounded queue or admission limit and define the overload response.

## Run and check

```powershell
Set-Location lessons/S2/094-timeouts-retries-backpressure-and-recovery/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Resilient client

Implement a client with deadlines, bounded retries, exponential backoff, and backpressure. Test slow, flaky, and permanently failing dependencies.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
