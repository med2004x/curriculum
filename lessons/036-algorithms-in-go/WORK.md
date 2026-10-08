# Lesson 036 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 4. Go and Core Algorithms  
**Primary:** Go/Algorithms  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What is linear search's precondition and worst-case cost?
2. What is binary search's precondition and comparison cost?
3. What does the half-open interval `[low, high)` represent?
4. Why must the binary-search interval strictly shrink?
5. What two properties define a correct sorting result?
6. What are insertion sort's worst-case cost and useful input pattern?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Trace binary search for target `7` in `[2, 4, 7, 9, 13]`.
2. Predict the result for a target smaller than every element.
3. Predict the result on an empty slice.
4. State why the interval is half-open and what `high` means.
5. Predict how many comparisons linear search needs for a missing value at size `n`.
6. State the ordering and permutation conditions for a correct sort.
7. Predict which sorting case is worst for insertion sort.
8. Identify the input assumption that makes binary search valid.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Implement linear search and binary search over integer slices in Go.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **State and preserve the sorted-input precondition for binary search.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Implement a sorting algorithm with an explicit permutation-and-order postcondition.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

Solve an unseen problem by writing a brute-force baseline first. Then identify the repeated work or structure and derive a better approach. Compare both on the same examples.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### Algorithm toolkit

Implement three standard algorithms in Go, instrument comparisons or allocations, and write a short complexity justification for each.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one precondition: allow duplicate values or search for the first occurrence rather than any occurrence. State how the postcondition and boundary tests change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Algorithm statement:** Write the input, output, invariant, and termination argument before coding.
2. **Dry run:** Solve one non-trivial example by hand and show each state change.
3. **Complexity:** Give worst-case time and extra-space bounds and justify them.
4. **Transfer:** Change one constraint and explain which part of the algorithm must change.

## 11. Self-test

- [ ] What is linear search's precondition and worst-case cost?
- [ ] What is binary search's precondition and comparison cost?
- [ ] What does the half-open interval `[low, high)` represent?
- [ ] Why must the binary-search interval strictly shrink?
- [ ] What two properties define a correct sorting result?
- [ ] What are insertion sort's worst-case cost and useful input pattern?

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
