# Lesson 025 — Processes and Your First Go Program

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand how a running process starts another program and how the parent learns whether the child completed successfully. The Go portion introduces the language's program structure and explicit error handling; the signal exercise is a separate lifecycle extension.

**The contract.** On POSIX systems, `fork()` creates a child process with a distinct PID and inherited process state; implementations commonly use copy-on-write rather than eagerly copying all memory. An `exec` function replaces the calling process image with a new program and does not create a new PID. The parent uses `wait`/`waitpid` to collect child state and avoid leaving a terminated child unreaped.

**Worked example.** Write a small C program that calls `fork()`. The child calls `execvp` with a program name and argument vector; if `execvp` returns, it failed, so report the error and exit with failure. The parent calls `waitpid`, checks that the child exited normally, and extracts the child's exit code. Go's `os/exec` offers a higher-level process API; it is not a direct Go wrapper around `fork` followed by user-managed `exec`.

## 2. Mechanism and vocabulary

**Mechanism.** `fork` returns twice: zero in the child and the child's PID in the parent; failure returns -1 in the parent. Successful `exec` does not return. The child inherits open descriptors unless their close-on-exec flag is set; shells can use this interval to set up redirection before executing the command.

**Second point.** A Go program begins in `main.main` within package `main`; short declarations such as `count := 3` introduce a local variable, and functions that can fail conventionally return an `error` for the caller to inspect. Do not discard an error merely to keep an example short.

Signals are a distinct process-control mechanism: a process may install a handler for `SIGTERM` and choose a bounded cleanup response. `SIGKILL` cannot be caught, blocked, or ignored. A handler that waits indefinitely can still be terminated after an external grace period; graceful shutdown therefore needs a deadline.

**Worked example.** In the parent, call `wait()` to retrieve the child's exit status and print it.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
package main

import "fmt"

func main() {
    count := 3
    fmt.Println(count)
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

This tiny Go program illustrates `package main`, imports, `main`, short declaration, and output. Keep it separate from the C `fork`/`exec` exercise; Go's standard `os/exec` API is the appropriate high-level way to start a child program in ordinary Go code.

## 4. Verification, failure, and evidence

**Verification.** For process creation, record parent PID, child PID, whether `exec` succeeded, and the decoded wait status for both a successful command and a command that exits nonzero. For the separate signal exercise, record the signal sent and whether cleanup ran; do not expect a `SIGKILL` handler.

**Failure mode.** Make the child request a nonexistent executable and handle the failed `exec` path. In the signal extension, let `SIGTERM` cleanup exceed a supervisor's grace period and explain why a later `SIGKILL` cannot be handled.

**Completion criterion.** You can explain fork/exec/wait, check the child's exit result, write the introductory Go example, and describe bounded graceful shutdown without treating signals as interchangeable.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
