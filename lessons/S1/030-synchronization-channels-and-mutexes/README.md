# Lesson 030 — Synchronization, Channels, and Mutexes

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** A mutex provides exclusive access, but it does not itself let a producer sleep until a buffer has space or a consumer sleep until data exists. Condition variables and channels provide coordination; lock ordering addresses a separate liveness risk.

**The contract.** For a capacity-`C` ring buffer, maintain `0 <= count <= C`; enqueue is allowed only when `count < C`, dequeue only when `count > 0`. In a condition-variable design, inspect and update this state while holding the mutex. Wait atomically releases the mutex while sleeping and reacquires it before returning.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty).

## 2. Mechanism and vocabulary

**Mechanism.** A condition variable notification is not the predicate itself. A waiter uses `while (count == 0) cond_wait(not_empty, &mutex)` (or `while (count == C)` for a producer), because wakeups can be spurious or another awakened thread may consume the condition first.

**Second point.** In Go, an unbuffered channel synchronizes sender and receiver handoff; a buffered channel of capacity `C` can provide bounded producer/consumer buffering. A mutex is often simpler for shared mutable state; channels are useful when ownership or work transfer is the central model. Neither choice removes the need to reason about blocking and shutdown.

Deadlock is a liveness failure. With two locks, one thread can hold A while waiting for B as another holds B while waiting for A. If all code acquires A before B, this particular circular-wait pattern is prevented; a consistent order does not prove the whole program free of every deadlock.

**Worked example.** Have multiple producer threads and multiple consumer threads running concurrently against the same buffer, and verify — after a run with a known total number of produced items — that the consumer(s) received exactly that many items, with no lost or duplicated items.

## 3. Worked case from the reference material

Worked example: Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty).

## 4. Verification, failure, and evidence

**Verification.** Use a known finite number of uniquely identifiable items. Confirm exactly that many items are consumed, each once, and the buffer count remains between 0 and capacity. Include empty/full boundary cases and ensure all producers/consumers terminate.

**Failure mode.** Remove the predicate re-check and explain how a waiter can proceed when the buffer is still empty/full. Separately create opposite lock acquisition order in a two-lock example and draw the resulting wait-for cycle.

**Completion criterion.** You can state the buffer invariant, justify predicate loops around condition waits, explain channel or mutex coordination, and show how a global lock order removes a specific circular wait.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
