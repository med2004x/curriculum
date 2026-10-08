# Lesson 025 — Problem Set

**Topic:** Processes and Your First Go Program.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Write a C program (or use `fork()` via a small wrapper if working in Go, which exposes `os/exec` instead — pick whichever language makes the fork/exec split easiest to observe directly; C is more instructive here since it exposes both syscalls explicitly) that calls `fork()`, and in the child, calls `execve` (or `execvp`) to run `/bin/ls`.

2. Extend it with one explicit boundary case: In the parent, call `wait()` to retrieve the child's exit status and print it.

3. Rebuild the same behavior with one changed constraint: Confirm the parent process's PID is unchanged before and after, while the child's PID (printed from inside the child, before `exec`) differs from the parent's.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/025-processes-and-go-basics/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Go process launcher

Write a Go command that starts a child process, captures its output and exit status, and handles a failed command cleanly.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
