# Lesson 089 — Replication and State Machines

## Threads
**Primary:** Distributed/DB

**Supporting:** replicas; logs; deterministic state machines

**Math companion:** invariants

**Time:** 3–4 hours. **Prerequisites:** Lesson 088.

**Time:** 3–4 hours. **Prerequisites:** Lesson 088.

## What you will be able to do

1. Explain why replicas applying the same deterministic command sequence reach the same state.
2. Track log index, applied index, and state-machine state separately.
3. Recover a lagging follower by replaying missing ordered log entries.
4. Explain what asynchronous replication can lose and why a toy leader/follower demo is not consensus.

## Retrieve first (no notes)

1. What does happened-before capture?
2. What does a Lamport timestamp guarantee for causally ordered events?
3. Can Lamport timestamps distinguish all concurrent event pairs?
4. How can wall-clock skew reverse the apparent order of causal events?

## 1. Problem and model

**What problem this solves.** Understand replicated state machines as the general pattern behind most fault-tolerant distributed systems, and directly experience the core consistency tradeoff: keeping replicas identical costs coordination, and skipping that coordination risks replicas disagreeing.

**The contract.** If replicas start from equivalent state and apply the same deterministic commands in the same order, they reach equivalent state. Replication must define who chooses the order, how entries become committed, and what happens when messages or nodes fail; simple copying alone does not provide consensus.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Keep `last_log_index` (received/persisted) distinct from `last_applied_index` (executed by the state machine). Apply entries strictly in order, and make command execution deterministic with respect to replicated state; local time, random choices, or unreplicated external reads can cause divergence.

**Second point.** Asynchronous replication may acknowledge before a follower persists/applies an entry, so a leader failure can lose acknowledged data under some failover policies. Synchronous replication can wait for specified acknowledgments, but the exact durability and availability guarantee depends on quorum/commit rules; do not infer consensus from a single follower acknowledgment.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Test basic replication correctness: perform a sequence of writes against the leader, confirm each follower's local state exactly matches the leader's after all operations have propagated and been applied.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** For step 2, report the exact final state (all key-value pairs) from the leader and from each follower after your test sequence, confirming byte-for-byte agreement.

**Failure mode.** Using your asynchronous replication mode from step 4, perform a write against the leader, and — before it has propagated to any follower — simulate the leader crashing (kill the leader process or discard its state entirely) while at least one follower is still consistent with the leader's state *before* that specific write.

**Completion criterion.** You can replay a deterministic ordered log, track persistence versus application progress, recover a lagging follower, and state the acknowledged-data-loss limits of the chosen replication policy.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict replica state after replaying an identical ordered command log.
2. Predict what happens if an entry is applied twice without command deduplication.
3. Distinguish last received/persisted index from last applied index.
4. Predict how nondeterministic local time can make replicas diverge.
5. Predict which acknowledged writes can be lost under asynchronous replication.
6. Predict what data a lagging follower needs to catch up.
7. Predict what a single follower acknowledgment does not prove about consensus.
8. State the log-order and durability assumptions behind the prediction.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.

2. Extend it with one explicit boundary case: Test basic replication correctness: perform a sequence of writes against the leader, confirm each follower's local state exactly matches the leader's after all operations have propagated and been applied.

3. Rebuild the same behavior with one changed constraint: Simulate a follower crash and recovery: stop a follower process partway through a sequence of leader operations (so it misses some), then restart it and implement a simple catch-up mechanism (the follower requests all log entries after the last one it successfully applied, and replays them in order) — confirm the recovered follower's state correctly converge.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Using your asynchronous replication mode from step 4, perform a write against the leader, and — before it has propagated to any follower — simulate the leader crashing (kill the leader process or discard its state entirely) while at least one follower is still consistent with the leader's state *before* that specific write.

2. Confirm the write is now permanently lost — no follower has it, and the leader (if restarted from scratch, with no separately-persisted log of its own) has no record of it either.

3. Explain why this is not a bug in your implementation but the precise, expected cost of asynchronous replication's tradeoff — the client was told the write succeeded (since asynchronous mode acknowledges before propagation), yet the write is now gone, which is exactly the kind of "confirmed but actually lost" scenario synchronous replication (at the cost of t.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 089 — Replication and State Machines: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 089 — Replication and State Machines?
2. What two concrete practices are required in the assigned material for # 089 — Replication and State Machines?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson089/EVIDENCE.md` with exactly this structure:

```text
# Lesson 089 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 089 — Replication and State Machines.
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

You pass when you can explain the central model of # 089 — Replication and State Machines without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your replicated state machine correctly converges follower state to match the leader across a sequence of operations, you've demonstrated a follower correctly catching up after a simulated crash and restart, and you've directly reproduced — not just read about — the specific data-loss scenario asynchronous replication risks, with your own measured latency nu. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
