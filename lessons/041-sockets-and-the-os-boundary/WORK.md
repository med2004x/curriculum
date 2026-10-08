# Lesson 041 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 5. Networking Foundations  
**Primary:** Networking/OS/Go  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What are the server and client socket lifecycle steps for TCP?
2. Why are TCP reads not message reads?
3. What framing fields identify a length-prefixed message?
4. How should the receiver handle a partial header or payload?
5. What should happen if EOF arrives before the declared length?
6. Why must the declared length be bounded?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict whether two writes map to exactly two reads.
2. Predict what `Read` returns when fewer bytes are currently available than requested.
3. Predict the receiver behavior when EOF arrives before the declared payload length.
4. Predict how an oversized length prefix should be handled.
5. Predict the listener, accepted-connection, and client socket roles.
6. Predict whether TCP preserves application message boundaries.
7. Predict which resources should be closed after a failed accept or read.
8. State the framing invariant before interpreting a read trace.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Explain how a socket endpoint is created and connected through the operating-system networking interface.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Trace the server lifecycle from bind/listen/accept and the client lifecycle from connect.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Explain why TCP reads do not correspond to sender write calls.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

Given a client/server trace with missing, delayed, or duplicated messages, reconstruct the state transitions and identify the first point where the application can no longer infer the desired state.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### Socket protocol

Implement a small client/server protocol over TCP with explicit framing and error responses. Test partial reads and client disconnects.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Write a *naive* receiver that assumes a single `read()` call always returns a complete message (no length-prefix framing, no read-loop — just one `read()` call per expected message) and test it against step 4's large-message, small-buffer-size scenario.

2. Confirm it receives only a fragment of the intended message (whatever fit in the first `read()` call) and either processes that truncated fragment as if it were complete, or errors out unexpectedly depending on what the fragment happens to contain.

3. Explain why this naive version might have appeared to work correctly during earlier, smaller-message testing (steps 1-2), lulling a developer into believing the assumption was safe, when in fact it was only ever true by coincidence of small message size and favorable timing, not because TCP guarantees it.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one framing rule: use a delimiter with escaping instead of a length prefix, or lower the maximum payload. Explain how the receiver recognizes message completion and handles malformed input.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Formal model:** State the relevant relation, invariant, transaction property, or storage rule.
2. **Schedule/table:** Work one small example manually.
3. **Counterexample:** Construct one invalid or surprising case.
4. **Verification:** State the SQL command, trace, or test that would falsify your answer.

This lesson is also an academic integration point. Do the written drill before the engineering implementation so the two reinforce each other rather than becoming two unrelated memorisation exercises.

## 11. Self-test

- [ ] What are the server and client socket lifecycle steps for TCP?
- [ ] Why are TCP reads not message reads?
- [ ] What framing fields identify a length-prefixed message?
- [ ] How should the receiver handle a partial header or payload?
- [ ] What should happen if EOF arrives before the declared length?
- [ ] Why must the declared length be bounded?

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
