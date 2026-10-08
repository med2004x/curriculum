# Lesson 088 — Clocks, Ordering, and Logical Time

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand why wall-clock timestamps cannot reliably order events across different machines, and implement Lamport logical clocks to establish a causally-consistent ordering instead.

**The contract.** Happened-before (`→`) is a causal partial order: program order and message send-before-receive imply order, and the relation is transitive. If neither `a → b` nor `b → a`, the events are concurrent. Lamport clocks guarantee `a → b` implies `L(a) < L(b)`, but not the converse.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a Lamport clock as a small class/struct with a `tick()` method (increment for a local event, return new value) and a `receive(remoteTimestamp)` method (implementing the max-plus-one rule from Learn).

## 2. Mechanism and vocabulary

**Mechanism.** Each process increments its local counter before a local event. On receiving a message with timestamp `t`, update local time to `max(local,t)+1`; include the resulting counter in outgoing messages.

**Second point.** Lamport timestamps can produce a deterministic total ordering when combined with a process ID, but that order does not prove events were causally related. Physical clocks estimate real time and require synchronization assumptions and uncertainty bounds.

**Worked example.** Write a unit test confirming the basic increment behavior in isolation.

## 3. Worked case from the reference material

Worked example: Implement a Lamport clock as a small class/struct with a `tick()` method (increment for a local event, return new value) and a `receive(remoteTimestamp)` method (implementing the max-plus-one rule from Learn).

## 4. Verification, failure, and evidence

**Verification.** Produce an event table containing process, event, message relation, Lamport timestamp, and happened-before/concurrent classification. Confirm every causal edge has increasing Lamport time and include a concurrent pair whose timestamps are still numerically ordered.

**Failure mode.** Assign different synthetic wall-clock offsets to two processes and show that wall-clock order can contradict a known message-causality edge. Explain that Lamport time preserves that edge but does not measure elapsed time.

**Completion criterion.** You can implement Lamport updates, classify causality versus concurrency from an event graph, and explain the limits of logical and wall-clock timestamps.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
