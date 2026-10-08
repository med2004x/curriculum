# Lesson 090 — Problem Set

**Topic:** Raft and Consensus.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: MIT's 6.5840 (referenced throughout this track) provides a rigorous, well-tested lab sequence for implementing Raft incrementally, with a real test harness that specifically injects network partitions, delays, and crashes — using it directly is strongly preferred over reconstructing an equivalent test harness from scratch, since correctly testing a consensus.

2. Extend it with one explicit boundary case: Implement Raft leader election (the course's first Raft lab): nodes start as followers, time out and become candidates, request votes, and a majority-vote winner becomes leader for that term.

3. Rebuild the same behavior with one changed constraint: Test it under the course's provided test harness, including scenarios where the current leader is forcibly disconnected and a new election must correctly occur.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/090-raft-and-consensus/code
go test -race exercise.go exercise_test.go
```

## Stretch problem

### MIT 6.5840 Raft lab

Implement the official Raft lab in the matching MIT 6.5840 checkout. Use its supplied distributed tests for leader election, log replication, persistence, and unreliable-network behavior as assigned by the matching release. The local `exercise.go` is a small state-transition reading exercise, not the course lab.

## Upstream lab: MIT 6.5840 Raft

Use the official [MIT 6.5840 lab source](git://g.csail.mit.edu/6.5840-golabs-2026) and [course guide](https://pdos.csail.mit.edu/6.5840/), with the handout and tests from the same release. In the supplied 2026 checkout, run `make -C src raft1`; its Makefile runs the race-enabled `raft1` Go tests. Follow the selected release's course page and Makefile, and record the commit, command, and results in `EVIDENCE.md`. The local Go file is a state-transition exercise, not a substitute for the replicated state machine lab.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
