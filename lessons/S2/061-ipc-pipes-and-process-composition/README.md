# Lesson 061 — IPC, Pipes, and Process Composition

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build a real producer/consumer pipeline using OS-level inter-process communication (pipes and/or local sockets), and handle backpressure and EOF correctly — the two edge cases that separate a toy IPC example from one that behaves correctly under real conditions.

**The contract.** **Why IPC exists at all.** Processes are isolated by design (Lesson 1's user/kernel boundary and per-process address spaces) — this isolation is a feature, not an obstacle, but it means processes that *need* to cooperate must use an explicit, OS-provided channel.

**Worked example.** Start with the smallest concrete case described by the source material: Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end.

## 2. Mechanism and vocabulary

**Mechanism.** Pipes, sockets, shared memory, and signals are different IPC mechanisms with different tradeoffs: pipes and sockets move data through the kernel with built-in flow control; shared memory is faster (no data copying through the kernel) but requires the processes to handle their own synchronization, since the OS provides no automatic ordering guarantees on shar.

**Second point.** **Pipes, reprised programmatically.** Linux tools Lesson 1 covered shell-level pipes conceptually.

**Worked example.** Ensure the child closes the write end when done, and the parent's read loop correctly detects EOF (a `read()` returning 0) and exits its loop rather than blocking forever.

## 3. Worked case from the reference material

Worked example: Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end.

## 4. Verification, failure, and evidence

**Verification.** For step 2, produce actual `strace` output (or timing measurements) showing at least one `write()` call from the producer taking noticeably longer than the others, correlating with the point where the pipe buffer would have filled given your artificial consumer delay — connect this specific observation back to the pipe buffer size mentioned in Learn.

**Failure mode.** Take step 4's hung consumer and, before killing it, inspect it with `strace -p <pid>` attached to the already-running, hung process (rather than starting a fresh `strace` session) — confirm you can see it sitting in a blocking `read()` syscall with no data arriving, distinguishing "hung waiting for more input that will never come because the pipe's write end.

**Completion criterion.** You've built a working producer/consumer pipeline using OS-level IPC (not just language-level channels) and correctly handle EOF, you've directly observed backpressure via `strace` on a real blocking `write()` call rather than just reading about it, and you've deliberately broken EOF handling and observed the resulting hang, then explained why `strace` alone.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
