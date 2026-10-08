# Lesson 028 — Signals, Context, and Cancellation

## Threads
**Primary:** Linux/Go

**Supporting:** signals; deadlines; cancellation propagation

**Math companion:** probability and expected waiting

**Time:** 3–4 hours. **Prerequisites:** Lesson 027.

**Time:** 3–4 hours. **Prerequisites:** Lesson 027.

## What you will be able to do

1. Explain the difference between an operating-system signal and cooperative Go context cancellation.
2. Propagate a context deadline through a worker call chain and stop work when `Done` is closed.
3. Handle `SIGTERM` for bounded cleanup while recognizing that `SIGKILL` cannot be caught.
4. Test cancellation without relying on arbitrary sleeps or leaked goroutines.

## Retrieve first (no notes)

1. What does a system call request from the kernel?
2. What information can `/proc/<pid>/status` expose?
3. How does Go conventionally report a failed operation?
4. Why should an error not be confused with a valid zero result?

## 1. Problem and model

**What problem this solves.** Long-running work needs a way to stop when a caller disconnects, a deadline expires, or the process is asked to shut down. Signals request process-level action; Go contexts communicate cooperative cancellation through function calls.

**The contract.** A `context.Context` carries cancellation, deadlines, and request-scoped values. A callee that can block or perform repeated work should observe `ctx.Done()` and return `ctx.Err()` when cancellation is requested. A signal handler can begin graceful shutdown on `SIGTERM`, but cleanup must be bounded because a supervisor may later send `SIGKILL`.

**Worked example.** Build a worker that selects between receiving a job and `<-ctx.Done()`. A caller creates a context with a deadline, starts the worker, and waits for either a result or the deadline. Make every potentially blocking send, receive, and wait cancellation-aware.

Self-check: does every blocked worker operation have a route to observe cancellation, or can a goroutine remain stuck after its caller returns?

## 2. Mechanism and vocabulary

**Mechanism.** Calling a context's cancel function closes its `Done` channel; it does not forcibly terminate a goroutine. Cancellation is cooperative. A deadline adds a time-based cancellation signal, but code must still check it and release resources.

**Second point.** Use `signal.NotifyContext` when a process-level signal should cancel a context. The signal initiates a shutdown path; it is not a replacement for context propagation inside the application.

`SIGTERM` can be handled or ignored; `SIGKILL` cannot be caught, blocked, or ignored. In Go, `signal.NotifyContext` can connect selected signals to cancellation, after which the application should stop accepting work, cancel requests, wait for bounded cleanup, and exit.

**Worked example.** Pass the same context through a handler, service, and worker. Give the context a short deadline; verify each layer returns when cancellation occurs and that the first meaningful error is preserved rather than replaced by a generic success value.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

For the OS side, send `SIGTERM` to a test process and record that its graceful path begins. Separately, explain why sending `SIGKILL` offers no opportunity to run a handler. Do not use a production process for destructive signal experiments.

For the Go side, show the context creation, propagation, cancellation observation, and returned error as a four-step trace.

Self-check: cancellation closes the context channel; it does not forcibly stop code that never checks it.

## 4. Verification, failure, and evidence

**Verification.** Record the deadline, the goroutine's cancellation path, the returned error, and whether all started work has completed. For signal handling, record the signal requested and the bounded cleanup outcome.

**Failure mode.** Create a goroutine that ignores cancellation.

**Completion criterion.** You can explain cooperative cancellation, propagate a deadline, and distinguish the process-level `SIGTERM` path from uncatchable `SIGKILL`.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict when a context created with a deadline closes `Done`.
2. Predict the value returned by `ctx.Err()` after deadline expiration.
3. Predict what happens to a goroutine that never checks its context.
4. Predict whether cancellation forcibly interrupts an arbitrary blocking operation.
5. Predict which signal can enter an application cleanup path.
6. Predict the effect of a worker blocked on a channel send after its receiver returns.
7. Predict how context cancellation should propagate through nested functions.
8. Record which goroutines the caller must still join before returning.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement a worker that receives a context and exits when `ctx.Done()` closes.

2. Add a deadline and propagate the context through at least two function layers.

3. Connect `SIGTERM` to cancellation with `signal.NotifyContext` and give cleanup a bounded wait.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Create a worker that ignores cancellation and identify why the caller cannot stop it by calling `cancel` alone.

2. Add a cancellation-aware blocking operation and a completion signal; make the test wait on synchronization rather than an arbitrary delay.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one condition: replace explicit cancellation with a deadline, or cancel while the worker is blocked on a channel send rather than doing computation. Explain how the worker observes the request and how completion is confirmed.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Is Go context cancellation cooperative or preemptive?
2. What does closing `ctx.Done()` communicate?
3. How does a deadline differ from an explicit cancel call?
4. What must a worker do to respond to context cancellation?
5. How does `SIGTERM` differ from `SIGKILL`?
6. Why should shutdown cleanup have a deadline?

## Evidence to record

Create `lesson028/EVIDENCE.md` with exactly this structure:

```text
# Lesson 028 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Assuming `cancel()` forcibly terminates goroutines.
- Failing to pass the context into a blocking operation.
- Returning from a caller while child goroutines remain blocked.
- Using `SIGKILL` while expecting application cleanup.
- Making cancellation tests depend on a guessed sleep duration.

## Precision checks

1. Draw cancellation propagation through caller, service, and worker.
2. Explain which goroutine observes each `Done` channel.
3. Identify every blocking point that must be cancellation-aware.
4. State how started work is joined before the caller returns.
5. Explain the signal used to begin graceful shutdown and the limit imposed by `SIGKILL`.

Cancellation is a request, not a kill switch: correctness requires both propagation and a path for each worker to return.
## University revision bridge

For the LIG track, connect signal delivery and process shutdown to OS semantics, then treat Go contexts as a cooperative application-level cancellation protocol.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can propagate a context deadline through a worker pipeline, show every blocking worker can exit on cancellation, join started work, and explain the distinction between catchable `SIGTERM` and uncatchable `SIGKILL`.
