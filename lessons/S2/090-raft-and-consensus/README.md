# Lesson 090 — Raft and Consensus

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand and implement core components of the Raft consensus algorithm — leader election and log replication — the rigorous protocol that solves Lesson 4's "how do you correctly elect a new leader after the old one fails" problem, which that lesson deliberately left unaddressed.

**The contract.** A candidate increments its term, votes for itself, and requests votes. A server grants at most one vote per term and checks candidate log freshness. A candidate becomes leader only after receiving votes from a majority. Terms and these voting rules are central to election safety; a timeout alone does not prove the old leader failed.

**Worked example.** Start with the smallest concrete case described by the source material: MIT's 6.5840 (referenced throughout this track) provides a rigorous, well-tested lab sequence for implementing Raft incrementally, with a real test harness that specifically injects network partitions, delays, and crashes — using it directly is strongly preferred over reconstructing an equivalent test harness from scratch, since correctly testing a consensus.

## 2. Mechanism and vocabulary

**Mechanism.** Raft uses a majority quorum. Any two majorities of the same fixed cluster intersect, and the one-vote-per-term plus log-up-to-date rules help ensure a later leader contains prior committed entries. Election safety is at most one leader per term; progress requires an available majority and timely communication.

**Second point.** AppendEntries replicates log entries and also serves as a heartbeat. A follower checks the previous log index/term, rejecting inconsistent prefixes so the leader can repair the follower's log.

**Worked example.** Implement Raft leader election (the course's first Raft lab): nodes start as followers, time out and become candidates, request votes, and a majority-vote winner becomes leader for that term.

## 3. Worked case from the reference material

Worked example: MIT's 6.5840 (referenced throughout this track) provides a rigorous, well-tested lab sequence for implementing Raft incrementally, with a real test harness that specifically injects network partitions, delays, and crashes — using it directly is strongly preferred over reconstructing an equivalent test harness from scratch, since correctly testing a consensus.

## 4. Verification, failure, and evidence

**Verification.** Use the lab's partition/crash harness to record election and commit outcomes for majority and minority partitions, then verify a reconnected follower catches up. State the tested cluster size and quorum; passing tests provide evidence for covered schedules, not a universal proof by themselves.

**Failure mode.** Attempt to deliberately break the "one vote per term per node" rule in your implementation (allow a node to vote for two different candidates within the same term) and rerun your step 1 election tests under network conditions likely to produce contested elections (e.g. multiple nodes becoming candidates near-simultaneously).

**Completion criterion.** You can explain term/vote/log invariants, trace majority election and log repair, and identify the availability limit imposed by quorum under partition.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
