# Lesson 038 — Recursion, Dynamic Programming, and Recurrences Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement naive recursive Fibonacci and draw its call tree for a small argument.

2. Add memoization and bottom-up tabulation; explain which repeated computation is removed.

3. Derive and solve the recurrences for binary search and merge sort, stating theorem assumptions.

## Run and check

```powershell
Set-Location lessons/S1/038-recursion-dp-and-recurrences/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Dynamic-programming solver

Solve a problem with both a recursive baseline and a memoized or tabulated version. Compare call counts and explain the recurrence.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
