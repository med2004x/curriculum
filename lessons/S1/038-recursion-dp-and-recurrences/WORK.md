# Lesson 038 — Problem Set

**Topic:** Recursion, Dynamic Programming, and Recurrences.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement naive recursive Fibonacci and draw its call tree for a small argument.

2. Add memoization and bottom-up tabulation; explain which repeated computation is removed.

3. Derive and solve the recurrences for binary search and merge sort, stating theorem assumptions.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/038-recursion-dp-and-recurrences/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Dynamic-programming solver

Solve a problem with both a recursive baseline and a memoized or tabulated version. Compare call counts and explain the recurrence.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
