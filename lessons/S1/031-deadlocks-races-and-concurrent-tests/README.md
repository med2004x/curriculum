# Lesson 031 — Deadlocks, Races, and Concurrent Tests

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Reason about programs that stop making progress even though every thread is still running.

**The contract.** Deadlock is a set of participants blocked indefinitely while each waits for a resource held by another. Livelock is continued activity without progress; starvation is indefinite postponement of a participant while other work may proceed. A wait-for graph represents who waits for whom; a cycle indicates deadlock in the relevant single-instance resource model.

**Worked example.** Thread A holds lock X and requests Y; thread B holds Y and requests X. Draw edges from each waiting thread to the thread holding its requested lock. The cycle `A → B → A` explains why neither can proceed. Requiring all code to acquire X before Y prevents this particular cycle by breaking circular wait.

## 2. Mechanism and vocabulary

**Mechanism.** The Coffman conditions are mutual exclusion, hold-and-wait, no preemption, and circular wait. A prevention strategy must rule out at least one necessary condition; imposing a total order on lock acquisition rules out circular wait for those locks.

**Second point.** Tests should arrange the schedule-relevant events explicitly using channels, barriers, or other synchronization. Sleeps impose time delays, not ordering guarantees, so they are poor evidence that a specific interleaving occurred.

**Worked example.** Build a test that starts two workers, signals when each has acquired its first lock, and then allows each to request the second. This makes the intended deadlock schedule observable; use a bounded timeout as a test guard, not as proof of the precise cause.

## 3. Worked case from the reference material

Worked example: Construct a two-lock deadlock, draw its wait-for graph, then fix it with a global lock-order rule.

## 4. Verification, failure, and evidence

**Verification.** Record the exact command, input, and observable result for the smallest case.

**Failure mode.** Introduce a data race in a shared variable or a lock-order cycle, then use the race detector or wait-for graph to distinguish the two. A race is a safety violation; a deadlock is a liveness failure.

**Completion criterion.** You can explain the wait-for cycle, name the broken liveness property, select a prevention rule, and construct a test that proves the relevant events occurred.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
