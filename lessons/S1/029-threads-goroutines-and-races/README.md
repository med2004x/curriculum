# Lesson 029 — Threads, Goroutines, and Data Races

## Threads
**Primary:** OS/Go

**Supporting:** threads vs goroutines; race model

**Math companion:** probability refresher

**Time:** 3–4 hours. **Prerequisites:** Lesson 028.

**Time:** 3–4 hours. **Prerequisites:** Lesson 028.

## What you will be able to do

1. Explain how threads and goroutines can access shared process memory.
2. Define a data race as conflicting concurrent accesses to the same memory location, with at least one write and no synchronization ordering them.
3. Use Go's race detector or ThreadSanitizer to detect a deliberately unsynchronized access.
4. Fix the race with a mutex or another synchronization strategy and state the protected invariant.

## Retrieve first (no notes)

1. What does a Go context's `Done` channel communicate?
2. Does calling `cancel` forcibly terminate a goroutine?
3. How can a deadline propagate through nested functions?
4. Why should a caller wait for its workers to finish?

## 1. Problem and model

**What problem this solves.** Understand threads as units of execution sharing an address space, deliberately produce a data race, and fix it correctly with synchronization — building the concrete intuition that makes "just add a mutex" an informed decision rather than a reflex.

**The contract.** Threads in one process share the process address space and many resources; goroutines are Go's lightweight concurrent execution units scheduled by the Go runtime. A data race is a correctness defect, not merely a way to get a variable final value that differs between runs.

**Worked example.** In Go, start several goroutines that update one shared counter without a mutex. Increment is a read-modify-write sequence, so concurrent operations can overlap. The exact final count is not a reliable diagnostic by itself: a run may appear correct despite the race.

Self-check: identify the shared memory location and both accesses that conflict; do not infer race freedom from one expected-looking final value.

## 2. Mechanism and vocabulary

**Mechanism.** A mutex establishes exclusive access to a critical section. Its lock/unlock operations also establish synchronization ordering for protected memory. Every read and write to a shared invariant must follow the same synchronization rule.

**Second point.** A race detector reports executed conflicting accesses that lack synchronization. It is dynamic evidence, not a proof that untested paths are race-free. The detector does not guarantee that a raced program will show a wrong counter value on every run.

For C, an unsynchronized data race has undefined behavior under the language memory model. In Go, programs with races are not guaranteed to behave as intended; use the race detector and synchronization rather than relying on observed output.

**Worked example.** Protect the shared Go counter with `sync.Mutex`: lock, increment, unlock. Ensure no observer reads that counter without taking the same lock. The expected total is then a useful assertion, while the race detector checks the access discipline.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Write a program (C with pthreads, or Go with goroutines — pick whichever your track context makes more natural, though C is more instructive here since Go's race detector will make the next steps almost too easy) with multiple threads/goroutines all incrementing a shared, unprotected counter many times (e.g.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Run the unsynchronized version with Go's race detector and retain the report identifying the conflicting accesses. Then run the synchronized version and confirm the invariant and expected total. Do not require the unsynchronized final value to vary or be incorrect in every execution.

**Failure mode.** Keep writes protected by a mutex but read the same state from a reporting goroutine without locking or another synchronization mechanism. The unprotected read remains a race even though the write path takes a lock.

**Completion criterion.** You can reproduce a detector report, explain the conflicting accesses, fix every access path to the shared invariant, and distinguish dynamic detector evidence from a proof over all executions.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Identify the memory location accessed by two concurrent goroutines.
2. Predict which accesses conflict and whether one writes.
3. Predict what synchronization relation is missing.
4. Predict whether one successful-looking counter value proves race freedom.
5. Predict which accesses a mutex must protect, including reads.
6. Predict what kind of stack report the race detector will produce.
7. Predict what invariant the synchronized version should preserve.
8. State one limitation of a dynamic race detector.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Write a Go program with several goroutines incrementing a shared counter without synchronization.

2. Add the race detector and explain the reported conflicting read/write operations.

3. Protect the counter with a mutex and make every read of the protected state use the same mutex.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take the fixed, mutex-protected version from step 3, and introduce a *second* shared variable that's incremented in the same critical section but *read* elsewhere in the program without holding the same lock (e.g. a separate reporting goroutine/thread that reads the counter periodically without locking).

2. Run under the race detector (Go) or ThreadSanitizer (C) again and confirm it flags this new, partial-protection race even though the *write* path is correctly locked — the lesson being that protecting a write doesn't automatically protect all reads;

3. Introduce a shared counter race.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one synchronization strategy: replace the mutex-protected counter with a channel-owned counter or an atomic counter. Explain which operations are indivisible and which invariant the design preserves.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What access pattern constitutes a data race?
2. Why is incrementing a shared counter not one indivisible operation?
3. What synchronization guarantee does a mutex provide?
4. Why must a read use the synchronization protecting writes?
5. What evidence does Go's race detector provide?
6. Why does one correct-looking output not prove race freedom?

## Evidence to record

Create `lesson029/EVIDENCE.md` with exactly this structure:

```text
# Lesson 029 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a data race as merely a nondeterministic final value.
- Expecting every raced execution to produce a visibly wrong counter.
- Protecting writes while leaving reads unprotected.
- Assuming one race-detector-clean run proves all possible paths are race-free.
- Confusing a mutex-protected critical section with an invariant protected everywhere.

## Precision checks

1. Name the conflicting memory location and the read/write pair.
2. Explain why a mutex establishes ordering as well as exclusion.
3. Show that each access to the invariant follows the same lock protocol.
4. Distinguish race-detector evidence from proof for all executions.
5. Explain why the unsynchronized counter's exact output is not guaranteed.

The invariant is protected only if every conflicting access follows the synchronization protocol.
## University revision bridge

For the LIG track, connect shared address spaces and synchronization to OS threads, and use Go's runtime tooling to inspect concrete race reports.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can identify the conflicting accesses, obtain and interpret a race-detector report, synchronize all access paths to the shared invariant, and explain why a correct-looking raced output is not evidence of correctness.
