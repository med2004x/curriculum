# Lesson 030 — Synchronization, Channels, and Mutexes

## Threads
**Primary:** OS/Go

**Supporting:** mutual exclusion; message passing; happens-before

**Math companion:** logic of interleavings

**Time:** 3–4 hours. **Prerequisites:** Lesson 029.

**Time:** 3–4 hours. **Prerequisites:** Lesson 029.

## What you will be able to do

1. State the bounded-buffer invariant and protect it with a mutex.
2. Explain why condition-variable waits must re-check their predicate in a loop.
3. Implement producer/consumer coordination with either a condition variable or a Go channel.
4. Recognize a two-lock deadlock and explain how a consistent global lock order prevents that cycle.

## Retrieve first (no notes)

1. What two accesses make a data race?
2. Why should every access to a shared invariant follow the same mutex protocol?
3. What does a race detector report?
4. Why can a raced counter occasionally appear to have the expected value?

## 1. Problem and model

**What problem this solves.** A mutex provides exclusive access, but it does not itself let a producer sleep until a buffer has space or a consumer sleep until data exists. Condition variables and channels provide coordination; lock ordering addresses a separate liveness risk.

**The contract.** For a capacity-`C` ring buffer, maintain `0 <= count <= C`; enqueue is allowed only when `count < C`, dequeue only when `count > 0`. In a condition-variable design, inspect and update this state while holding the mutex. Wait atomically releases the mutex while sleeping and reacquires it before returning.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty).

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** A condition variable notification is not the predicate itself. A waiter uses `while (count == 0) cond_wait(not_empty, &mutex)` (or `while (count == C)` for a producer), because wakeups can be spurious or another awakened thread may consume the condition first.

**Second point.** In Go, an unbuffered channel synchronizes sender and receiver handoff; a buffered channel of capacity `C` can provide bounded producer/consumer buffering. A mutex is often simpler for shared mutable state; channels are useful when ownership or work transfer is the central model. Neither choice removes the need to reason about blocking and shutdown.

Deadlock is a liveness failure. With two locks, one thread can hold A while waiting for B as another holds B while waiting for A. If all code acquires A before B, this particular circular-wait pattern is prevented; a consistent order does not prove the whole program free of every deadlock.

**Worked example.** Have multiple producer threads and multiple consumer threads running concurrently against the same buffer, and verify — after a run with a known total number of produced items — that the consumer(s) received exactly that many items, with no lost or duplicated items.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Use a known finite number of uniquely identifiable items. Confirm exactly that many items are consumed, each once, and the buffer count remains between 0 and capacity. Include empty/full boundary cases and ensure all producers/consumers terminate.

**Failure mode.** Remove the predicate re-check and explain how a waiter can proceed when the buffer is still empty/full. Separately create opposite lock acquisition order in a two-lock example and draw the resulting wait-for cycle.

**Completion criterion.** You can state the buffer invariant, justify predicate loops around condition waits, explain channel or mutex coordination, and show how a global lock order removes a specific circular wait.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the buffer count after enqueue and dequeue at capacity 1.
2. Predict what a consumer does when the buffer is empty.
3. Predict why condition waits use `while`, not `if`.
4. Predict what can occur when multiple consumers are awakened for one item.
5. Predict the lock ownership and wait edges in the two-lock deadlock.
6. Predict how a global lock order changes those wait edges.
7. Predict whether a buffered Go channel blocks when empty or full.
8. State how the finite workload proves no items were lost or duplicated.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement a fixed-capacity ring buffer with `count`, head, and tail state protected by one mutex.

2. Coordinate empty/full transitions using two condition variables and predicate loops, or implement the same bounded flow with a buffered Go channel.

3. Run multiple producers and consumers over uniquely identified finite items; verify conservation and termination.

4. Build a separate two-lock example with opposite acquisition order, draw the wait-for graph, then enforce a global order.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Change a condition-variable predicate wait from `while` to `if` and explain the broken assumption.

2. Construct the two-lock cycle and identify the thread holding and requesting each lock.

3. Correct the wait predicate and impose one lock order; state precisely which failure each correction prevents.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one condition: use a Go buffered channel instead of a condition-variable queue, or change buffer capacity from one to two. Explain how blocking and the invariant are represented in the new design.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. State the ring-buffer count invariant.
2. Why must a condition predicate be checked in a loop while holding its mutex?
3. What does a Go buffered channel do when full or empty?
4. Draw the wait-for cycle in a two-lock deadlock.
5. How does consistent lock ordering prevent that cycle?
6. What evidence shows that a finite set of items was neither lost nor duplicated?

## Evidence to record

Create `lesson030/EVIDENCE.md` with exactly this structure:

```text
# Lesson 030 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Waiting on a condition while holding a lock in a way that prevents the condition from changing.
- Using `if` instead of re-checking the predicate in a loop.
- Treating a notification as proof that the condition is true.
- Acquiring multiple locks in inconsistent orders.
- Checking only a final total when duplicate and missing items could cancel numerically.

## Precision checks

1. State the count invariant and the empty/full predicates.
2. Explain how a condition wait releases and later reacquires the mutex.
3. Explain why the predicate is re-checked after waking.
4. Give the global ordering rule for the two locks.
5. Show how item identities establish conservation, not just equal totals.

Synchronization establishes safety and coordination, but liveness still depends on wait conditions, lock ordering, and a shutdown path.
## University revision bridge

For the LIG track, connect mutual exclusion and condition synchronization to the operating-system model, then compare those primitives with Go's channel-based communication.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain and implement a bounded producer/consumer buffer, justify predicate loops, verify item conservation, and show why consistent lock ordering prevents a specified two-lock deadlock.
