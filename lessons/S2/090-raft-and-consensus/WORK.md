# Lesson 090 Work Packet

**Lecture:** [`README.md`](README.md)  
**Module:** 9. Security, Database Internals, and Distributed Systems  
**Primary:** Distributed/Algorithms  
**Session target:** 2â€“4 hours  

This file is the practical half of the lesson. Do the lecture first. Do not treat the worked examples in the lecture as evidence that you can reproduce the skill from memory.

## 1. Session sequence

Use this exact order: **retrieve â†’ guided practice â†’ independent exercises â†’ problem solving â†’ coding challenge â†’ review â†’ evidence**. Skip a phase only when you have already demonstrated mastery of the lesson gate.

Suggested pacing: 10â€“15 min retrieval, 25â€“45 min guided practice, 45â€“75 min exercises, 30â€“60 min problem solving, 60â€“120 min challenge, 15 min review.

## 2. Retrieve before learning

1. What is the main objective stated for # 090 â€” Raft and Consensus?
2. What two concrete practices are required in the assigned material for # 090 â€” Raft and Consensus?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

Do not open the lecture until the predictions are written. Retrieval is not warm-up decoration. It is part of the lesson.

## 3. Guided practice

### Practice A â€” Predict
1. Predict how many votes are required in a five-node cluster.
2. Predict what a node does after receiving a higher term.
3. Predict whether a minority partition can commit a new entry.
4. Predict why two majorities intersect in a fixed cluster.
5. Predict how a stale log affects a RequestVote decision.
6. Predict what happens when AppendEntries reports a mismatched previous entry.
7. State which Raft guarantee is safety and which depends on liveness assumptions.
8. Identify the lab test evidence supporting each claim.

Run the smallest case first. Record prediction and observation side by side.

### Practice B â€” Modify one variable
Take the smallest runnable example from the lecture and change exactly one thing. Do not add new concepts yet. Predict the change, run it, and explain the delta in one paragraph.

### Practice C â€” Guided instrumentation
Instrument one internal state: an index, pointer, counter, transaction state, request ID, queue length, timestamp, or equivalent. The instrumentation must let you answer â€œwhat changed?â€ rather than merely â€œdid it work?â€

## 4. Independent exercises

### Exercise 1 â€” Recall and prediction
Without running anything, explain the main mechanism in **Explain Raft terms, randomized election timeouts, and majority-based leader election.**. Predict the result for one ordinary input and one boundary input. Record both predictions before checking them.

### Exercise 2 â€” Guided-to-independent implementation
Rebuild the smallest useful version of the lesson's main mechanism from a blank file. Your implementation must satisfy **State RequestVote and AppendEntries roles and the log-matching/safety invariants they support.**. Start with a written data/control-flow sketch, then code, then verify.

### Exercise 3 â€” Debug and repair
Take your implementation and introduce one realistic defect: wrong boundary, wrong state update, missing cleanup, incorrect error branch, or invalid assumption. Reproduce the failure, classify it, fix it, and add a regression test or evidence step.

### Exercise 4 â€” Transfer
Change one important constraint from the lesson. Examples: larger input, partial failure, concurrent callers, invalid data, persistence, or a different representation. Explain which invariant survives, which one changes, and why.

### Exercise 5 â€” Explanation
Write a 150â€“250 word technical explanation of **Explain why a committed log entry survives leader change under Raft's voting rules.** as if another student must implement it from your explanation. No screenshots; the explanation must stand on its own.

## 5. Problem-solving lab

Assume messages can be delayed, duplicated, or lost. State the guarantees you need, then design the minimum protocol state required to preserve them.

### Problem-solving protocol

1. Restate the problem and constraints.
2. Produce a small concrete example.
3. State the invariant or contract.
4. Write the simplest baseline solution or model.
5. Identify the bottleneck, failure, or repeated work.
6. Improve it only after you can explain the baseline.
7. Verify against a case designed to break the naive solution.

## 6. Coding challenge

### MIT 6.5840 Raft lab

Implement the official Raft lab in the matching MIT 6.5840 checkout. Use its supplied distributed tests for leader election, log replication, persistence, and unreliable-network behavior as assigned by the matching release. The local `exercise.go` is a small state-transition reading exercise, not the course lab.

### Acceptance criteria

- Follow one release's handout and tests without mixing test harnesses across years.
- Make the assigned tests pass and retain the upstream harness.
- Explain the safety invariant behind one failing-then-passing test.
- Record the upstream commit, lab part, exact test command, and output in `EVIDENCE.md`.
- Describe one limitation of the lab's fault model.


## 8. Debugging and failure lab

1. Attempt to deliberately break the "one vote per term per node" rule in your implementation (allow a node to vote for two different candidates within the same term) and rerun your step 1 election tests under network conditions likely to produce contested elections (e.g. multiple nodes becoming candidates near-simultaneously).

2. Attempt to reproduce a split-brain scenario â€” two nodes each believing they won the same term's election.

3. Explain, using the majority-overlap pigeonhole argument from Learn, precisely which specific guarantee your broken version violated, and why restoring the one-vote-per-term rule is what makes the overlap argument (and therefore the whole safety guarantee) hold again.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint in the main case for # 090 â€” Raft and Consensus: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Model:** State the system model and assumptions before solving.
2. **Trace:** Work a small failure or state-transition example by hand.
3. **Trade-off:** Compare two designs and state the condition under which each wins.
4. **Evidence:** Identify the measurement, test, or trace that would distinguish the designs in practice.

## 11. Self-test

- [ ] What is the main objective stated for # 090 â€” Raft and Consensus?
- [ ] What two concrete practices are required in the assigned material for # 090 â€” Raft and Consensus?
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

## Upstream lab: MIT 6.5840 Raft

Use the official [MIT 6.5840 lab source](git://g.csail.mit.edu/6.5840-golabs-2026) and [course guide](https://pdos.csail.mit.edu/6.5840/), with the handout and tests from the same release. In the supplied 2026 checkout, run `make -C src raft1`; its Makefile runs the race-enabled `raft1` Go tests. Follow the selected release's course page and Makefile, and record the commit, command, and results in `EVIDENCE.md`. The local Go file is a state-transition exercise, not a substitute for the replicated state machine lab.
