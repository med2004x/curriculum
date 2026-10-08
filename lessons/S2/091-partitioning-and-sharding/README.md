# Lesson 091 — Partitioning and Sharding

## Threads
**Primary:** Distributed/DB

**Supporting:** hash/range partitioning; hotspots; rebalancing

**Math companion:** probability and hashing

**Time:** 3–4 hours. **Prerequisites:** Lesson 090.

**Time:** 3–4 hours. **Prerequisites:** Lesson 090.

## What you will be able to do

1. Compare modulo hashing, range partitioning, and consistent hashing.
2. Measure key movement and load balance when membership changes.
3. Explain virtual nodes, hotspots, and the operational cost of rebalancing.
4. Identify when a partition key creates skew or complicates queries/transactions.

## Retrieve first (no notes)

1. What does a Raft majority quorum provide?
2. Why can a minority partition not commit new log entries?
3. What do terms and one-vote-per-term rules help protect?
4. Why is a single leader/follower demo not enough to establish consensus?

## 1. Problem and model

**What problem this solves.** Distribute data and workload across partitions while balancing ownership, query locality, and reconfiguration cost. Consistent hashing can reduce movement for a ring-based assignment, but range/hash partitioning, directory maps, and database-native partitioning may be better for other workloads.

**The contract.** `hash(key) % M` can remap many keys when `M` changes. Consistent hashing places nodes and keys on a ring so adding/removing a node typically moves keys in the affected arc, subject to hash distribution and virtual-node configuration. Movement is not zero and imbalance remains possible.

**Worked example.** Start with the smallest concrete case described by the source material: Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Range partitioning can support range scans but risks hotspots when key distributions are skewed. Hash partitioning tends to spread keys but can make range queries scatter. Consistent hashing trades straightforward routing for ring metadata, virtual nodes, and explicit data movement/rebalancing.

**Second point.** A partition key affects not only ownership balance but also transaction locality, query fan-out, and hot-key behavior. Rebalancing must move data safely and coordinate ownership metadata; changing the lookup function alone does not migrate records.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Populate it with, say, 10,000 test keys across 4 nodes, then simulate adding a 5th node (recompute `% 5` for every key) and measure exactly what fraction of the 10,000 keys ended up mapped to a *different* node than before — report the actual percentage, and compare it to what you'd naively expect (should be a very large fraction, close to `(numNodes-1)/numN.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Produce a table comparing: naive modulo (step 1), plain consistent hashing (steps 2-3), and consistent hashing with virtual nodes (step 4) — for each, report the percentage of keys that moved when adding the 5th node, and the evenness of the resulting distribution across nodes.

**Failure mode.** Implement consistent hashing with a deliberately *small* number of virtual nodes per physical node (e.g. just 2, instead of 150) and rerun the distribution-evenness measurement from step 4.

**Completion criterion.** You can compare movement and balance for stated strategies, explain virtual-node trade-offs, and identify query/locality risks of the chosen partition key.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict how modulo ownership changes when node count changes from four to five.
2. Predict which ring interval moves when one node is added.
3. Predict how virtual-node count affects expected balance and metadata.
4. Predict how a monotonically increasing range key can hotspot range partitions.
5. Predict the query fan-out for a range scan across hash partitions.
6. Predict one hot-key failure that partition count alone does not solve.
7. State why key movement and load balance are separate measurements.
8. Identify the data transfer needed after ownership changes.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`.

2. Extend it with one explicit boundary case: Populate it with, say, 10,000 test keys across 4 nodes, then simulate adding a 5th node (recompute `% 5` for every key) and measure exactly what fraction of the 10,000 keys ended up mapped to a *different* node than before — report the actual percentage, and compare it to what you'd naively expect (should be a very large fraction, close to `(numNodes-1)/numN.

3. Rebuild the same behavior with one changed constraint: Implement consistent hashing: hash both nodes and keys onto a numeric ring (e.g. using a 32-bit or 64-bit hash function), and implement the "find the first node clockwise from the key's position" lookup (a sorted list of node positions plus a binary search, discrete math/algorithms track, is a natural implementation approach).

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Implement consistent hashing with a deliberately *small* number of virtual nodes per physical node (e.g. just 2, instead of 150) and rerun the distribution-evenness measurement from step 4.

2. Confirm the distribution is noticeably more uneven than with 150 virtual nodes (though still likely better than a single hash position with zero virtual nodes) — report the actual numbers.

3. Explain, using the law-of-large-numbers-style reasoning from Learn, why increasing the virtual-node count specifically improves evenness — you're averaging over more independent random ring positions per physical node, which (echoing probability track Lesson 4's Central Limit Theorem discussion, in spirit if not exact mechanism) reduces the variance of each.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 091 — Partitioning and Sharding: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 091 — Partitioning and Sharding?
2. What two concrete practices are required in the assigned material for # 091 — Partitioning and Sharding?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson091/EVIDENCE.md` with exactly this structure:

```text
# Lesson 091 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 091 — Partitioning and Sharding.
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

You pass when you can explain the central model of # 091 — Partitioning and Sharding without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've measured and directly compared naive modulo sharding's catastrophic key movement against consistent hashing's dramatically smaller movement when adding a node, using your own generated data rather than trusting the claim abstractly, and you've demonstrated that virtual nodes measurably improve distribution evenness, with numbers showing the effect sca. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
