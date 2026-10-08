# Lesson 074 — Graceful Shutdown and Lifecycles

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**The contract.** A shutdown request moves the service through explicit lifecycle states: mark it unready, stop accepting new work, drain in-flight requests for a bounded interval, stop and join background workers, then close dependencies such as the database pool.

**Worked example.** Start with the smallest concrete case described by the source material: Build a server that receives a shutdown signal, stops accepting new work, finishes allowed in-flight work, and exits with a clean status.

## 2. Mechanism and vocabulary

**Mechanism.** In Go, `signal.NotifyContext` can initiate cancellation on `SIGTERM`; `http.Server.Shutdown(ctx)` stops accepting new connections and waits for active handlers until the context expires. Background goroutines need explicit cancellation and joining; returning from `main` terminates the process even if goroutines remain.

**Second point.** Shutdown order follows dependencies: stop admission, drain handlers that use the database, stop workers, then close the database. If the deadline expires, cleanup remains bounded; `SIGKILL` cannot run application cleanup.

**Worked example.** Modify the worker pipeline from Lesson 28 to stop when a deadline expires and to return the first meaningful error.

## 3. Worked case from the reference material

Worked example: Build a server that receives a shutdown signal, stops accepting new work, finishes allowed in-flight work, and exits with a clean status.

## 4. Verification, failure, and evidence

**Verification.** Record readiness/admission behavior, in-flight request completion, worker exit, resource-close order, and the outcome when the drain deadline expires.

**Failure mode.** Create a worker that ignores cancellation or a handler that exceeds the drain deadline. Confirm shutdown remains bounded and document which work may be abandoned.

**Completion criterion.** You can draw the lifecycle state machine, implement bounded graceful shutdown, and explain the forced-termination boundary.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
