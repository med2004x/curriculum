# Lesson 076 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 8. Backend Engineering  
**Primary:** Go/DB/API/Operations  
**Session target:** 3–6 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What is the selected resource and its invariant?
2. Which API errors are part of the public contract?
3. Which features must compose in the cross-feature integration test?
4. What runtime configuration and dependencies are required?
5. How does the service handle cancellation and shutdown?
6. What failure did the deliberate interaction test reveal?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict each layer's responsibility for one create request.
2. Predict the behavior for invalid input and unauthorized resource access.
3. Predict which integration tests need PostgreSQL.
4. Predict the stored state after a failed transaction.
5. Predict the response to a duplicate idempotency key.
6. Predict what happens when a dependency is unavailable at startup.
7. Predict the shutdown order for server, workers, and database pool.
8. Identify a cross-feature interleaving that can violate an invariant.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Design and implement one production-shaped Go API over PostgreSQL from an explicit contract.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Integrate authentication/authorization, validation, transactions, idempotency or version checks, and observability.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Verify isolated integration behavior, migrations, container configuration, and graceful shutdown.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### Production backend capstone

Build a production-shaped Go service combining routing, validation, persistence, authentication, idempotency, async work, observability, integration tests, and deployment notes.

### Acceptance criteria

- Start from a clean workspace or the latest lesson project without copying a completed solution.
- Implement the required behavior end-to-end.
- Include automated tests where the technology allows it, plus one deliberate failure test.
- Record design assumptions and the main invariant(s) in Markdown.
- Demonstrate the project from a reproducible command sequence.
- Finish with a short post-mortem: what failed, why, and what you changed.

## 8. Debugging and failure lab

1. Pick one specific cross-feature interaction from your capstone and deliberately break it — for example, make your rate limiter apply *before* authentication (checking IP-based limits before knowing who the authenticated user is, contradicting Lesson 10's failure drill about per-identity limiting being preferable to per-IP), or make your async job endpoint (i.

2. Demonstrate the resulting vulnerability or incorrect behavior with a real test, then fix it, and explain in your own words why integration points between separately-correct features are disproportionately likely to hide bugs, compared to bugs within any single feature's own isolated logic.

3. Deliver a portion of your talk from memorized wording, then have someone interrupt you mid-sentence with an unrelated question, then ask you to continue from where you left off.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one contract requirement, such as adding a status transition or a new access policy. Update schema/migration, service behavior, API error contract, and tests together while preserving unrelated interfaces.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Contract recall:** State the protocol/API/security rule without notes.
2. **Scenario analysis:** Given one malformed or adversarial request, determine the expected system response.
3. **Why:** Explain the invariant or security property that makes that response correct.
4. **Implementation bridge:** Name the test or runtime observation that would prove the rule is enforced.

This lesson is also an academic integration point. Do the written drill before the engineering implementation so the two reinforce each other rather than becoming two unrelated memorisation exercises.

## 11. Self-test

- [ ] What is the selected resource and its invariant?
- [ ] Which API errors are part of the public contract?
- [ ] Which features must compose in the cross-feature integration test?
- [ ] What runtime configuration and dependencies are required?
- [ ] How does the service handle cancellation and shutdown?
- [ ] What failure did the deliberate interaction test reveal?

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
