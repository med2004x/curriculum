# Lesson 064 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 7. Operating Systems and Systems Programming  
**Primary:** C/OS  
**Session target:** 3â€“6 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve â†’ guided practice â†’ independent exercises â†’ problem solving â†’ coding challenge â†’ review â†’ evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10â€“15 min retrieval, 25â€“45 min guided practice, 45â€“75 min exercises, 30â€“60 min problem solving, 60â€“120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What is the sequence from a user syscall stub to a kernel handler?
2. Which components must agree when adding a syscall?
3. What kernel invariant does the selected lab modify?
4. How does a dispatcher-registration failure differ from a handler bug?
5. What does the lab-provided test suite establish?
6. Why must the xv6 source version match the lab instructions?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A â€” Predict
1. Predict the user-visible result when a syscall number is absent from the dispatcher.
2. Predict which part fails if a user stub exists but the handler is not linked.
3. Predict what the trap handler must preserve before dispatch.
4. Predict the result of passing an invalid user address to a syscall.
5. Predict what state must be protected by a lock in the selected lab.
6. Predict what each lab test observes and what it cannot establish.
7. Identify which code layer you would inspect first for a dispatch failure.
8. State the kernel invariant before making a change.

Run the smallest case first. Record prediction and observation side by side.

### Practice B â€” Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C â€” Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer â€œwhat changed?â€ rather than merely â€œdid it work?â€

## 4. Independent exercises

### Exercise 1 â€” Recall and prediction
Without running anything, explain the main mechanism in **Navigate the xv6-RISC-V source path for a user system call.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 â€” Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Add or trace a small system call across user stub, trap entry, dispatcher, kernel implementation, and return path.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 â€” Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 â€” Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 â€” Explanation
Write a 150â€“250 word technical explanation of **State the kernel invariant affected by a change and preserve it across process state transitions.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

You are given a small specification with two valid states and one invalid state. Decompose it into inputs, invariant, transition, and verification before writing code.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### MIT xv6 lab assignment

Complete the MIT 6.1810 system-call lab in its official xv6-labs checkout. The local `exercise.c` is only a source-tracing prompt.

### Acceptance criteria

- Clone the course lab repo named in the handout, check out its `syscall` branch, and keep the upstream tests unchanged.
- Trace the user/kernel boundary touched by your change and state the invariant it must preserve.
- Run the provided test target and one relevant regression or failure case.
- Record the upstream commit, lab name, exact commands, and test results in `EVIDENCE.md`.
- Explain one failing test from its assertion back to the kernel state transition.


## 8. Debugging and failure lab

1. Pick one lab you completed and deliberately introduce a regression related to a specific concept from this track â€” for example, in the syscall lab, remove the syscall number's registration in the dispatch table while leaving the handler function itself intact, and observe the specific failure mode (the syscall not being found/dispatched, likely surfacing as.

2. Explain, using OS Lesson 1's syscall-dispatch discussion, exactly why this specific break produces this specific symptom rather than some other kind of failure â€” connecting the abstract "syscalls go through a dispatch table" idea to the literal line of xv6 source code responsible for it.

3. Deliberately violate your own documented ownership contract in a *separate* test function: if you documented `strbuf_cstr`'s pointer as borrowed (invalidated by the next `strbuf_append`), write a test that holds onto that pointer across an `append` call and then reads it anyway.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint in the main case for # 064 â€” Systems Lab: xv6 and Kernel Boundaries: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Formal model:** State the relevant relation, invariant, transaction property, or storage rule.
2. **Schedule/table:** Work one small example manually.
3. **Counterexample:** Construct one invalid or surprising case.
4. **Verification:** State the SQL command, trace, or test that would falsify your answer.

## 11. Self-test

- [ ] What is the sequence from a user syscall stub to a kernel handler?
- [ ] Which components must agree when adding a syscall?
- [ ] What kernel invariant does the selected lab modify?
- [ ] How does a dispatcher-registration failure differ from a handler bug?
- [ ] What does the lab-provided test suite establish?
- [ ] Why must the xv6 source version match the lab instructions?

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

- Use the lesson's assigned material and primary technical documentation when a precise semantic detail needs verification. The repository challenge remains the assessment artifact.
