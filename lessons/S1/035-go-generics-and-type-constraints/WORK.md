# Lesson 035 — Problem Set

**Topic:** Generics and Type Constraints.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement `contains` for strings and integers first.

2. Test empty, singleton, duplicate, found, and missing-value cases for both types.

3. Write `Contains[T comparable]` and explain why the constraint is necessary for equality.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/035-go-generics-and-type-constraints/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Generic utility

Implement a generic collection helper with a meaningful type constraint. Demonstrate one useful specialization and explain why the constraint is necessary.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
