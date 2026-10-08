# Lesson 027 — Syscalls, Errors, and Explicit Control Flow Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Run a sleeping process and inspect its PID, parent PID, open descriptors, and command line.

2. Send a signal and record the state transition.

3. Trace a program that opens and reads a file.

## Run and check

```powershell
Set-Location lessons/S1/027-syscalls-errors-and-go-control-flow/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Syscall-style wrapper

Design a Go API that turns low-level failures into explicit typed or sentinel errors. Exercise success, missing resource, and malformed input cases.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
