# Lesson 058 — Processes, Threads, and Scheduling Revisited

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand the process abstraction as the OS's mechanism for isolating and multiplexing programs, and the user/kernel boundary that every system call crosses.

**The contract.** A process has a protected virtual address space and operating-system-managed resources such as open descriptors. The scheduler selects runnable execution contexts for available CPU time; blocked tasks are not currently eligible to run until the event they await occurs. A process is not guaranteed to run continuously or to own a physical CPU.

**Worked example.** Trace a small program that opens, reads, and prints a file. Identify `openat`, `read`, and `write` calls that appear, while noting libc buffering, startup, and runtime behavior can add or combine syscalls beyond the source-level call sequence.

## 2. Mechanism and vocabulary

**Mechanism.** System calls transition from user execution to a kernel service under the architecture's ABI. The kernel validates the request and may block, return a result, or report an error. Scheduling can preempt a runnable process/thread so another can use a CPU.

**Second point.** A terminated child whose status has not been collected is a zombie: most of its execution resources have been released, but the kernel retains a small process-table record and exit status until the parent or a reaper collects it with `wait`/`waitpid`.

Do not describe a zombie as still running or as retaining its full address space. A large number of unreaped children can consume process identifiers and kernel process-table resources, but the child itself is not executing.

**Worked example.** Compare the program's source-level file operations with `strace` output. Explain why a `printf` may buffer output and why process startup may issue calls not visible in `main`.

## 3. Worked case from the reference material

Worked example: Use `strace -c` (from Linux tools Lesson 5) on a simple program (e.g. one that reads a file and prints its contents) and identify which specific syscalls correspond to which C library calls you wrote (`fopen`→likely `openat`, `fread`→`read`, `printf`→eventually `write`).

## 4. Verification, failure, and evidence

**Verification.** Produce a table of observed syscalls and plausible source/runtime causes. Then create a child process that exits while its parent delays collection; identify its terminated/zombie state and show how `waitpid` removes the entry.

**Failure mode.** Leave one terminated child unreaped temporarily, then reap it. Explain which process-table record persists and which execution resources have already been released.

**Completion criterion.** You can interpret representative syscalls, describe process scheduling states, and explain the zombie/reaping lifecycle without conflating it with a running child.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
