# Lesson 033 — Go Slices, Maps, Structs, and Data Structures Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Define an `Account` struct and an update function that returns a modified value.

2. Store multiple accounts in a slice and aggregate balances in a map keyed by owner.

3. Construct two slices sharing a backing array; record the results of reslicing and appending within and beyond capacity.

## Run and check

```powershell
Set-Location lessons/S1/033-go-slices-maps-structs-and-data/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Collection-backed service

Implement a Go service around slices, maps, and structs. Add lookup, update, deletion, and deterministic listing with tests for missing and duplicate keys.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
