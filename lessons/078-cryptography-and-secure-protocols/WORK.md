# Lesson 078 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 9. Security, Database Internals, and Distributed Systems  
**Primary:** Security/Networking  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What does a cryptographic hash provide?
2. How does HMAC differ from an unkeyed hash?
3. How do digital-signature keys differ from HMAC keys?
4. Why is SHA-256 unsuitable for direct password storage?
5. What does constant-time MAC comparison address?
6. Why does a valid signature not by itself establish trust in a public key?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict which party can verify an HMAC and what key it needs.
2. Predict which party can create a signature and which key verifies it.
3. Predict what an altered message does to HMAC/signature verification.
4. Explain why a cryptographic hash does not conceal a predictable password.
5. Predict the effect of using a wrong public key.
6. Predict why constant-time MAC comparison is preferred.
7. State which key must remain secret in each construction.
8. Distinguish integrity, authenticity, confidentiality, and password-guessing resistance.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Distinguish cryptographic hashes, keyed MACs, and public-key signatures by their guarantees and key requirements.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Use standard-library or established-library APIs to compute and verify SHA-256, HMAC, and a digital signature.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Explain why password storage needs a salted, deliberately expensive password-hashing function rather than fast SHA-256.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### Secure protocol

Implement a small authenticated message exchange using a standard cryptographic primitive from the standard library or a vetted package. Explain the threat it solves and the one it does not.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Implement a deliberately naive "custom" password verification scheme: instead of using bcrypt's built-in, constant-time comparison (which handles the timing-safety concern automatically), implement your own comparison of a computed hash against a stored one using a standard, non-constant-time string comparison (`==` or equivalent).

2. This directly reuses API engineering Lesson 11's constant-time-comparison concern, now applied to password verification specifically rather than webhook signatures.

3. You don't need to demonstrate a successful timing attack (genuinely hard to do reliably in a small exercise), but research and explain, in your own words, why a naive comparison here reintroduces exactly the timing side-channel risk that lesson identified, and why using your standard library's provided, already-hardened comparison function (rather than writi.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one primitive: replace HMAC with a digital signature or change password-hash parameters. Explain which trust/key model changes and which guarantees remain.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Algorithm statement:** Write the input, output, invariant, and termination argument before coding.
2. **Dry run:** Solve one non-trivial example by hand and show each state change.
3. **Complexity:** Give worst-case time and extra-space bounds and justify them.
4. **Transfer:** Change one constraint and explain which part of the algorithm must change.

## 11. Self-test

- [ ] What does a cryptographic hash provide?
- [ ] How does HMAC differ from an unkeyed hash?
- [ ] How do digital-signature keys differ from HMAC keys?
- [ ] Why is SHA-256 unsuitable for direct password storage?
- [ ] What does constant-time MAC comparison address?
- [ ] Why does a valid signature not by itself establish trust in a public key?

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
