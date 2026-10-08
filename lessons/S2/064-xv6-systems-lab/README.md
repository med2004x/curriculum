# Lesson 064 — Systems Lab: xv6 and Kernel Boundaries

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Bridge from OS theory (Lessons 1-7) to real kernel code by working directly in xv6, a small, teaching-oriented Unix-like kernel — reading actual kernel source before modifying it, and implementing at least one real system call or scheduling change end to end.

**The contract.** xv6 is a teaching kernel with a small codebase for studying real kernel mechanisms. A syscall crosses from user mode through trap-entry code to kernel dispatch and returns a result to user space. A change must preserve process, memory, and locking invariants as well as satisfy the user-level API contract.

**Worked example.** Trace one syscall end to end: user program calls a generated stub; the stub places the syscall number/arguments as required and executes the architecture trap instruction; trap code saves state; the dispatcher selects the kernel handler; the result is returned through the trap path.

## 2. Mechanism and vocabulary

**Mechanism.** Read the corresponding existing syscall, trap handler, and build/test wiring before editing. The syscall number, user stub, declarations, dispatcher entry, and handler must agree; the kernel handler should validate user-provided arguments before using them.

**Second point.** Kernel correctness is invariant-based: identify which state is protected by a lock, who owns a process/page-table resource, and when references become valid or are released. Keep lock ordering and cleanup consistent with neighboring code.

Use the exact xv6 branch and lab instructions for this repository. The upstream xv6 source and lab sequence can change; do not transplant line numbers or APIs from a different course version.

**Worked example.** Reference: `https://github.com/mit-pdos/xv6-riscv` and the course's published lab assignments.

## 3. Worked case from the reference material

Worked example: MIT's 6.1810 (formerly 6.S081) course provides the canonical xv6 labs, freely available; work through them using the actual course materials rather than reconstructing equivalent labs from scratch, since the lab infrastructure (test harness, provided starter code) is part of what makes this exercise tractable.

## 4. Verification, failure, and evidence

**Verification.** Run the test target supplied by the selected xv6 lab version and record the exact failing/passing test names. Also explain what the test observes at the user/kernel boundary; a green suite is evidence for covered behavior, not a proof of all kernel invariants.

**Failure mode.** Remove the syscall's dispatcher registration while leaving the user stub and handler intact. Trace why the request cannot reach the handler and distinguish dispatch failure from argument-validation or handler-logic failure.

**Completion criterion.** You can trace the syscall path, state the modified invariant, explain the kernel/user boundary, and use the matching lab tests to diagnose one deliberately introduced wiring regression.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
