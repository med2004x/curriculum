# Lesson 012 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 1. Machine and Programming Foundations  
**Primary:** C/Linux  
**Supporting:** arguments; files; parsing; arrays; functions; tests; Git  
**Math companion:** problem decomposition  
**Session target:** 3–6 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve → guided practice → independent exercises → problem solving → coding challenge → review → evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10–15 min retrieval, 25–45 min guided practice, 45–75 min exercises, 30–60 min problem solving, 60–120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What are the four valid fields or constraints in one log record?
2. Why does the parser receive a writable line buffer?
3. Which channel receives parse errors, and what process status follows any malformed line?
4. Why does a search with no matches still return status zero?
5. Which source file owns file I/O, and which one owns record parsing?
6. What should happen when the record capacity of 1,000 is exceeded?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A — Predict
1. Predict which fields the parser extracts from the first sample record.
2. Predict whether `disk` and `Disk` both match.
3. Predict the output stream and exit status for a missing file.
4. Predict whether a malformed middle line prevents later valid records from being searched.
5. Predict what happens when a file contains 1,001 valid records.
6. Predict which unit test detects accepting `BOGUS` as a level.

Record each prediction before running. Compare stdout, stderr, and status separately; they are distinct observations.

### Practice B — Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C — Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer “what changed?” rather than merely “did it work?”

## 4. Independent exercises

### Exercise 1 — Recall and prediction
Without running anything, explain the main mechanism in **Turn a command-line contract into a multi-file C program.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 — Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **Parse a delimited record without confusing malformed input with valid data.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 — Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 — Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 — Explanation
Write a 150–250 word technical explanation of **Keep a bounded set of records in memory and search it by message text.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

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

### Integrated C utility

Build a small CSV-like command-line data tool using the earlier C concepts. It must parse input, maintain records, compute a summary, and report malformed rows.

### Acceptance criteria

- Start from a clean workspace or the latest lesson project without copying a completed solution.
- Implement the required behavior end-to-end.
- Include automated tests where the technology allows it, plus one deliberate failure test.
- Record design assumptions and the main invariant(s) in Markdown.
- Demonstrate the project from a reproducible command sequence.
- Finish with a short post-mortem: what failed, why, and what you changed.

## 8. Debugging and failure lab

1. Treat `BOGUS` as `INFO`; confirm the invalid-level test fails.
2. Remove the second separator from one input; confirm parsing rejects it.
3. Search for an absent term; confirm status is still zero.
4. Rename the input file; confirm the filename diagnostic appears on stderr and the process exits nonzero.
5. Make a line longer than the input buffer; confirm it is rejected as one malformed line, not read as two records.

For each experiment, identify whether the defect is in parsing, file I/O, output routing, or command-line validation. Restore the source and rerun all tests.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change the record format to `LEVEL|timestamp|message` while preserving all other behavior. Update the parser tests first, including one valid record and one invalid level. Only then change the implementation. Your evidence must show the same search results on equivalent data and a nonzero status for malformed input.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Closed-book definition:** State the key formal definition from this lesson in one or two sentences.
2. **Worked derivation:** Draw the state transition or representation for the concrete example used in the lesson.
3. **Exam variation:** Change one parameter and recompute the result without running code.
4. **Engineering bridge:** Identify the exact observation in the terminal that would confirm your written result.

LIG support: [ACCOUNTING](../S1/s1_accounting.py). Complete the linked subject-pack drill before the engineering implementation; keep the university derivation and code evidence as separate artifacts.

## 11. Self-test

- [ ] What are the four valid fields or constraints in one log record?
- [ ] Why does the parser receive a writable line buffer?
- [ ] Which channel receives parse errors, and what process status follows any malformed line?
- [ ] Why does a search with no matches still return status zero?
- [ ] Which source file owns file I/O, and which one owns record parsing?
- [ ] What should happen when the record capacity of 1,000 is exceeded?

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
