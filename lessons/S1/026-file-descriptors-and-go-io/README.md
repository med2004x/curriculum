# Lesson 026 — File Descriptors and Go I/O

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand why Linux treats files, terminals, pipes, and sockets as file-descriptor-backed I/O resources.

**The contract.** On Unix-like systems, stdin, stdout, and stderr conventionally use descriptors 0, 1, and 2. `open` returns a descriptor or failure; `read` and `write` operate on it; `close` releases it. A pipe connects one process's output to another process's input. Shell redirection changes these connections before the command runs.

**Worked example.** Start with the smallest concrete case described by the source material: Build a small C program that writes to stdout and stderr.

## 2. Mechanism and vocabulary

**Mechanism.** A descriptor is a process-local integer handle. The kernel associates it with an open file description or another I/O resource. Duplicated descriptors can refer to the same open description and therefore share state such as a file offset.

**Second point.** Go's `os.File` provides methods that implement interfaces such as `io.Reader` and `io.Writer`. Those interfaces describe operations; a writer need not be a disk file.

`command >out.txt 2>err.txt` routes stdout and stderr separately. `producer | consumer` connects producer stdout to consumer stdin; stderr remains separate unless redirected too.

**Worked example.** Redirect each stream separately.

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

## 4. Verification, failure, and evidence

**Verification.** Record stdout and stderr separately for the two-stream program. Trace one pipeline's descriptor wiring and inspect what happens when the consumer exits early. Record return values and errors from the I/O calls.

**Failure mode.** Create an unclosed file descriptor and inspect the process's open descriptors.

**Completion criterion.** You can trace an I/O operation through the process descriptor table to its resource and explain how a pipe or redirection changes that path.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
