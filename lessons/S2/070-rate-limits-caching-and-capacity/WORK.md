# Lesson 070 — Rate Limits, Caching, and Capacity Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a fixed-window counter and reproduce the boundary burst with a controllable clock.

2. Implement a token bucket with capacity, refill rate, and last-refill timestamp.

3. Test burst, idle refill, empty bucket, shared-key concurrency, and separate logical client identities.

## Run and check

```powershell
Set-Location lessons/S2/070-rate-limits-caching-and-capacity/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Rate-limit + cache

Add a bounded rate limiter and a cache to a small HTTP service. Measure hit rate and rejection behavior, then document the consistency trade-off.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
