# Lesson 035 — Generics and Type Constraints Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement `contains` for strings and integers first.

2. Test empty, singleton, duplicate, found, and missing-value cases for both types.

3. Write `Contains[T comparable]` and explain why the constraint is necessary for equality.

## Run and check

```powershell
Set-Location lessons/S1/035-go-generics-and-type-constraints/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Generic utility

Implement a generic collection helper with a meaningful type constraint. Demonstrate one useful specialization and explain why the constraint is necessary.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
