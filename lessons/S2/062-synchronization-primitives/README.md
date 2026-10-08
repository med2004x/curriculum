# Lesson 062 — Synchronization Primitives in the Kernel and in Go

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Shared-memory coordination requires both a safety rule for accessing state and a communication or waiting mechanism for progress. Go provides high-level synchronization; the operating system and hardware provide lower-level scheduling and blocking support.

**The contract.** **Why a mutex alone isn't enough for producer/consumer coordination.** A mutex protects a critical section from concurrent access, but says nothing about *waiting for a condition* — e.g. a consumer thread that finds a shared buffer empty needs to wait until a producer adds something, without holding the lock the whole time (which would prevent the producer f.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty).

## 2. Mechanism and vocabulary

**Mechanism.** Go's memory model defines happens-before relationships. A mutex unlock synchronizes before a later lock that acquires it; channel communication and other documented primitives also establish ordering. Race-free Go programs have sequentially consistent behavior, while data races are correctness errors.

**Second point.** Linux futexes let synchronization state be managed mostly in user space when uncontended. The kernel provides wait/wake operations when a thread must block. This is an implementation mechanism, not a guarantee that every Go mutex operation makes a syscall or maps directly to one futex call.

For `sync.Cond`, hold its associated locker while checking or changing the predicate. `Wait` atomically unlocks while sleeping and reacquires before returning; re-check the predicate in a loop because another goroutine may change it before the waiter reacquires the lock.

**Worked example.** Use a `WaitGroup` to join workers and a channel to return results; use a mutex only for state genuinely shared by workers. State who owns each value and which operation establishes the required happens-before relation.

## 3. Worked case from the reference material

Worked example: Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty).

## 4. Verification, failure, and evidence

**Verification.** Draw the happens-before relation for a protected counter and a channel-owned counter. For a `Cond` example, identify predicate, lock, wait loop, state change, and signal/broadcast.

**Failure mode.** Read or write the counter outside its mutex and identify the race and missing synchronization edge. Separately omit a condition-predicate re-check after a wake and explain why a notification does not guarantee the predicate still holds.

**Completion criterion.** You can match a primitive to a coordination need, state its memory-ordering guarantee, and explain which work occurs in user space versus when the kernel may block or wake a thread.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
