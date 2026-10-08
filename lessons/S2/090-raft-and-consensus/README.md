# Lesson 090 — Raft and Consensus

## Threads
**Primary:** Distributed/Algorithms

**Supporting:** leader election; log replication; safety/liveness

**Math companion:** proofs and invariants

**Time:** 3–4 hours. **Prerequisites:** Lesson 089.

**Time:** 3–4 hours. **Prerequisites:** Lesson 089.

## What you will be able to do

1. Explain Raft terms, randomized election timeouts, and majority-based leader election.
2. State RequestVote and AppendEntries roles and the log-matching/safety invariants they support.
3. Explain why a committed log entry survives leader change under Raft's voting rules.
4. Distinguish a majority partition that can make progress from a minority partition that cannot commit new entries.

## Retrieve first (no notes)

1. Why must replicas apply the same deterministic command sequence?
2. What is the difference between persisted log index and applied index?
3. What data loss can asynchronous replication permit?
4. Why is a simple leader/follower demo not consensus?

## 1. Problem and model

**What problem this solves.** Understand and implement core components of the Raft consensus algorithm — leader election and log replication — the rigorous protocol that solves Lesson 4's "how do you correctly elect a new leader after the old one fails" problem, which that lesson deliberately left unaddressed.

**The contract.** A candidate increments its term, votes for itself, and requests votes. A server grants at most one vote per term and checks candidate log freshness. A candidate becomes leader only after receiving votes from a majority. Terms and these voting rules are central to election safety; a timeout alone does not prove the old leader failed.

**Worked example.** Start with the smallest concrete case described by the source material: MIT's 6.5840 (referenced throughout this track) provides a rigorous, well-tested lab sequence for implementing Raft incrementally, with a real test harness that specifically injects network partitions, delays, and crashes — using it directly is strongly preferred over reconstructing an equivalent test harness from scratch, since correctly testing a consensus.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Raft uses a majority quorum. Any two majorities of the same fixed cluster intersect, and the one-vote-per-term plus log-up-to-date rules help ensure a later leader contains prior committed entries. Election safety is at most one leader per term; progress requires an available majority and timely communication.

**Second point.** AppendEntries replicates log entries and also serves as a heartbeat. A follower checks the previous log index/term, rejecting inconsistent prefixes so the leader can repair the follower's log.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Implement Raft leader election (the course's first Raft lab): nodes start as followers, time out and become candidates, request votes, and a majority-vote winner becomes leader for that term.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: MIT's 6.5840 (referenced throughout this track) provides a rigorous, well-tested lab sequence for implementing Raft incrementally, with a real test harness that specifically injects network partitions, delays, and crashes — using it directly is strongly preferred over reconstructing an equivalent test harness from scratch, since correctly testing a consensus.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Use the lab's partition/crash harness to record election and commit outcomes for majority and minority partitions, then verify a reconnected follower catches up. State the tested cluster size and quorum; passing tests provide evidence for covered schedules, not a universal proof by themselves.

**Failure mode.** Attempt to deliberately break the "one vote per term per node" rule in your implementation (allow a node to vote for two different candidates within the same term) and rerun your step 1 election tests under network conditions likely to produce contested elections (e.g. multiple nodes becoming candidates near-simultaneously).

**Completion criterion.** You can explain term/vote/log invariants, trace majority election and log repair, and identify the availability limit imposed by quorum under partition.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict how many votes are required in a five-node cluster.
2. Predict what a node does after receiving a higher term.
3. Predict whether a minority partition can commit a new entry.
4. Predict why two majorities intersect in a fixed cluster.
5. Predict how a stale log affects a RequestVote decision.
6. Predict what happens when AppendEntries reports a mismatched previous entry.
7. State which Raft guarantee is safety and which depends on liveness assumptions.
8. Identify the lab test evidence supporting each claim.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: MIT's 6.5840 (referenced throughout this track) provides a rigorous, well-tested lab sequence for implementing Raft incrementally, with a real test harness that specifically injects network partitions, delays, and crashes — using it directly is strongly preferred over reconstructing an equivalent test harness from scratch, since correctly testing a consensus.

2. Extend it with one explicit boundary case: Implement Raft leader election (the course's first Raft lab): nodes start as followers, time out and become candidates, request votes, and a majority-vote winner becomes leader for that term.

3. Rebuild the same behavior with one changed constraint: Test it under the course's provided test harness, including scenarios where the current leader is forcibly disconnected and a new election must correctly occur.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Attempt to deliberately break the "one vote per term per node" rule in your implementation (allow a node to vote for two different candidates within the same term) and rerun your step 1 election tests under network conditions likely to produce contested elections (e.g. multiple nodes becoming candidates near-simultaneously).

2. Attempt to reproduce a split-brain scenario — two nodes each believing they won the same term's election.

3. Explain, using the majority-overlap pigeonhole argument from Learn, precisely which specific guarantee your broken version violated, and why restoring the one-vote-per-term rule is what makes the overlap argument (and therefore the whole safety guarantee) hold again.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 090 — Raft and Consensus: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 090 — Raft and Consensus?
2. What two concrete practices are required in the assigned material for # 090 — Raft and Consensus?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson090/EVIDENCE.md` with exactly this structure:

```text
# Lesson 090 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 090 — Raft and Consensus.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 090 — Raft and Consensus without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've implemented and passed the test harness's leader election and log replication tests, including under simulated partitions and crashes, you've hand-traced a concrete 5-node scenario confirming why a newly-elected leader necessarily has all previously-committed entries, and you've directly observed a minority partition correctly failing to make progress. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
