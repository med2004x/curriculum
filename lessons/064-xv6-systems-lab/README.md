# Lesson 064 — Systems Lab: xv6 and Kernel Boundaries

## Threads
**Primary:** C/OS

**Supporting:** read/modify/test a teaching kernel

**Math companion:** proof by invariant

**Time:** 3–4 hours. **Prerequisites:** Lesson 063.

**Time:** 3–4 hours. **Prerequisites:** Lesson 063.

## What you will be able to do

1. Navigate the xv6-RISC-V source path for a user system call.
2. Add or trace a small system call across user stub, trap entry, dispatcher, kernel implementation, and return path.
3. State the kernel invariant affected by a change and preserve it across process state transitions.
4. Use the matching xv6 lab tests and a deliberate regression to localize a boundary failure.

## Retrieve first (no notes)

1. What does a process mapping show?
2. How is virtual address space different from resident memory?
3. What does a page fault signal?
4. How does Go's GC reachability differ from releasing physical pages to the OS?

## 1. Problem and model

**What problem this solves.** Bridge from OS theory (Lessons 1-7) to real kernel code by working directly in xv6, a small, teaching-oriented Unix-like kernel — reading actual kernel source before modifying it, and implementing at least one real system call or scheduling change end to end.

**The contract.** xv6 is a teaching kernel with a small codebase for studying real kernel mechanisms. A syscall crosses from user mode through trap-entry code to kernel dispatch and returns a result to user space. A change must preserve process, memory, and locking invariants as well as satisfy the user-level API contract.

**Worked example.** Trace one syscall end to end: user program calls a generated stub; the stub places the syscall number/arguments as required and executes the architecture trap instruction; trap code saves state; the dispatcher selects the kernel handler; the result is returned through the trap path.

Self-check: distinguish the user API, architecture trap mechanism, dispatch table, and kernel implementation; a missing link produces a different failure than a bug inside the handler.

## 2. Mechanism and vocabulary

**Mechanism.** Read the corresponding existing syscall, trap handler, and build/test wiring before editing. The syscall number, user stub, declarations, dispatcher entry, and handler must agree; the kernel handler should validate user-provided arguments before using them.

**Second point.** Kernel correctness is invariant-based: identify which state is protected by a lock, who owns a process/page-table resource, and when references become valid or are released. Keep lock ordering and cleanup consistent with neighboring code.

Use the exact xv6 branch and lab instructions for this repository. The upstream xv6 source and lab sequence can change; do not transplant line numbers or APIs from a different course version.

**Worked example.** Reference: `https://github.com/mit-pdos/xv6-riscv` and the course's published lab assignments.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: MIT's 6.1810 (formerly 6.S081) course provides the canonical xv6 labs, freely available; work through them using the actual course materials rather than reconstructing equivalent labs from scratch, since the lab infrastructure (test harness, provided starter code) is part of what makes this exercise tractable.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Run the test target supplied by the selected xv6 lab version and record the exact failing/passing test names. Also explain what the test observes at the user/kernel boundary; a green suite is evidence for covered behavior, not a proof of all kernel invariants.

**Failure mode.** Remove the syscall's dispatcher registration while leaving the user stub and handler intact. Trace why the request cannot reach the handler and distinguish dispatch failure from argument-validation or handler-logic failure.

**Completion criterion.** You can trace the syscall path, state the modified invariant, explain the kernel/user boundary, and use the matching lab tests to diagnose one deliberately introduced wiring regression.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the user-visible result when a syscall number is absent from the dispatcher.
2. Predict which part fails if a user stub exists but the handler is not linked.
3. Predict what the trap handler must preserve before dispatch.
4. Predict the result of passing an invalid user address to a syscall.
5. Predict what state must be protected by a lock in the selected lab.
6. Predict what each lab test observes and what it cannot establish.
7. Identify which code layer you would inspect first for a dispatch failure.
8. State the kernel invariant before making a change.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Select the repository's xv6-RISC-V version and one lab focused on syscalls, traps, memory, or scheduling.

2. Trace an existing implementation in source before changing it; write down the user API, dispatch path, handler, and invariant.

3. Implement the smallest specified change, then run the lab's provided tests and inspect the relevant failure path.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Pick one lab you completed and deliberately introduce a regression related to a specific concept from this track — for example, in the syscall lab, remove the syscall number's registration in the dispatch table while leaving the handler function itself intact, and observe the specific failure mode (the syscall not being found/dispatched, likely surfacing as.

2. Explain, using OS Lesson 1's syscall-dispatch discussion, exactly why this specific break produces this specific symptom rather than some other kind of failure — connecting the abstract "syscalls go through a dispatch table" idea to the literal line of xv6 source code responsible for it.

3. Deliberately violate your own documented ownership contract in a *separate* test function: if you documented `strbuf_cstr`'s pointer as borrowed (invalidated by the next `strbuf_append`), write a test that holds onto that pointer across an `append` call and then reads it anyway.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 064 — Systems Lab: xv6 and Kernel Boundaries: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the sequence from a user syscall stub to a kernel handler?
2. Which components must agree when adding a syscall?
3. What kernel invariant does the selected lab modify?
4. How does a dispatcher-registration failure differ from a handler bug?
5. What does the lab-provided test suite establish?
6. Why must the xv6 source version match the lab instructions?

## Evidence to record

Create `lesson064/EVIDENCE.md` with exactly this structure:

```text
# Lesson 064 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Editing a handler without wiring its number, declaration, and user stub.
- Passing unchecked user pointers into kernel code.
- Changing a protected process invariant without following its locking discipline.
- Using source instructions from a different xv6 revision.
- Treating test success as proof of untested kernel behavior.

## Precision checks

1. Trace user entry, trap, dispatcher, handler, and return.
2. State the changed kernel invariant and its protection.
3. Explain the regression symptom at the correct layer.
4. Record the matching lab test evidence and its scope.
5. Confirm the exact source/lab version used.

Kernel changes are boundary changes: trace the request and preserve the invariant across every layer it crosses.
## University revision bridge

For the LIG track, connect syscall/trap mechanisms to the OS model and use xv6 as a small implementation for inspecting those abstractions directly.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 064 — Systems Lab: xv6 and Kernel Boundaries without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've completed and passed the automated tests for at least the three labs specified above, you wrote a design note before implementing each one (not after, as documentation-after-the-fact), and you can trace, using your own completed syscall lab as the concrete example, the full path a system call takes from user-space code through the kernel and back, cit. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
