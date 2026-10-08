# Lesson 027 — Syscalls, Errors, and Explicit Control Flow

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Observe process state from Linux instead of inferring it from program output alone.

**The contract.** Linux exposes selected process information through pseudo-files under `/proc/<pid>`, including `status`, `cmdline`, and descriptor links under `fd/`. These are kernel-provided views, not ordinary persistent files. A system call requests a kernel operation; its result and failure indication must be checked according to the API used.

**Worked example.** Start a process that waits, note its PID, then inspect `/proc/<pid>/status`, `/proc/<pid>/cmdline`, and `/proc/<pid>/fd`. Match `Pid` and `PPid` with the observed process tree. Expect access restrictions and kernel/version differences in some fields.

## 2. Mechanism and vocabulary

**Mechanism.** A system call uses the architecture's defined transition mechanism to request privileged kernel work. The kernel validates arguments and permissions, performs or rejects the operation, and returns a result. A library wrapper may translate a low-level failure indicator into an `errno` value; Go APIs commonly return an `error` alongside ordinary results.

**Second point.** A regular function call stays within the process's user-space execution. A syscall crosses the user/kernel boundary, although library calls may avoid a kernel transition for some operations or use optimized mechanisms on particular systems; do not assume every function named in a library invokes a syscall.

In Go, handle both values returned by a fallible operation. Wrap errors with context using `%w` when callers may need to test the cause with `errors.Is` or `errors.As`. Do not ignore an error with `_` unless the failure is deliberately irrelevant and that decision is justified.

**Worked example.** Send a signal and record the state transition.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
package main

import (
    "fmt"
    "strconv"
)

func parseCount(text string) (int, error) {
    count, err := strconv.Atoi(text)
    if err != nil {
        return 0, fmt.Errorf("parse count %q: %w", text, err)
    }
    return count, nil
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

The parser returns a useful value only on success and preserves the underlying error for callers. Compare input `"12"` with `"twelve"`; the second must not silently become a zero count.

## 4. Verification, failure, and evidence

**Verification.** Record the PID, parent PID, command line, and selected descriptor targets for the sleeping process. For `parseCount`, show a valid input and malformed input and state the returned value/error for each.

**Failure mode.** Inspect a process that has already exited and explain why `/proc/<pid>` may disappear. Then call `parseCount` with nonnumeric text and verify the caller detects and reports the error rather than using a default zero accidentally.

**Completion criterion.** You can identify selected process metadata in `/proc`, explain the user/kernel boundary, and propagate a Go error with enough context to diagnose the failed operation.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
