# Lesson 070 — Problem Set

**Topic:** Rate Limits, Caching, and Capacity.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a fixed-window counter and reproduce the boundary burst with a controllable clock.

2. Implement a token bucket with capacity, refill rate, and last-refill timestamp.

3. Test burst, idle refill, empty bucket, shared-key concurrency, and separate logical client identities.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/070-rate-limits-caching-and-capacity/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Rate-limit + cache

Add a bounded rate limiter and a cache to a small HTTP service. Measure hit rate and rejection behavior, then document the consistency trade-off.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
