# Lesson 026 — File Descriptors and Go I/O

## Threads
**Primary:** Linux/Go

**Supporting:** FDs; stdin/stdout; readers/writers

**Math companion:** functions and composition

**Time:** 3–4 hours. **Prerequisites:** Lesson 025.

**Time:** 3–4 hours. **Prerequisites:** Lesson 025.

## What you will be able to do

1. Explain that a file descriptor is a process-local integer handle for an open I/O resource.
2. Trace bytes through stdin, stdout, stderr, a file, and a pipe.
3. Explain how shell redirection and pipes change descriptor connections.
4. Relate Go's `io.Reader`, `io.Writer`, and `os.File` to concrete I/O resources.

## Retrieve first (no notes)

1. What does `fork` create?
2. What does successful `exec` replace?
3. How does a parent collect a child's exit status?
4. Why can a child inherit open descriptors across `fork`?

## 1. Problem and model

**What problem this solves.** Understand why Linux treats files, terminals, pipes, and sockets as file-descriptor-backed I/O resources.

**The contract.** On Unix-like systems, stdin, stdout, and stderr conventionally use descriptors 0, 1, and 2. `open` returns a descriptor or failure; `read` and `write` operate on it; `close` releases it. A pipe connects one process's output to another process's input. Shell redirection changes these connections before the command runs.

**Worked example.** Start with the smallest concrete case described by the source material: Build a small C program that writes to stdout and stderr.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** A descriptor is a process-local integer handle. The kernel associates it with an open file description or another I/O resource. Duplicated descriptors can refer to the same open description and therefore share state such as a file offset.

**Second point.** Go's `os.File` provides methods that implement interfaces such as `io.Reader` and `io.Writer`. Those interfaces describe operations; a writer need not be a disk file.

`command >out.txt 2>err.txt` routes stdout and stderr separately. `producer | consumer` connects producer stdout to consumer stdin; stderr remains separate unless redirected too.

**Worked example.** Redirect each stream separately.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Use a tiny Go program to compare standard output and standard error:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprintln(os.Stdout, "normal output")
    fmt.Fprintln(os.Stderr, "diagnostic output")
}
```

`os.Stdout` and `os.Stderr` are `*os.File` handles for the conventional standard streams. Redirect each stream separately, then contrast this Go abstraction with the assigned C calls to `open`, `read`, `write`, and `close`.

Worked example check: identify the process descriptor receiving each write and the shell redirection that changes its destination.

Self-check: if stdout is redirected to a file, where does stderr go when no stderr redirection is supplied?

## 4. Verification, failure, and evidence

**Verification.** Record stdout and stderr separately for the two-stream program. Trace one pipeline's descriptor wiring and inspect what happens when the consumer exits early. Record return values and errors from the I/O calls.

**Failure mode.** Create an unclosed file descriptor and inspect the process's open descriptors.

**Completion criterion.** You can trace an I/O operation through the process descriptor table to its resource and explain how a pipe or redirection changes that path.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict which descriptor `os.Stdout` writes to.
2. Predict the destinations under `>out.txt 2>err.txt`.
3. Predict which stream is connected by `producer | consumer`.
4. Predict where a diagnostic goes if only stdout is redirected.
5. Predict the result of writing to a closed descriptor.
6. Predict why a pipe read returns EOF only after all write ends are closed.
7. Predict the resource effect of leaving a descriptor open.
8. Draw the descriptor flow before checking it in a shell.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Build a small C program that writes to stdout and stderr.

2. Extend it with one explicit boundary case: Redirect each stream separately.

3. Rebuild the same behavior with one changed constraint: Then connect two processes with a pipe and describe the byte flow.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Create an unclosed file descriptor and inspect the process's open descriptors.

2. Explain the resource leak.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one connection: redirect stdout, redirect stderr separately, or connect stdout to a pipe. Explain how the descriptor wiring changes without changing the program's write calls.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is a file descriptor?
2. What are the conventional descriptors for stdin, stdout, and stderr?
3. What does a pipe connect?
4. How does redirection change a process's I/O destinations?
5. What does Go's `io.Writer` specify?
6. Why must unused descriptors be closed?

## Evidence to record

Create `lesson026/EVIDENCE.md` with exactly this structure:

```text
# Lesson 026 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Assuming stdout and stderr are the same stream because both display in a terminal by default.
- Treating a descriptor as a filename rather than a process-local handle.
- Forgetting that inherited descriptors can keep a pipe open and prevent EOF.
- Leaving descriptors open after their last use.
- Assuming an `io.Writer` must be a regular disk file.

## Precision checks

1. Draw stdout/stderr routing for a redirected process.
2. Draw the descriptor connection for a two-process pipe.
3. Explain how descriptor inheritance affects a pipe's lifetime.
4. Relate a Go `io.Writer` call to its concrete destination.
5. State when and why a descriptor is closed.

I/O behavior is easiest to reason about as a graph: processes are endpoints, descriptors are handles, and kernel resources or pipes are connections.
## University revision bridge

For the LIG track, connect file-descriptor operations to the syscall boundary, then relate Go's reader/writer interfaces to abstraction and composability.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain descriptor ownership, trace separate stdout/stderr redirection and a pipe, identify an unclosed-descriptor failure, and explain how Go's `io.Reader`/`io.Writer` abstractions compose over concrete I/O resources.
