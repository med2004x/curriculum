# Lesson 025 — Processes and Your First Go Program

## Threads
**Primary:** Linux/Go

**Supporting:** process lifecycle; Go syntax; executable model

**Math companion:** probability review

**Time:** 3–4 hours. **Prerequisites:** Lesson 024.

**Time:** 3–4 hours. **Prerequisites:** Lesson 024.

## What you will be able to do

1. Explain process creation, program replacement, waiting, and exit status.
2. Distinguish `fork` from the `exec` family and describe what state is inherited.
3. Write a small Go program with an explicit entry point, values, and error handling.
4. Distinguish catchable `SIGTERM` from uncatchable `SIGKILL` and state why cleanup must be bounded.

## Retrieve first (no notes)

1. What is a virtual address space?
2. What does a page table translate?
3. What is the difference between a serviced page fault and an invalid access?
4. What is the difference between a process's virtual size and its RSS?

## 1. Problem and model

**What problem this solves.** Understand how a running process starts another program and how the parent learns whether the child completed successfully. The Go portion introduces the language's program structure and explicit error handling; the signal exercise is a separate lifecycle extension.

**The contract.** On POSIX systems, `fork()` creates a child process with a distinct PID and inherited process state; implementations commonly use copy-on-write rather than eagerly copying all memory. An `exec` function replaces the calling process image with a new program and does not create a new PID. The parent uses `wait`/`waitpid` to collect child state and avoid leaving a terminated child unreaped.

**Worked example.** Write a small C program that calls `fork()`. The child calls `execvp` with a program name and argument vector; if `execvp` returns, it failed, so report the error and exit with failure. The parent calls `waitpid`, checks that the child exited normally, and extracts the child's exit code. Go's `os/exec` offers a higher-level process API; it is not a direct Go wrapper around `fork` followed by user-managed `exec`.

Self-check: which PID changes after `fork`, and what is replaced by a successful `exec`?

## 2. Mechanism and vocabulary

**Mechanism.** `fork` returns twice: zero in the child and the child's PID in the parent; failure returns -1 in the parent. Successful `exec` does not return. The child inherits open descriptors unless their close-on-exec flag is set; shells can use this interval to set up redirection before executing the command.

**Second point.** A Go program begins in `main.main` within package `main`; short declarations such as `count := 3` introduce a local variable, and functions that can fail conventionally return an `error` for the caller to inspect. Do not discard an error merely to keep an example short.

Signals are a distinct process-control mechanism: a process may install a handler for `SIGTERM` and choose a bounded cleanup response. `SIGKILL` cannot be caught, blocked, or ignored. A handler that waits indefinitely can still be terminated after an external grace period; graceful shutdown therefore needs a deadline.

**Worked example.** In the parent, call `wait()` to retrieve the child's exit status and print it.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: what does the process API report if the requested executable cannot be found?

## 4. Verification, failure, and evidence

**Verification.** For process creation, record parent PID, child PID, whether `exec` succeeded, and the decoded wait status for both a successful command and a command that exits nonzero. For the separate signal exercise, record the signal sent and whether cleanup ran; do not expect a `SIGKILL` handler.

**Failure mode.** Make the child request a nonexistent executable and handle the failed `exec` path. In the signal extension, let `SIGTERM` cleanup exceed a supervisor's grace period and explain why a later `SIGKILL` cannot be handled.

**Completion criterion.** You can explain fork/exec/wait, check the child's exit result, write the introductory Go example, and describe bounded graceful shutdown without treating signals as interchangeable.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the return values of `fork` in the parent and child.
2. Predict whether a successful `exec` returns to the old program.
3. Predict which PID remains associated with the process after `exec`.
4. Predict the wait result when the child exits with status 7.
5. Predict what happens when the executable path is invalid.
6. Predict which file descriptors survive `exec` by default and what close-on-exec changes.
7. Predict whether a `SIGTERM` handler can run after `SIGKILL` is sent.
8. Record the process transition and status expected for each case.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Write a C program (or use `fork()` via a small wrapper if working in Go, which exposes `os/exec` instead — pick whichever language makes the fork/exec split easiest to observe directly; C is more instructive here since it exposes both syscalls explicitly) that calls `fork()`, and in the child, calls `execve` (or `execvp`) to run `/bin/ls`.

2. Extend it with one explicit boundary case: In the parent, call `wait()` to retrieve the child's exit status and print it.

3. Rebuild the same behavior with one changed constraint: Confirm the parent process's PID is unchanged before and after, while the child's PID (printed from inside the child, before `exec`) differs from the parent's.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Modify step 3's `SIGTERM` handler to deliberately take a long time (e.g.

2. `sleep(30)` inside the handler, simulating slow cleanup like draining a large in-flight request queue) before exiting.

3. Send `SIGTERM`, then — simulating what a container orchestrator does after its grace period expires — send `SIGKILL` a few seconds later, before the handler finishes.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint: use `os/exec` in Go instead of explicit C `fork`/`exec`, or change the child program and its arguments. Explain which low-level process details the Go API manages for you.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does `fork` do, and how does its return value differ between parent and child?
2. What does successful `exec` replace, and does it change the PID?
3. How does a parent retrieve a child's exit status?
4. How does `os/exec` relate to low-level `fork` and `exec`?
5. Which signals can and cannot be caught?
6. Why should graceful cleanup have a deadline?

## Evidence to record

Create `lesson025/EVIDENCE.md` with exactly this structure:

```text
# Lesson 025 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Saying that `exec` creates a new process; it replaces the current process image.
- Assuming a successful `exec` returns to the caller.
- Ignoring the child's wait status or confusing it with the PID.
- Treating Go's `os/exec` as a direct, portable `fork` interface.
- Expecting a `SIGKILL` handler or unbounded cleanup after `SIGTERM`.

## Precision checks

1. Draw the `fork` → child setup → `exec` → parent `waitpid` sequence.
2. State which process ID changes or remains stable at each step.
3. Explain how a nonzero child exit is represented to the parent.
4. Name the Go entry point and explain why errors should be checked.
5. Explain why `SIGKILL` cannot perform application cleanup.

Keep process creation, program replacement, waiting, and signal delivery as distinct transitions; each has a different contract.
## University revision bridge

For the LIG track, connect process lifecycle and signal semantics to the OS model, and treat Go's entry-point syntax as a separate language foundation.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain and trace fork/exec/wait, inspect a child exit status, write the introductory Go program, and state the operational difference between `SIGTERM` and `SIGKILL`.
