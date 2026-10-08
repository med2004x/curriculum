# Lesson 075 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 8. Backend Engineering  
**Primary:** Go/Production  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What behavior belongs in a unit test versus an integration test?
2. Why is container running distinct from dependency readiness?
3. Which container state is ephemeral by default?
4. How do volumes change persistence?
5. How are test fixtures isolated and cleaned up?
6. Which runtime configuration should not be baked into the image?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict whether a container's writable layer persists after replacement.
2. Predict what happens if integration tests begin before database readiness.
3. Predict which test scope can verify SQL constraints.
4. Predict how two parallel tests can interfere through shared fixture IDs.
5. Predict what state a mounted volume preserves.
6. Predict which environment variables or secrets must be supplied at runtime.
7. Predict how a migration failure should affect test startup.
8. State the cleanup mechanism before creating test data.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Choose unit, integration, or end-to-end scope for a stated behavior.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Start a disposable PostgreSQL dependency and run integration tests against an isolated database.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Build a container image with explicit runtime configuration and health/readiness expectations.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### Container integration

Containerize a Go/PostgreSQL service and create integration tests that start dependencies, run migrations, test the service, and clean up deterministically.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Remove the database health/readiness gate and observe startup/test behavior when the dependency is unavailable.

2. Reuse fixture IDs across parallel tests and identify the resulting cross-test interference.

3. Restore explicit readiness and fixture isolation; explain why containerization alone does not make tests deterministic.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one test environment condition: use a fresh database per test suite or one shared database with isolated schemas. Explain the isolation and cleanup trade-off.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Contract recall:** State the protocol/API/security rule without notes.
2. **Scenario analysis:** Given one malformed or adversarial request, determine the expected system response.
3. **Why:** Explain the invariant or security property that makes that response correct.
4. **Implementation bridge:** Name the test or runtime observation that would prove the rule is enforced.

## 11. Self-test

- [ ] What behavior belongs in a unit test versus an integration test?
- [ ] Why is container running distinct from dependency readiness?
- [ ] Which container state is ephemeral by default?
- [ ] How do volumes change persistence?
- [ ] How are test fixtures isolated and cleaned up?
- [ ] Which runtime configuration should not be baked into the image?

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
