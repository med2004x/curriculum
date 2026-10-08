# Lesson 029 — Threads, Goroutines, and Data Races

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand threads as units of execution sharing an address space, deliberately produce a data race, and fix it correctly with synchronization — building the concrete intuition that makes "just add a mutex" an informed decision rather than a reflex.

**The contract.** Threads in one process share the process address space and many resources; goroutines are Go's lightweight concurrent execution units scheduled by the Go runtime. A data race is a correctness defect, not merely a way to get a variable final value that differs between runs.

**Worked example.** In Go, start several goroutines that update one shared counter without a mutex. Increment is a read-modify-write sequence, so concurrent operations can overlap. The exact final count is not a reliable diagnostic by itself: a run may appear correct despite the race.

## 2. Mechanism and vocabulary

**Mechanism.** A mutex establishes exclusive access to a critical section. Its lock/unlock operations also establish synchronization ordering for protected memory. Every read and write to a shared invariant must follow the same synchronization rule.

**Second point.** A race detector reports executed conflicting accesses that lack synchronization. It is dynamic evidence, not a proof that untested paths are race-free. The detector does not guarantee that a raced program will show a wrong counter value on every run.

For C, an unsynchronized data race has undefined behavior under the language memory model. In Go, programs with races are not guaranteed to behave as intended; use the race detector and synchronization rather than relying on observed output.

**Worked example.** Protect the shared Go counter with `sync.Mutex`: lock, increment, unlock. Ensure no observer reads that counter without taking the same lock. The expected total is then a useful assertion, while the race detector checks the access discipline.

## 3. Worked case from the reference material

Worked example: Write a program (C with pthreads, or Go with goroutines — pick whichever your track context makes more natural, though C is more instructive here since Go's race detector will make the next steps almost too easy) with multiple threads/goroutines all incrementing a shared, unprotected counter many times (e.g.

## 4. Verification, failure, and evidence

**Verification.** Run the unsynchronized version with Go's race detector and retain the report identifying the conflicting accesses. Then run the synchronized version and confirm the invariant and expected total. Do not require the unsynchronized final value to vary or be incorrect in every execution.

**Failure mode.** Keep writes protected by a mutex but read the same state from a reporting goroutine without locking or another synchronization mechanism. The unprotected read remains a race even though the write path takes a lock.

**Completion criterion.** You can reproduce a detector report, explain the conflicting accesses, fix every access path to the shared invariant, and distinguish dynamic detector evidence from a proof over all executions.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
