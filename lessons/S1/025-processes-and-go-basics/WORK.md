# Lesson 025 — Processes and Your First Go Program Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Write a C program (or use `fork()` via a small wrapper if working in Go, which exposes `os/exec` instead — pick whichever language makes the fork/exec split easiest to observe directly; C is more instructive here since it exposes both syscalls explicitly) that calls `fork()`, and in the child, calls `execve` (or `execvp`) to run `/bin/ls`.

2. In the parent, call `wait()` to retrieve the child's exit status and print it.

3. Confirm the parent process's PID is unchanged before and after, while the child's PID (printed from inside the child, before `exec`) differs from the parent's.

## Run and check

```powershell
Set-Location lessons/S1/025-processes-and-go-basics/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Go process launcher

Write a Go command that starts a child process, captures its output and exit status, and handles a failed command cleanly.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
