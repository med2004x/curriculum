# Lesson 079 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 9. Security, Database Internals, and Distributed Systems  
**Primary:** Security/API/DB  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. Why do SQL parameters keep values separate from query syntax?
2. Why is HTML escaping context-sensitive?
3. What browser behavior does CSRF rely on?
4. What makes a server-side URL fetch a trust boundary?
5. Why must SSRF defenses consider redirects and resolved addresses?
6. Why must every output sink be reviewed independently?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict whether a SQL parameter can alter the query's syntax tree.
2. Predict how contextual HTML escaping changes a synthetic marker in a local page.
3. Predict which state-changing request is protected by the selected CSRF policy.
4. Predict whether a server-side fetch to a disallowed destination should proceed.
5. Predict how a redirect can affect destination validation.
6. Predict whether escaping one template location protects a second output sink.
7. Identify the trust boundary for each test fixture.
8. State the expected before/after evidence before patching.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Identify trust-boundary crossings for SQL, HTML, browser state-changing requests, and server-side URL fetches.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Reproduce each vulnerability only in a local toy application with synthetic data.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Apply parameterized SQL, contextual output escaping, CSRF defenses, and restricted outbound destinations.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

Assume an attacker controls one untrusted input or one lower-trust component. Identify the trust boundary, attacker goal, violated invariant, and smallest preventive control. Verify the control with a test.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### Web vulnerability lab

Take a deliberately unsafe web endpoint, reproduce a data or authorization flaw, patch it, and write a regression test that would fail before the fix.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Take your step 2 XSS fix (proper output escaping) and construct a *second*, different injection point in the same application that you didn't originally fix (e.g. if you fixed escaping in a comment-rendering template, try injecting via a different field, like a username, rendered in a different part of the page that uses a different, unescaped rendering path.

2. Confirm this second injection point is still vulnerable despite your first fix, demonstrating that fixing one specific instance of a vulnerability class doesn't automatically fix every instance — each place untrusted data crosses into a rendering (or query, or fetch) context needs its own correct handling, and a security review needs to systematically check.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint in the main case for # 079 — Web Vulnerabilities and Data Safety: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Contract recall:** State the protocol/API/security rule without notes.
2. **Scenario analysis:** Given one malformed or adversarial request, determine the expected system response.
3. **Why:** Explain the invariant or security property that makes that response correct.
4. **Implementation bridge:** Name the test or runtime observation that would prove the rule is enforced.

## 11. Self-test

- [ ] Why do SQL parameters keep values separate from query syntax?
- [ ] Why is HTML escaping context-sensitive?
- [ ] What browser behavior does CSRF rely on?
- [ ] What makes a server-side URL fetch a trust boundary?
- [ ] Why must SSRF defenses consider redirects and resolved addresses?
- [ ] Why must every output sink be reviewed independently?

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
