# Lesson 085 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 9. Security, Database Internals, and Distributed Systems  
**Primary:** DB Internals/C/Algorithms  
**Session target:** 3–6 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What is the main objective stated for # 085 — Database Internals Lab?
2. What two concrete practices are required in the assigned material for # 085 — Database Internals Lab?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. State the chosen component's invariant.
2. Predict what fails if a pinned frame is evicted.
3. Predict the observable behavior of an iterator after exhaustion.
4. Predict which component owns page serialization and which owns frame lifetime.
5. Predict how the test verifies a split or recovery boundary.
6. Predict one measurement affected by the chosen workload.
7. Distinguish a correctness test from a performance experiment.
8. State the environment assumptions before interpreting results.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Select one database component boundary: page storage, buffer pool, B+ tree, iterator execution, or WAL/recovery.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **State invariants and interfaces before implementing or instrumenting that component.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Integrate it with at least one neighboring component and verify their contract together.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

Given a small relational dataset and a business question, write the query or transaction schedule on paper first. Then verify it against a deliberately chosen edge case that would break a naive solution.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### CMU BusTub project

Select the matching BusTub project on pages, buffer pools, B+ trees, query execution, or transactions from the semester that matches the local upstream checkout. Implement the official project API and use the supplied tests. `exercise.c` is a concept prompt only; it is not a miniature BusTub substitute.

### Acceptance criteria

- Follow the official project handout and preserve its test harness.
- State the component invariant and how its neighboring component depends on it.
- Run the supplied project tests plus one failure or boundary case.
- Record the semester, upstream commit, exact build/test commands, and outcomes in `EVIDENCE.md`.
- Use measurements from the supplied workload to explain one performance or correctness tradeoff.

## 8. Debugging and failure lab

1. Take your integrated system (or the BusTub project you completed) and deliberately break one cross-component assumption — for example, if you built your own system, bypass the buffer pool's pin discipline in exactly one code path (directly access a page without pinning it first) while the rest of the system correctly uses pins.

2. Run your workload under concurrent access (multiple goroutines/threads performing operations simultaneously) and attempt to reproduce a corruption or crash caused specifically by this one broken assumption — e.g. a page being evicted while the unprotected code path is still reading it, per Lesson 2's pin/unpin discussion.

3. If you cannot reliably reproduce a visible failure (this specific race can be timing-dependent and hard to force), document your reasoning for why the assumption violation is still a genuine bug even if your specific test run didn't happen to trigger a visible symptom — directly echoing C track Lesson 5's core point about undefined behavior not requiring a v.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint in the main case for # 085 — Database Internals Lab: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Formal model:** State the relevant relation, invariant, transaction property, or storage rule.
2. **Schedule/table:** Work one small example manually.
3. **Counterexample:** Construct one invalid or surprising case.
4. **Verification:** State the SQL command, trace, or test that would falsify your answer.

## 11. Self-test

- [ ] What is the main objective stated for # 085 — Database Internals Lab?
- [ ] What two concrete practices are required in the assigned material for # 085 — Database Internals Lab?
- [ ] What does the lesson ask you to reproduce before making changes?
- [ ] What failure mode is explicitly used as a diagnostic exercise?
- [ ] What evidence does the verification step require you to record?
- [ ] What condition in the "Done when" criterion proves mastery of this lesson?

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
