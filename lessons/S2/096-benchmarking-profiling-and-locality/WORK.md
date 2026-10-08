# Lesson 096 — Problem Set

**Topic:** Benchmarking, Profiling, and Locality.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Benchmark two implementations of the same operation.

2. Extend it with one explicit boundary case: Repeat enough times to report a distribution, not one favorite number.

3. Rebuild the same behavior with one changed constraint: Write (or use) a deliberately inefficient Go function — e.g. one doing unnecessary repeated string concatenation in a loop (a classic Go performance anti-pattern, since naive `+=` string concatenation reallocates on every iteration) — and benchmark it using Lesson 1's methodology to establish a baseline.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/096-benchmarking-profiling-and-locality/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Performance investigation

Benchmark two implementations, profile the slower one, identify one hot path, change it, and verify the improvement without changing semantics.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
