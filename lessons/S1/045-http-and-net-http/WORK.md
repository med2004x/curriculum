# Lesson 045 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 5. Networking Foundations  
**Primary:** Networking/Go  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What components make up an HTTP/1.1 request message?
2. What does the blank line after headers mark?
3. How can HTTP/1.1 frame a message body?
4. What does `net/http` provide beyond a handler function?
5. How does a persistent HTTP connection differ from a TCP message boundary?
6. Why is content length measured in bytes?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Identify the request line, headers, header terminator, and body in an HTTP/1.1 request.
2. Predict the status for `/health` and an unknown path.
3. Predict how `Content-Length` differs from chunked transfer coding.
4. Predict whether HTTP/1.1 normally permits connection reuse.
5. Predict what occurs if a declared body length exceeds the bytes sent.
6. Predict why a UTF-8 body can have a byte length different from its character count.
7. State which parsing and connection behavior `net/http` manages.
8. Explain why one TCP read is not one HTTP request.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Explain the HTTP request/response structure and the role of method, target, headers, status, and body.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Distinguish body framing with `Content-Length` from HTTP/1.1 chunked transfer coding.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Explain persistent HTTP/1.1 connections and when a connection closes.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### HTTP service

Build a small `net/http` service with two resources, clear status codes, request logging, and tests that exercise malformed and valid requests.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Deliberately send a `Content-Length` header from your server that doesn't match the actual body's real byte length (e.g. claim 100 bytes but send only 50).

2. Test with `curl -v` and observe the client either hang waiting for the remaining, never-arriving bytes (if it's a keep-alive connection, since the client believes more of the response is still coming), or produce a truncated/corrupted body.

3. Explain, using Lesson 1's framing discussion directly, why this specific bug is a framing-boundary violation — the client is doing exactly what correct framing logic should do (trust the length prefix and read that many bytes), and the bug is entirely in your server providing an inaccurate length, not in the client's read logic.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint in the main case for # 045 — HTTP and net/http: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Formal model:** State the relevant relation, invariant, transaction property, or storage rule.
2. **Schedule/table:** Work one small example manually.
3. **Counterexample:** Construct one invalid or surprising case.
4. **Verification:** State the SQL command, trace, or test that would falsify your answer.

This lesson is also an academic integration point. Do the written drill before the engineering implementation so the two reinforce each other rather than becoming two unrelated memorisation exercises.

## 11. Self-test

- [ ] What components make up an HTTP/1.1 request message?
- [ ] What does the blank line after headers mark?
- [ ] How can HTTP/1.1 frame a message body?
- [ ] What does `net/http` provide beyond a handler function?
- [ ] How does a persistent HTTP connection differ from a TCP message boundary?
- [ ] Why is content length measured in bytes?

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
