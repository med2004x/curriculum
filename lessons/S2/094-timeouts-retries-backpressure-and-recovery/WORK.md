# Lesson 094 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 10. Reliability, Performance, Production, and Graduation  
**Primary:** System Design/Go/Distributed  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. Why propagate one deadline through a call chain?
2. Which operations are safe to retry?
3. What do exponential backoff and jitter each accomplish?
4. How does a retry budget limit amplification?
5. What is backpressure and how can a bounded queue enforce it?
6. Why can a timeout be too short or too long?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict how a child deadline relates to its caller's deadline.
2. Predict the maximum attempt count under a retry budget.
3. Predict how jitter changes synchronized retry traffic.
4. Predict whether retrying a non-idempotent operation is safe without a key.
5. Predict queue behavior when capacity is reached.
6. Predict what happens to latency when a timeout is too short.
7. Predict the recovery risk if all clients retry immediately after an outage.
8. State which overload metric should trigger admission control or shedding.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Propagate an end-to-end deadline through a multi-service call chain.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Retry only suitable failures with bounded exponential backoff, jitter, and an overall retry budget.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Apply backpressure through bounded queues, admission control, or load shedding.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### Resilient client

Implement a client with deadlines, bounded retries, exponential backoff, and backpressure. Test slow, flaky, and permanently failing dependencies.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Configure your step 3 timeout to be deliberately too short (much shorter than C's *normal*, healthy response time, not just its slowed-down time) and rerun the simulation with C responding at its normal, healthy speed.

2. Confirm B now frequently times out on requests that would have succeeded, generating unnecessary failures and retries even though nothing is actually wrong with C — directly demonstrating the "too short" failure mode from Learn, using your own measured false-timeout rate as evidence, and connect this concretely to why timeout values need to be chosen based o.

3. Take your step 2 bounded-queue-with-rejection service and deliberately configure the queue capacity far too small (e.g. capacity of 1, when your actual expected burst size is much larger) relative to normal, healthy traffic patterns (not just overload conditions).

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one dependency or retry policy: alter the overall deadline, queue capacity, or maximum attempts. Explain how resource use, success probability, and overload response change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Model:** State the system model and assumptions before solving.
2. **Trace:** Work a small failure or state-transition example by hand.
3. **Trade-off:** Compare two designs and state the condition under which each wins.
4. **Evidence:** Identify the measurement, test, or trace that would distinguish the designs in practice.

## 11. Self-test

- [ ] Why propagate one deadline through a call chain?
- [ ] Which operations are safe to retry?
- [ ] What do exponential backoff and jitter each accomplish?
- [ ] How does a retry budget limit amplification?
- [ ] What is backpressure and how can a bounded queue enforce it?
- [ ] Why can a timeout be too short or too long?

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
