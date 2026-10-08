# Lesson 033 — Problem Set

**Topic:** Go Slices, Maps, Structs, and Data Structures.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define an `Account` struct and an update function that returns a modified value.

2. Store multiple accounts in a slice and aggregate balances in a map keyed by owner.

3. Construct two slices sharing a backing array; record the results of reslicing and appending within and beyond capacity.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/033-go-slices-maps-structs-and-data/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Collection-backed service

Implement a Go service around slices, maps, and structs. Add lookup, update, deletion, and deterministic listing with tests for missing and duplicate keys.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
