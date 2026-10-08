# Lesson 026 — File Descriptors and Go I/O Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Build a small C program that writes to stdout and stderr.

2. Redirect each stream separately.

3. Then connect two processes with a pipe and describe the byte flow.

## Run and check

```powershell
Set-Location lessons/S1/026-file-descriptors-and-go-io/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### File copier with boundaries

Build a Go file-copy utility with explicit buffering, error handling, and close-order checks. Test empty, large, and unreadable inputs.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
