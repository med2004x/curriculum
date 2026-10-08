# Lesson 031 — Deadlocks, Races, and Concurrent Tests

## Threads
**Primary:** OS/Go

**Supporting:** liveness; deadlock; race detection; testing

**Math companion:** graphs and wait-for relations

**Time:** 3–4 hours. **Prerequisites:** Lesson 030.

**Time:** 3–4 hours. **Prerequisites:** Lesson 030.

## What you will be able to do

1. Distinguish deadlock, livelock, and starvation as different liveness failures.
2. Construct and analyze a wait-for graph for a two-lock deadlock.
3. Explain the four Coffman conditions and identify which condition a chosen prevention strategy breaks.
4. Design concurrency tests that coordinate on events and completion rather than arbitrary sleeps.

## Retrieve first (no notes)

1. What invariant does a bounded buffer maintain?
2. Why must a condition-variable predicate be re-checked after waking?
3. What is a consistent lock-order rule?
4. How can item identities reveal loss or duplication that a final count misses?

## 1. Problem and model

**What problem this solves.** Reason about programs that stop making progress even though every thread is still running.

**The contract.** Deadlock is a set of participants blocked indefinitely while each waits for a resource held by another. Livelock is continued activity without progress; starvation is indefinite postponement of a participant while other work may proceed. A wait-for graph represents who waits for whom; a cycle indicates deadlock in the relevant single-instance resource model.

**Worked example.** Thread A holds lock X and requests Y; thread B holds Y and requests X. Draw edges from each waiting thread to the thread holding its requested lock. The cycle `A → B → A` explains why neither can proceed. Requiring all code to acquire X before Y prevents this particular cycle by breaking circular wait.

Self-check: distinguish a wait-for cycle from a mere long-running operation; a hang alone does not prove deadlock.

## 2. Mechanism and vocabulary

**Mechanism.** The Coffman conditions are mutual exclusion, hold-and-wait, no preemption, and circular wait. A prevention strategy must rule out at least one necessary condition; imposing a total order on lock acquisition rules out circular wait for those locks.

**Second point.** Tests should arrange the schedule-relevant events explicitly using channels, barriers, or other synchronization. Sleeps impose time delays, not ordering guarantees, so they are poor evidence that a specific interleaving occurred.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Build a test that starts two workers, signals when each has acquired its first lock, and then allows each to request the second. This makes the intended deadlock schedule observable; use a bounded timeout as a test guard, not as proof of the precise cause.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Construct a two-lock deadlock, draw its wait-for graph, then fix it with a global lock-order rule.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record the exact command, input, and observable result for the smallest case.

**Failure mode.** Introduce a data race in a shared variable or a lock-order cycle, then use the race detector or wait-for graph to distinguish the two. A race is a safety violation; a deadlock is a liveness failure.

**Completion criterion.** You can explain the wait-for cycle, name the broken liveness property, select a prevention rule, and construct a test that proves the relevant events occurred.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Draw the wait-for graph when A holds X and requests Y while B holds Y and requests X.
2. Predict which Coffman conditions hold in the example.
3. Predict which condition is removed by a global lock order.
4. Predict the difference between a deadlock and a livelock in observable progress.
5. Predict how starvation can occur without a wait-for cycle.
6. Predict why a sleep-based test may pass without exercising its intended interleaving.
7. Predict what synchronization event would prove each worker acquired its first lock.
8. Identify what a race detector can report and what it cannot prove about liveness.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Write the two-thread/two-lock schedule and draw its wait-for cycle.

2. Change the implementation to acquire locks in one global order and explain which Coffman condition is broken.

3. Design a concurrency test with explicit barriers/channels to control the relevant schedule and a timeout only as a bounded failure guard.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Introduce opposite lock ordering and identify the exact hold-and-wait cycle.

2. Replace timing assumptions with event notifications that establish each required ordering.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one condition: add a third lock to the ordering rule or introduce a third worker. Update the wait-for graph and explain whether the same prevention argument still applies.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Define deadlock, livelock, and starvation.
2. What does each edge in a wait-for graph represent?
3. Name the four Coffman conditions.
4. Which condition does consistent global lock ordering prevent?
5. Why are arbitrary sleeps not reliable synchronization in a test?
6. How can a bounded timeout help a concurrency test without proving deadlock by itself?

## Evidence to record

Create `lesson031/EVIDENCE.md` with exactly this structure:

```text
# Lesson 031 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Calling any slow operation a deadlock without evidence of a wait cycle.
- Confusing a race (safety) with a deadlock (liveness).
- Using sleeps as though they guarantee an execution order.
- Claiming a lock-order rule prevents deadlocks involving unrelated resources not included in that order.
- Treating a test timeout alone as a diagnosis of the failure cause.

## Precision checks

1. Draw a wait-for graph and identify a directed cycle.
2. Classify the incident as safety, deadlock, livelock, or starvation.
3. Name the Coffman condition your prevention rule breaks.
4. Identify the synchronization events that establish the tested interleaving.
5. State what a timeout reveals and what it does not establish.

Concurrency diagnosis requires evidence about both state and progress: a race report and a wait-for cycle answer different questions.
## University revision bridge

For the LIG track, connect deadlock prevention to resource-allocation graphs and formal liveness reasoning, then test the concrete locking discipline with controlled synchronization.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can draw and explain a two-lock wait cycle, distinguish deadlock from livelock, starvation, and data races, justify a prevention rule using the Coffman conditions, and design a concurrency test whose ordering comes from synchronization rather than sleeps.
