# Lesson 062 — Synchronization Primitives in the Kernel and in Go

## Threads
**Primary:** OS/Go

**Supporting:** mutexes; condvars; channels; scheduling interaction

**Math companion:** logic of interleavings

**Time:** 3–4 hours. **Prerequisites:** Lesson 061.

**Time:** 3–4 hours. **Prerequisites:** Lesson 061.

## What you will be able to do

1. Explain how synchronization establishes happens-before and makes protected data visible.
2. Choose among channels, `sync.Mutex`, `sync.RWMutex`, `sync.Cond`, and `sync.WaitGroup` for a stated coordination need.
3. Explain Linux futexes as user-space synchronization with kernel-assisted blocking under contention.
4. State the invariant and lock rules for condition-variable code.

## Retrieve first (no notes)

1. What does EOF on a pipe indicate?
2. Why must processes close unused pipe ends?
3. How can a full pipe create backpressure?
4. What does a process-local file descriptor refer to?

## 1. Problem and model

**What problem this solves.** Shared-memory coordination requires both a safety rule for accessing state and a communication or waiting mechanism for progress. Go provides high-level synchronization; the operating system and hardware provide lower-level scheduling and blocking support.

**The contract.** **Why a mutex alone isn't enough for producer/consumer coordination.** A mutex protects a critical section from concurrent access, but says nothing about *waiting for a condition* — e.g. a consumer thread that finds a shared buffer empty needs to wait until a producer adds something, without holding the lock the whole time (which would prevent the producer f.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty).

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Go's memory model defines happens-before relationships. A mutex unlock synchronizes before a later lock that acquires it; channel communication and other documented primitives also establish ordering. Race-free Go programs have sequentially consistent behavior, while data races are correctness errors.

**Second point.** Linux futexes let synchronization state be managed mostly in user space when uncontended. The kernel provides wait/wake operations when a thread must block. This is an implementation mechanism, not a guarantee that every Go mutex operation makes a syscall or maps directly to one futex call.

For `sync.Cond`, hold its associated locker while checking or changing the predicate. `Wait` atomically unlocks while sleeping and reacquires before returning; re-check the predicate in a loop because another goroutine may change it before the waiter reacquires the lock.

**Worked example.** Use a `WaitGroup` to join workers and a channel to return results; use a mutex only for state genuinely shared by workers. State who owns each value and which operation establishes the required happens-before relation.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Draw the happens-before relation for a protected counter and a channel-owned counter. For a `Cond` example, identify predicate, lock, wait loop, state change, and signal/broadcast.

**Failure mode.** Read or write the counter outside its mutex and identify the race and missing synchronization edge. Separately omit a condition-predicate re-check after a wake and explain why a notification does not guarantee the predicate still holds.

**Completion criterion.** You can match a primitive to a coordination need, state its memory-ordering guarantee, and explain which work occurs in user space versus when the kernel may block or wake a thread.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict which accesses to a shared counter must hold the mutex.
2. Predict the happens-before edge from unlock to a later successful lock.
3. Predict when sends block on buffered and unbuffered channels.
4. Predict what `WaitGroup.Wait` guarantees about task completion.
5. Predict why a `Cond.Wait` caller checks its predicate in a loop.
6. Predict whether an uncontended Go mutex necessarily enters the kernel.
7. Choose a primitive for ownership transfer versus shared mutable state.
8. State one synchronization guarantee from documentation rather than inferring implementation details.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement a protected counter and ensure every access follows the mutex protocol.

2. Rewrite the counter as channel-owned state and identify the send/receive synchronization relation.

3. Add a `sync.Cond` predicate wait and explain lock release/reacquisition plus the required loop.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Remove synchronization from one counter access and identify the race and missing happens-before edge.

2. Replace a one-time condition check with a predicate loop around `Cond.Wait`.

3. Explain why futex behavior is an implementation detail and why code relies on Go's documented synchronization contract.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one primitive: replace the mutex with an atomic counter or a channel-owned state. Explain which invariant the design protects and what synchronization edge it provides.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is a happens-before relationship?
2. What ordering does `sync.Mutex` provide?
3. Why must a condition predicate be checked in a loop?
4. How does an unbuffered channel synchronize sender and receiver?
5. What role does a Linux futex play under contention?
6. Why need not every Go mutex operation invoke a futex syscall?

## Evidence to record

Create `lesson062/EVIDENCE.md` with exactly this structure:

```text
# Lesson 062 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a mutex as protecting accesses that bypass it.
- Treating `WaitGroup` as a substitute for data synchronization.
- Checking a condition only once after `Cond.Wait` returns.
- Assuming channels and mutexes provide identical ownership semantics.
- Treating futex behavior as a Go API guarantee.

## Precision checks

1. Draw happens-before edges for mutex, channel, and `WaitGroup` examples.
2. State the protected invariant and identify every access to it.
3. Explain the `Cond` locker, predicate, and wait loop.
4. Distinguish documented Go behavior from Linux implementation details.
5. Choose a primitive based on ownership, shared state, and waiting requirements.

Synchronization is both exclusion and ordering; a correct design names the shared invariant and the edges that make observations valid.
## University revision bridge

For the LIG track, connect kernel-assisted blocking to synchronization primitives and Go's happens-before model, keeping interface guarantees separate from runtime implementation.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 062 — Synchronization Primitives in the Kernel and in Go without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your bounded buffer implementation is correct under concurrent producer/consumer load across multiple runs, you've reliably reproduced a two-lock deadlock and diagnosed its exact structure using a debugger rather than just observing a hang, and you've fixed it with lock ordering and confirmed — across many repeated runs — that the fix is structural, not just. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
