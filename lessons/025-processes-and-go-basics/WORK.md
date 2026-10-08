# Lesson 025 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 3. Architecture, Linux, Processes, and Concurrency  
**Primary:** Linux/Go  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What does `fork` do, and how does its return value differ between parent and child?
2. What does successful `exec` replace, and does it change the PID?
3. How does a parent retrieve a child's exit status?
4. How does `os/exec` relate to low-level `fork` and `exec`?
5. Which signals can and cannot be caught?
6. Why should graceful cleanup have a deadline?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict the return values of `fork` in the parent and child.
2. Predict whether a successful `exec` returns to the old program.
3. Predict which PID remains associated with the process after `exec`.
4. Predict the wait result when the child exits with status 7.
5. Predict what happens when the executable path is invalid.
6. Predict which file descriptors survive `exec` by default and what close-on-exec changes.
7. Predict whether a `SIGTERM` handler can run after `SIGKILL` is sent.
8. Record the process transition and status expected for each case.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Explain process creation, program replacement, waiting, and exit status.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Distinguish `fork` from the `exec` family and describe what state is inherited.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Write a small Go program with an explicit entry point, values, and error handling.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

You are given a small service component that must behave correctly under both normal and failure inputs. Design the state machine, then implement the smallest testable unit.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### Go process launcher

Write a Go command that starts a child process, captures its output and exit status, and handles a failed command cleanly.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Modify step 3's `SIGTERM` handler to deliberately take a long time (e.g.

2. `sleep(30)` inside the handler, simulating slow cleanup like draining a large in-flight request queue) before exiting.

3. Send `SIGTERM`, then — simulating what a container orchestrator does after its grace period expires — send `SIGKILL` a few seconds later, before the handler finishes.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint: use `os/exec` in Go instead of explicit C `fork`/`exec`, or change the child program and its arguments. Explain which low-level process details the Go API manages for you.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Algorithm statement:** Write the input, output, invariant, and termination argument before coding.
2. **Dry run:** Solve one non-trivial example by hand and show each state change.
3. **Complexity:** Give worst-case time and extra-space bounds and justify them.
4. **Transfer:** Change one constraint and explain which part of the algorithm must change.

## 11. Self-test

- [ ] What does `fork` do, and how does its return value differ between parent and child?
- [ ] What does successful `exec` replace, and does it change the PID?
- [ ] How does a parent retrieve a child's exit status?
- [ ] How does `os/exec` relate to low-level `fork` and `exec`?
- [ ] Which signals can and cannot be caught?
- [ ] Why should graceful cleanup have a deadline?

Score yourself only after the attempt: **0 = cannot answer, 1 = answer with notes, 2 = answer from memory, 3 = answer and justify a new variation.**

## 12. Evidence and grading

Create `EVIDENCE.md` in this lesson directory. Record:

- the exact commands or test commands used;
- one before/after prediction;
- one failure reproduced deliberately and its diagnosis;
- the result of the coding challenge;
- one transfer variation;
- your self-test score.

### Gate rubric

| Dimension | Pass condition |
|---|---|
| Knowledge | Explain the core model without notes. |
| Procedure | Produce a working implementation or derivation from blank. |
| Debugging | Reproduce, classify, and repair one realistic failure. |
| Transfer | Solve a changed constraint without being handed the algorithm. |
| Evidence | Another person can reproduce the verification from your notes. |

Suggested pass threshold: all five dimensions pass. If one fails, review that dimension before moving to the next lesson.

## 13. Mastery gate

Use the lecture's mastery gate in [`README.md`](README.md) as the primary criterion. This work packet adds the stronger requirement that you can satisfy the gate **after** completing independent work, not by recognising the lecture examples.

For integrated assessment lessons, the gate is summative: treat the challenge as an exam/project submission and do not consult previous solutions while working.

## 14. Spaced review

- **Next day:** answer the self-test from memory and redo one exercise without looking.
- **One week:** solve the transfer problem again under a changed constraint.
- **Later module review:** revisit this lesson when the concept reappears in a later module and compare the old and new evidence.

## 15. External practice reference

- [A Tour of Go](https://go.dev/tour/): work its exercises only after attempting the repository exercise so the external solution path does not replace your own reasoning.
- [Exercism Go Track](https://exercism.org/tracks/go): select a concept-matched exercise for additional independent practice and use feedback after your own attempt.
- [Go documentation](https://go.dev/doc/): use the primary documentation when the lesson depends on exact language or standard-library semantics.
