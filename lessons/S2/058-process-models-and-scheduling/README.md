# Lesson 058 — Processes, Threads, and Scheduling Revisited

## Threads
**Primary:** OS/Linux

**Supporting:** process states; scheduling; runnable work

**Math companion:** queue intuition

**Time:** 3–4 hours. **Prerequisites:** Lesson 057.

**Time:** 3–4 hours. **Prerequisites:** Lesson 057.

## What you will be able to do

1. Explain which resources are associated with a process and how processes are isolated.
2. Distinguish process states such as running/runnable, blocked, and terminated.
3. Relate library I/O calls to the system calls visible in a trace without assuming a one-to-one mapping.
4. Explain how a parent reaps child exit status and why an unreaped terminated child remains a zombie.

## Retrieve first (no notes)

1. What is the request path through handler, service, repository, and database?
2. Why must a mutation not be reported successful before commit?
3. How should a missing row differ from a database failure?
4. Why should integration-test fixtures be isolated?

## 1. Problem and model

**What problem this solves.** Understand the process abstraction as the OS's mechanism for isolating and multiplexing programs, and the user/kernel boundary that every system call crosses.

**The contract.** A process has a protected virtual address space and operating-system-managed resources such as open descriptors. The scheduler selects runnable execution contexts for available CPU time; blocked tasks are not currently eligible to run until the event they await occurs. A process is not guaranteed to run continuously or to own a physical CPU.

**Worked example.** Trace a small program that opens, reads, and prints a file. Identify `openat`, `read`, and `write` calls that appear, while noting libc buffering, startup, and runtime behavior can add or combine syscalls beyond the source-level call sequence.

Self-check: a system call trace shows kernel-boundary operations, not every library call or every user-space function.

## 2. Mechanism and vocabulary

**Mechanism.** System calls transition from user execution to a kernel service under the architecture's ABI. The kernel validates the request and may block, return a result, or report an error. Scheduling can preempt a runnable process/thread so another can use a CPU.

**Second point.** A terminated child whose status has not been collected is a zombie: most of its execution resources have been released, but the kernel retains a small process-table record and exit status until the parent or a reaper collects it with `wait`/`waitpid`.

Do not describe a zombie as still running or as retaining its full address space. A large number of unreaped children can consume process identifiers and kernel process-table resources, but the child itself is not executing.

**Worked example.** Compare the program's source-level file operations with `strace` output. Explain why a `printf` may buffer output and why process startup may issue calls not visible in `main`.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Use `strace -c` (from Linux tools Lesson 5) on a simple program (e.g. one that reads a file and prints its contents) and identify which specific syscalls correspond to which C library calls you wrote (`fopen`→likely `openat`, `fread`→`read`, `printf`→eventually `write`).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Produce a table of observed syscalls and plausible source/runtime causes. Then create a child process that exits while its parent delays collection; identify its terminated/zombie state and show how `waitpid` removes the entry.

**Failure mode.** Leave one terminated child unreaped temporarily, then reap it. Explain which process-table record persists and which execution resources have already been released.

**Completion criterion.** You can interpret representative syscalls, describe process scheduling states, and explain the zombie/reaping lifecycle without conflating it with a running child.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether a blocked process is runnable for the CPU.
2. Predict what `waitpid` returns for an exited child.
3. Predict which record remains while a child is a zombie.
4. Predict what changes after the parent reaps that child.
5. Predict why one `fread` or `printf` call may correspond to zero, one, or several observed syscalls.
6. Predict how scheduler preemption can allow several processes to make progress on one core.
7. Distinguish user-space function calls from system-call transitions.
8. State what the trace directly shows versus what you infer from source code.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Use `strace -c` (from Linux tools Lesson 5) on a simple program (e.g. one that reads a file and prints its contents) and identify which specific syscalls correspond to which C library calls you wrote (`fopen`→likely `openat`, `fread`→`read`, `printf`→eventually `write`).

2. Extend it with one explicit boundary case: Note any syscalls you didn't expect (e.g.

3. Rebuild the same behavior with one changed constraint: `mmap` calls from the C library's own startup or buffering).

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Extend step 4's zombie-creation scenario: fork *many* children in a loop (e.g.

2. 50) from a parent that never calls `wait()` on any of them, and observe via `ps aux | grep defunct` (or counting zombie states in `/proc`) that all 50 remain as zombies simultaneously.

3. Explain why this is a real, not just theoretical, resource leak — each zombie entry consumes a slot in the kernel's process table (a finite resource), and a long-running server process that forks children without ever reaping them will eventually exhaust available process table entries, a real historical class of production bug in daemon processes with buggy.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one condition: add explicit output flushing, or reap the child immediately instead of after a delay. Explain which observable syscalls or process states may change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is a process's virtual address space?
2. What makes an execution context runnable versus blocked?
3. What does a system call trace show?
4. Why is the mapping from library calls to syscalls not one-to-one?
5. What is a zombie and how is it removed?
6. Which child resources remain after termination?

## Evidence to record

Create `lesson058/EVIDENCE.md` with exactly this structure:

```text
# Lesson 058 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a system-call trace as a direct trace of all source-level function calls.
- Confusing blocked with runnable state.
- Describing a zombie as still executing.
- Assuming a zombie retains its full memory image.
- Ignoring output buffering when interpreting write syscalls.

## Precision checks

1. Identify the process state and scheduler eligibility.
2. Explain the user/kernel boundary of a syscall.
3. Correlate a syscall trace with source and runtime behavior.
4. Explain zombie retention and the effect of `waitpid`.
5. State which observations are direct and which are interpretations.

Process state describes eligibility and lifecycle; a syscall trace describes kernel requests. Use each as evidence for a different layer.
## University revision bridge

For the LIG track, connect process state transitions and scheduling to the OS model, then interpret syscall traces with libc and runtime behavior in mind.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 058 — Processes, Threads, and Scheduling Revisited without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can map a program's library calls to the actual syscalls they invoke using `strace`, you've directly observed a zombie process existing in the process table and explained why it persists until reaped, and you can articulate why the user/kernel boundary — not just convention — is what actually prevents user code from directly touching hardware or other pr. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
