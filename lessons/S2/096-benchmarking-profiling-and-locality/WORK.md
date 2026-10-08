# Lesson 096 — Benchmarking, Profiling, and Locality Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Benchmark two implementations of the same operation.

2. Repeat enough times to report a distribution, not one favorite number.

3. Write (or use) a deliberately inefficient Go function — e.g. one doing unnecessary repeated string concatenation in a loop (a classic Go performance anti-pattern, since naive `+=` string concatenation reallocates on every iteration) — and benchmark it using Lesson 1's methodology to establish a baseline.

## Run and check

```powershell
Set-Location lessons/S2/096-benchmarking-profiling-and-locality/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Performance investigation

Benchmark two implementations, profile the slower one, identify one hot path, change it, and verify the improvement without changing semantics.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
