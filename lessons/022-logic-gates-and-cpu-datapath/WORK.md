# Lesson 022 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 3. Architecture, Linux, Processes, and Concurrency  
**Primary:** Architecture/C  
**Session target:** 2–4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What is NAND's Boolean expression and truth table?
2. How can NAND construct NOT, AND, and OR?
3. Write the `Sum` and `Cout` equations for a full-adder.
4. What connects neighboring stages in a ripple-carry adder?
5. How does unsigned carry-out differ from signed overflow?
6. How does a latch differ from combinational logic?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict NAND's output for all four input pairs.
2. Predict `NAND(A,A)` for `A=0` and `A=1`.
3. Predict `Sum` and `Cout` for `A=1, B=1, Cin=1`.
4. Trace `0110 + 0101` from least-significant to most-significant bit.
5. Predict the four-bit result and final carry of `1111 + 0001`.
6. Predict what is lost if each stage's `Cin` is tied to zero.
7. Does `0111 + 0001` produce signed overflow, unsigned carry-out, both, or neither?
8. Record the truth table and carry trace before checking a simulator.

Run the smallest case first. Record prediction and observation side by side.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Construct NOT, AND, OR, and XOR using NAND gates.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Derive half-adder and full-adder outputs from Boolean expressions and truth tables.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Trace a four-bit ripple-carry addition, including every carry bit.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### Tiny ALU model

Model a four-operation ALU from Boolean building blocks or code that explicitly mirrors gates. Verify truth tables and explain the datapath.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Chain the full adders without connecting Cin from one stage to Cout of the previous stage — leave each Cin tied to 0.

2. Compute `0110 + 0101` again with this broken wiring.

3. Explain exactly which additions now produce a wrong 4-bit result and why, and connect this to what "ripple carry" actually means — the propagation delay per stage is the source of the ripple-carry adder's speed limit, which is why real CPUs use carry-lookahead adders instead.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint: widen the adder from four bits to eight, or change one input pair. State how many full-adder stages are needed and trace the new carry path.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Closed-book definition:** State the key formal definition from this lesson in one or two sentences.
2. **Worked derivation:** Draw the state transition or representation for the concrete example used in the lesson.
3. **Exam variation:** Change one parameter and recompute the result without running code.
4. **Engineering bridge:** Identify the exact observation in the terminal that would confirm your written result.

LIG support: [LOGIC-ARCH](../../LIG/S1/LOGIC-ARCH.md). Complete the linked subject-pack drill before the engineering implementation; keep the university derivation and code evidence as separate artifacts.

## 11. Self-test

- [ ] What is NAND's Boolean expression and truth table?
- [ ] How can NAND construct NOT, AND, and OR?
- [ ] Write the `Sum` and `Cout` equations for a full-adder.
- [ ] What connects neighboring stages in a ripple-carry adder?
- [ ] How does unsigned carry-out differ from signed overflow?
- [ ] How does a latch differ from combinational logic?

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
