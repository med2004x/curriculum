# Lesson 074 — Graceful Shutdown and Lifecycles

## Threads
**Primary:** Go/Linux/Backend

**Supporting:** signals; context; draining work

**Math companion:** state transitions

**Time:** 3–4 hours. **Prerequisites:** Lesson 073.

**Time:** 3–4 hours. **Prerequisites:** Lesson 073, # Lesson 073 — Webhooks, Async Jobs, and Channels.

## What you will be able to do

1. Define startup, readiness, serving, draining, and stopped lifecycle states.
2. Stop admitting new work before waiting for in-flight requests.
3. Propagate bounded cancellation to background workers and close dependencies in order.
4. Explain behavior when the drain deadline expires or the process receives an uncatchable kill.

## Retrieve first (no notes)

1. Why are webhook deliveries commonly at-least-once?
2. How does an event ID help a receiver avoid duplicate effects?
3. Why verify HMAC over the raw request bytes?
4. Why should delivery retries be bounded?

## 1. Problem and model

**The contract.** A shutdown request moves the service through explicit lifecycle states: mark it unready, stop accepting new work, drain in-flight requests for a bounded interval, stop and join background workers, then close dependencies such as the database pool.

**Worked example.** Start with the smallest concrete case described by the source material: Build a server that receives a shutdown signal, stops accepting new work, finishes allowed in-flight work, and exits with a clean status.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** In Go, `signal.NotifyContext` can initiate cancellation on `SIGTERM`; `http.Server.Shutdown(ctx)` stops accepting new connections and waits for active handlers until the context expires. Background goroutines need explicit cancellation and joining; returning from `main` terminates the process even if goroutines remain.

**Second point.** Shutdown order follows dependencies: stop admission, drain handlers that use the database, stop workers, then close the database. If the deadline expires, cleanup remains bounded; `SIGKILL` cannot run application cleanup.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Modify the worker pipeline from Lesson 28 to stop when a deadline expires and to return the first meaningful error.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Build a server that receives a shutdown signal, stops accepting new work, finishes allowed in-flight work, and exits with a clean status.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record readiness/admission behavior, in-flight request completion, worker exit, resource-close order, and the outcome when the drain deadline expires.

**Failure mode.** Create a worker that ignores cancellation or a handler that exceeds the drain deadline. Confirm shutdown remains bounded and document which work may be abandoned.

**Completion criterion.** You can draw the lifecycle state machine, implement bounded graceful shutdown, and explain the forced-termination boundary.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the readiness state after shutdown begins.
2. Predict whether new requests should be admitted during draining.
3. Predict what `http.Server.Shutdown` waits for.
4. Predict what happens when an in-flight handler exceeds the deadline.
5. Predict when the database pool can safely close.
6. Predict what happens to a goroutine that ignores context cancellation.
7. Predict which signal cannot run cleanup.
8. State how the overall shutdown deadline bounds lifecycle work.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement a server that receives shutdown cancellation and stops accepting new work.
2. Add readiness transition, worker cancellation/join, and a bounded shutdown deadline.
3. Write a test that observes worker completion using synchronization rather than an arbitrary sleep.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Create a goroutine that ignores cancellation.

2. Write a test that catches the leaked work instead of sleeping arbitrarily.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change the grace period or add a background worker using the database. Explain which lifecycle transition and resource ordering must change.

## Retrieval questions

1. What are the admission, draining, and stopped lifecycle states?
2. Why stop accepting new work before closing the database?
3. What does `http.Server.Shutdown` wait for?
4. What happens when a worker ignores context cancellation?
5. How does a deadline bound cleanup?
6. Which signal cannot be caught for cleanup?

## Common mistakes

- Closing the database before in-flight handlers finish using it.
- Returning from `main` while background workers are still expected to complete.
- Assuming context cancellation forcibly terminates a goroutine.
- Allowing shutdown to wait forever.
- Expecting `SIGKILL` to run application cleanup.

## Precision checks

1. Draw startup/readiness/serving/draining/stopped transitions.
2. State when admission stops and when in-flight work is joined.
3. List shutdown order for workers, handlers, and database resources.
4. Explain timeout and forced-termination behavior.
5. Confirm no cleanup step can block beyond the overall deadline.
