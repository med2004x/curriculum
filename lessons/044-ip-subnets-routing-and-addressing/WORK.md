# Lesson 044 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 5. Networking Foundations  
**Primary:** Networking/Math  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What does an IPv4 CIDR prefix specify?
2. How do you calculate a network address from an address and prefix?
3. What is longest-prefix match?
4. When is a destination on-link?
5. Which MAC address is used for an off-link packet's Ethernet frame?
6. What does ARP resolve, and what does routing decide?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Calculate the block size and range for `192.0.2.64/26`.
2. Predict whether `192.0.2.100` is on-link for that prefix.
3. Predict the next hop for an off-link destination.
4. Select between `/8`, `/16`, and default routes for `10.4.8.9`.
5. Predict which MAC address is resolved for an off-link packet on Ethernet.
6. State how IPv6 neighbor discovery differs from IPv4 ARP.
7. Predict what changes when a more-specific route is added.
8. Separate route selection from neighbor resolution in your explanation.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Calculate an IPv4 network address, broadcast address, and usable host range from a CIDR prefix.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Decide whether a destination is on-link or must be sent to a gateway.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Apply longest-prefix match to a routing table.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### Subnet calculator

Build a subnet calculator that accepts an IPv4 address and prefix, then reports network, broadcast, host range, and capacity with invalid-input checks.

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

Change one route or prefix: add a more-specific destination prefix, or change the local network from `/26` to `/27`. Recompute on-link membership and the longest-prefix route without changing the packet destination.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Formal model:** State the relevant relation, invariant, transaction property, or storage rule.
2. **Schedule/table:** Work one small example manually.
3. **Counterexample:** Construct one invalid or surprising case.
4. **Verification:** State the SQL command, trace, or test that would falsify your answer.

This lesson is also an academic integration point. Do the written drill before the engineering implementation so the two reinforce each other rather than becoming two unrelated memorisation exercises.

## 11. Self-test

- [ ] What does an IPv4 CIDR prefix specify?
- [ ] How do you calculate a network address from an address and prefix?
- [ ] What is longest-prefix match?
- [ ] When is a destination on-link?
- [ ] Which MAC address is used for an off-link packet's Ethernet frame?
- [ ] What does ARP resolve, and what does routing decide?

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
