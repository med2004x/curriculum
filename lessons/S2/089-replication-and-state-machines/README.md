# Lesson 089 — Replication and State Machines

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand replicated state machines as the general pattern behind most fault-tolerant distributed systems, and directly experience the core consistency tradeoff: keeping replicas identical costs coordination, and skipping that coordination risks replicas disagreeing.

**The contract.** If replicas start from equivalent state and apply the same deterministic commands in the same order, they reach equivalent state. Replication must define who chooses the order, how entries become committed, and what happens when messages or nodes fail; simple copying alone does not provide consensus.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.

## 2. Mechanism and vocabulary

**Mechanism.** Keep `last_log_index` (received/persisted) distinct from `last_applied_index` (executed by the state machine). Apply entries strictly in order, and make command execution deterministic with respect to replicated state; local time, random choices, or unreplicated external reads can cause divergence.

**Second point.** Asynchronous replication may acknowledge before a follower persists/applies an entry, so a leader failure can lose acknowledged data under some failover policies. Synchronous replication can wait for specified acknowledgments, but the exact durability and availability guarantee depends on quorum/commit rules; do not infer consensus from a single follower acknowledgment.

**Worked example.** Test basic replication correctness: perform a sequence of writes against the leader, confirm each follower's local state exactly matches the leader's after all operations have propagated and been applied.

## 3. Worked case from the reference material

Worked example: Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.

## 4. Verification, failure, and evidence

**Verification.** For step 2, report the exact final state (all key-value pairs) from the leader and from each follower after your test sequence, confirming byte-for-byte agreement.

**Failure mode.** Using your asynchronous replication mode from step 4, perform a write against the leader, and — before it has propagated to any follower — simulate the leader crashing (kill the leader process or discard its state entirely) while at least one follower is still consistent with the leader's state *before* that specific write.

**Completion criterion.** You can replay a deterministic ordered log, track persistence versus application progress, recover a lagging follower, and state the acknowledged-data-loss limits of the chosen replication policy.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
