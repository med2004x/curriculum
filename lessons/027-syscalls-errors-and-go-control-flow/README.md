# Lesson 027 — Syscalls, Errors, and Explicit Control Flow

## Threads
**Primary:** Linux/Go

**Supporting:** user/kernel boundary; Go errors; return values

**Math companion:** logic and predicates

**Time:** 3–4 hours. **Prerequisites:** Lesson 026.

**Time:** 3–4 hours. **Prerequisites:** Lesson 026.

## What you will be able to do

1. Explain how a system call crosses from user execution into a kernel service.
2. Inspect selected Linux process information under `/proc/<pid>` and distinguish it from ordinary files.
3. Explain why process state, system-call return values, and Go `error` values require explicit checks.
4. Trace a successful and failed operation without discarding the error information.

## Retrieve first (no notes)

1. What is a file descriptor?
2. Which conventional descriptors are stdin, stdout, and stderr?
3. What does a pipe connect?
4. How can inherited write descriptors prevent a reader from observing EOF?

## 1. Problem and model

**What problem this solves.** Observe process state from Linux instead of inferring it from program output alone.

**The contract.** Linux exposes selected process information through pseudo-files under `/proc/<pid>`, including `status`, `cmdline`, and descriptor links under `fd/`. These are kernel-provided views, not ordinary persistent files. A system call requests a kernel operation; its result and failure indication must be checked according to the API used.

**Worked example.** Start a process that waits, note its PID, then inspect `/proc/<pid>/status`, `/proc/<pid>/cmdline`, and `/proc/<pid>/fd`. Match `Pid` and `PPid` with the observed process tree. Expect access restrictions and kernel/version differences in some fields.

Self-check: distinguish an observed `/proc` field from an inference about the process's internal state.

## 2. Mechanism and vocabulary

**Mechanism.** A system call uses the architecture's defined transition mechanism to request privileged kernel work. The kernel validates arguments and permissions, performs or rejects the operation, and returns a result. A library wrapper may translate a low-level failure indicator into an `errno` value; Go APIs commonly return an `error` alongside ordinary results.

**Second point.** A regular function call stays within the process's user-space execution. A syscall crosses the user/kernel boundary, although library calls may avoid a kernel transition for some operations or use optimized mechanisms on particular systems; do not assume every function named in a library invokes a syscall.

In Go, handle both values returned by a fallible operation. Wrap errors with context using `%w` when callers may need to test the cause with `errors.Is` or `errors.As`. Do not ignore an error with `_` unless the failure is deliberately irrelevant and that decision is justified.

**Worked example.** Send a signal and record the state transition.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: why is returning `(0, err)` safer than returning the partially parsed integer together with an error?

## 4. Verification, failure, and evidence

**Verification.** Record the PID, parent PID, command line, and selected descriptor targets for the sleeping process. For `parseCount`, show a valid input and malformed input and state the returned value/error for each.

**Failure mode.** Inspect a process that has already exited and explain why `/proc/<pid>` may disappear. Then call `parseCount` with nonnumeric text and verify the caller detects and reports the error rather than using a default zero accidentally.

**Completion criterion.** You can identify selected process metadata in `/proc`, explain the user/kernel boundary, and propagate a Go error with enough context to diagnose the failed operation.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the relationship between a process's PID and its parent PID.
2. Predict which `/proc` entries are likely to vanish after the process exits.
3. Predict whether a process's `cmdline` contains ordinary space-separated display text or argument bytes separated by NULs.
4. Predict the result and error from parsing `"12"` and `"twelve"`.
5. Predict what wrapping an error with `%w` preserves for callers.
6. Predict which failures should be handled by the kernel-facing operation and which by application policy.
7. Distinguish a syscall failure from a Go function returning a valid zero value.
8. Record the expected values before inspecting the process or calling the parser.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Run a sleeping process and inspect its PID, parent PID, open descriptors, and command line.

2. Extend it with one explicit boundary case: Send a signal and record the state transition.

3. Rebuild the same behavior with one changed constraint: Trace a program that opens and reads a file.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one condition: inspect a short-lived process rather than a sleeping one, or parse an integer with leading/trailing whitespace. Explain whether the observed behavior follows from `/proc` lifecycle or the chosen parsing policy.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does `/proc/<pid>/status` expose?
2. Why can a `/proc/<pid>` entry disappear?
3. How does a syscall differ from an ordinary user-space function call?
4. What does Go return for a failed operation in the usual error convention?
5. Why wrap an error with `%w`?
6. Why is ignoring the error from `strconv.Atoi` unsafe?

## Evidence to record

Create `lesson027/EVIDENCE.md` with exactly this structure:

```text
# Lesson 027 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating `/proc` entries as persistent disk files.
- Assuming process metadata remains available after the process exits.
- Ignoring a Go error and using a zero value as if parsing succeeded.
- Claiming every library function call is a syscall.
- Conflating a syscall return convention with Go's `(value, error)` API convention.

## Precision checks

1. Identify at least three process fields under `/proc/<pid>` and explain their meaning.
2. Explain why a syscall crosses a protection boundary.
3. Show a valid parse and a parse failure with explicit error handling.
4. Explain how wrapped errors remain inspectable.
5. Distinguish process metadata observation from inference about internal kernel state.

Separate observation from interpretation: `/proc` exposes selected kernel-maintained process information, not every detail of execution.
## University revision bridge

For the LIG track, connect system-call transitions to privilege separation, and connect Go's explicit error return to program control flow and API contracts.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can inspect and explain selected `/proc` process fields, describe what makes a syscall a user/kernel boundary, and propagate a Go parsing error without silently using an invalid result.
