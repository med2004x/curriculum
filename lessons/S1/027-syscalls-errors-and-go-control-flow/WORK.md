# Lesson 027 — Problem Set

**Topic:** Syscalls, Errors, and Explicit Control Flow.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Run a sleeping process and inspect its PID, parent PID, open descriptors, and command line.

2. Extend it with one explicit boundary case: Send a signal and record the state transition.

3. Rebuild the same behavior with one changed constraint: Trace a program that opens and reads a file.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/027-syscalls-errors-and-go-control-flow/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Syscall-style wrapper

Design a Go API that turns low-level failures into explicit typed or sentinel errors. Exercise success, missing resource, and malformed input cases.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
