# Lesson 034 — Go Methods, Interfaces, and Abstraction Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Define a small `Payer` interface and a value-receiver implementation.

2. Add a second implementation and make one caller depend only on `Payer`.

3. Change the receiver to a pointer receiver and predict interface satisfaction for `T` and `*T`.

## Run and check

```powershell
Set-Location lessons/S1/034-go-methods-interfaces-and-abstraction/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Interface boundary

Define a narrow interface for a dependency, implement two concrete backends, and prove through tests that the service depends only on behavior, not representation.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
