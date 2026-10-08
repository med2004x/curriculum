# Lesson 014 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 2. Memory, Data Structures, and Algorithmic Thinking  
**Primary:** C  
**Supporting:** `malloc`; `realloc`; `free`; ownership; capacity; failure handling. **Math companion:** geometric growth and amortized work  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. Which object owns the allocation after a successful `vector_push`?
2. Why must `realloc` return into a temporary pointer?
3. What does the invariant say about `len`, `cap`, and `data`?
4. Why is the `vector_free` function safe to call twice?
5. What bug does LeakSanitizer detect in the first failure drill?
6. Why is a shallow copy of a populated `Vector` an ownership error?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict `len` and `cap` after pushes one through five.
2. Predict the result of reading index 8 after nine pushes.
3. Predict the result of reading index `len`.
4. Predict the vector state if `realloc` returns `NULL`.
5. Predict why the local variable `grown` is necessary.
6. Predict whether calling the correct `vector_free` twice frees twice.

Write the invariant values before running the test. Do not assume that `realloc` moves the buffer on every growth; correctness must hold whether it moves or extends in place.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **State who owns a dynamically allocated buffer and who must free it.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Grow a vector without losing its old allocation when `realloc` fails.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Enforce the invariant `0 <= len <= cap` and check indexed reads.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### Lifetime-safe container

Build a dynamically sized C container with explicit ownership rules. Add creation, insertion, removal, and destruction paths and verify allocation balance.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Remove the `vector_free` call in a test scope and run LeakSanitizer where available; identify the allocation that remains reachable only as a leak.
2. In a disposable broken copy, remove the `data = NULL` reset and call `vector_free` twice; run under AddressSanitizer and record the double-free report.
3. Assign `realloc` directly to `vector->data`; explain how allocation failure loses the owner’s only pointer.
4. Read index `len` without checking the return value; use the invariant to explain why it is out of bounds.

Do not rely on the unsanitized program crashing. Undefined behavior and allocator diagnostics vary; restore the correct implementation and rerun the clean test afterward.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change the vector element type from `int` to a small `struct Point { int x; int y; }`. Recalculate the allocation size using `sizeof *vector->data`, preserve the same owner/invariant rules, and test growth across capacity 4. Explain why the byte count changes but the capacity is still measured in elements.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Closed-book definition:** State the key formal definition from this lesson in one or two sentences.
2. **Worked derivation:** Draw the state transition or representation for the concrete example used in the lesson.
3. **Exam variation:** Change one parameter and recompute the result without running code.
4. **Engineering bridge:** Identify the exact observation in the terminal that would confirm your written result.

## 11. Self-test

- [ ] Which object owns the allocation after a successful `vector_push`?
- [ ] Why must `realloc` return into a temporary pointer?
- [ ] What does the invariant say about `len`, `cap`, and `data`?
- [ ] Why is the `vector_free` function safe to call twice?
- [ ] What bug does LeakSanitizer detect in the first failure drill?
- [ ] Why is a shallow copy of a populated `Vector` an ownership error?

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
