# Lesson 061 — IPC, Pipes, and Process Composition

## Threads
**Primary:** OS/Linux/Go

**Supporting:** pipes; child processes; composition

**Math companion:** graphs and state

**Time:** 3–4 hours. **Prerequisites:** Lesson 060.

**Time:** 3–4 hours. **Prerequisites:** Lesson 060, # Lesson 060 — Filesystems, Storage, and Durability.

## What you will be able to do

1. Explain the central model of # 061 — IPC, Pipes, and Process Composition and name the boundary where it applies.
2. Trace one concrete case of # 061 — IPC, Pipes, and Process Composition before executing it.
3. Rebuild the core exercise from an empty workspace without copying the reference implementation.
4. Produce terminal evidence for one success case and one failure case.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 060, # Lesson 060 — Filesystems, Storage, and Durability.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Build a real producer/consumer pipeline using OS-level inter-process communication (pipes and/or local sockets), and handle backpressure and EOF correctly — the two edge cases that separate a toy IPC example from one that behaves correctly under real conditions.

**The contract.** **Why IPC exists at all.** Processes are isolated by design (Lesson 1's user/kernel boundary and per-process address spaces) — this isolation is a feature, not an obstacle, but it means processes that *need* to cooperate must use an explicit, OS-provided channel.

**Worked example.** Start with the smallest concrete case described by the source material: Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Pipes, sockets, shared memory, and signals are different IPC mechanisms with different tradeoffs: pipes and sockets move data through the kernel with built-in flow control; shared memory is faster (no data copying through the kernel) but requires the processes to handle their own synchronization, since the OS provides no automatic ordering guarantees on shar.

**Second point.** **Pipes, reprised programmatically.** Linux tools Lesson 1 covered shell-level pipes conceptually.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Ensure the child closes the write end when done, and the parent's read loop correctly detects EOF (a `read()` returning 0) and exits its loop rather than blocking forever.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** For step 2, produce actual `strace` output (or timing measurements) showing at least one `write()` call from the producer taking noticeably longer than the others, correlating with the point where the pipe buffer would have filled given your artificial consumer delay — connect this specific observation back to the pipe buffer size mentioned in Learn.

**Failure mode.** Take step 4's hung consumer and, before killing it, inspect it with `strace -p <pid>` attached to the already-running, hung process (rather than starting a fresh `strace` session) — confirm you can see it sitting in a blocking `read()` syscall with no data arriving, distinguishing "hung waiting for more input that will never come because the pipe's write end.

**Completion criterion.** You've built a working producer/consumer pipeline using OS-level IPC (not just language-level channels) and correctly handle EOF, you've directly observed backpressure via `strace` on a real blocking `write()` call rather than just reading about it, and you've deliberately broken EOF handling and observed the resulting hang, then explained why `strace` alone.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 061 — IPC, Pipes, and Process Composition case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end.

2. Extend it with one explicit boundary case: Ensure the child closes the write end when done, and the parent's read loop correctly detects EOF (a `read()` returning 0) and exits its loop rather than blocking forever.

3. Rebuild the same behavior with one changed constraint: Modify step 1 so the parent (consumer) deliberately reads slowly (add a small `sleep()` between reads) while the child (producer) writes a large volume of data quickly with no delay.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take step 4's hung consumer and, before killing it, inspect it with `strace -p <pid>` attached to the already-running, hung process (rather than starting a fresh `strace` session) — confirm you can see it sitting in a blocking `read()` syscall with no data arriving, distinguishing "hung waiting for more input that will never come because the pipe's write end.

2. Create an unclosed file descriptor and inspect the process's open descriptors.

3. Explain the resource leak.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 061 — IPC, Pipes, and Process Composition: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 061 — IPC, Pipes, and Process Composition?
2. What two concrete practices are required in the assigned material for # 061 — IPC, Pipes, and Process Composition?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson061/EVIDENCE.md` with exactly this structure:

```text
# Lesson 061 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 061 — IPC, Pipes, and Process Composition.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 061 — IPC, Pipes, and Process Composition without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've built a working producer/consumer pipeline using OS-level IPC (not just language-level channels) and correctly handle EOF, you've directly observed backpressure via `strace` on a real blocking `write()` call rather than just reading about it, and you've deliberately broken EOF handling and observed the resulting hang, then explained why `strace` alone. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
