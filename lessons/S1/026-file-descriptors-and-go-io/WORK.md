# Lesson 026 — Problem Set

**Topic:** File Descriptors and Go I/O.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Build a small C program that writes to stdout and stderr.

2. Extend it with one explicit boundary case: Redirect each stream separately.

3. Rebuild the same behavior with one changed constraint: Then connect two processes with a pipe and describe the byte flow.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/026-file-descriptors-and-go-io/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### File copier with boundaries

Build a Go file-copy utility with explicit buffering, error handling, and close-order checks. Test empty, large, and unreadable inputs.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
