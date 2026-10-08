# Lesson 011 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 1. Machine and Programming Foundations  
**Primary:** C  
**Supporting:** test cases; assertions; code reading; evidence  
**Math companion:** logic review  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. State the valid score range and the exact passing rule.
2. Why are 49 and 50 a more useful pair than 20 and 80 for checking the threshold?
3. What is the expected result for 100? For 101?
4. What does the check-case function compare?
5. What process status means that every test passed in this program?
6. How does the score-50 test expose the `>` mutation?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict every expected classification in the table before running the program.
2. Predict the actual result at 50 after changing `>=` to `>`.
3. Predict the process exit status when one check fails.
4. Predict whether `101` reaches the threshold comparison or returns at the range check.
5. Predict what happens if the test for 49 is deleted while the implementation wrongly uses `>`.
6. Predict which evidence distinguishes a bad classification from a failed compilation.
7. Run the original and mutated versions; record prediction beside observation.

For item 6, the compiled test's `FAIL score=...` line means the program ran and compared values. A compiler diagnostic means no such test result has yet been produced.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Turn an input/output contract into normal, boundary, and invalid test cases.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Trace a C function on paper before compiling it.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Read a failing test as evidence about a specific expected and actual value.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

Model the resource, owner, state transitions, and failure path explicitly. Solve a small concrete case before implementing the general mechanism.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### Verification harness

Write a minimal test harness for a small C library. Include normal cases, boundary cases, and one test that exposes a seeded defect.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Change `>=` to `>` and confirm that the score-50 case catches it.
2. Change the expected result for 49 from `FAIL` to `PASS`; confirm that the test harness reports a mismatch even though the implementation is unchanged.
3. Remove the lower-bound check; confirm that -1 is no longer classified as `INVALID`.
4. For each failure, save the diagnostic line and exit status before restoring the source.

Classify each observation: compiler diagnostic, test assertion/mismatch, or process status. They answer different questions. A clean compile proves neither that the contract is correct nor that every test passes.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change the contract so a valid score passes at 60 instead of 50. Keep the valid range at 0 through 100. From a blank copy, choose cases 59, 60, and 61 plus both range endpoints and both invalid neighbors. Write expected outcomes before editing the code.

Do not merely change the macro and reuse the original tests: add the new boundary pair so the suite would reject both `score >= 59` and `score > 60`. Explain which case catches each defect, then run the suite and record its exit status.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Closed-book definition:** State the key formal definition from this lesson in one or two sentences.
2. **Worked derivation:** Draw the state transition or representation for the concrete example used in the lesson.
3. **Exam variation:** Change one parameter and recompute the result without running code.
4. **Engineering bridge:** Identify the exact observation in the terminal that would confirm your written result.

LIG support: [BUSINESS-COMMUNICATION](../S1/s1_business_communication.py). Complete the linked subject-pack drill before the engineering implementation; keep the university derivation and code evidence as separate artifacts.

## 11. Self-test

- [ ] State the valid score range and the exact passing rule.
- [ ] Why are 49 and 50 a more useful pair than 20 and 80 for checking the threshold?
- [ ] What is the expected result for 100? For 101?
- [ ] What does the check-case function compare?
- [ ] What process status means that every test passed in this program?
- [ ] How does the score-50 test expose the `>` mutation?

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
