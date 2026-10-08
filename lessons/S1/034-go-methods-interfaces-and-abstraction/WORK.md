# Lesson 034 — Problem Set

**Topic:** Go Methods, Interfaces, and Abstraction.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define a small `Payer` interface and a value-receiver implementation.

2. Add a second implementation and make one caller depend only on `Payer`.

3. Change the receiver to a pointer receiver and predict interface satisfaction for `T` and `*T`.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/034-go-methods-interfaces-and-abstraction/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Interface boundary

Define a narrow interface for a dependency, implement two concrete backends, and prove through tests that the service depends only on behavior, not representation.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
