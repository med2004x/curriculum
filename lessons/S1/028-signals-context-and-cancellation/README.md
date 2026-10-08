# Lesson 028 — Signals, Context, and Cancellation

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Long-running work needs a way to stop when a caller disconnects, a deadline expires, or the process is asked to shut down. Signals request process-level action; Go contexts communicate cooperative cancellation through function calls.

**The contract.** A `context.Context` carries cancellation, deadlines, and request-scoped values. A callee that can block or perform repeated work should observe `ctx.Done()` and return `ctx.Err()` when cancellation is requested. A signal handler can begin graceful shutdown on `SIGTERM`, but cleanup must be bounded because a supervisor may later send `SIGKILL`.

**Worked example.** Build a worker that selects between receiving a job and `<-ctx.Done()`. A caller creates a context with a deadline, starts the worker, and waits for either a result or the deadline. Make every potentially blocking send, receive, and wait cancellation-aware.

## 2. Mechanism and vocabulary

**Mechanism.** Calling a context's cancel function closes its `Done` channel; it does not forcibly terminate a goroutine. Cancellation is cooperative. A deadline adds a time-based cancellation signal, but code must still check it and release resources.

**Second point.** Use `signal.NotifyContext` when a process-level signal should cancel a context. The signal initiates a shutdown path; it is not a replacement for context propagation inside the application.

`SIGTERM` can be handled or ignored; `SIGKILL` cannot be caught, blocked, or ignored. In Go, `signal.NotifyContext` can connect selected signals to cancellation, after which the application should stop accepting work, cancel requests, wait for bounded cleanup, and exit.

**Worked example.** Pass the same context through a handler, service, and worker. Give the context a short deadline; verify each layer returns when cancellation occurs and that the first meaningful error is preserved rather than replaced by a generic success value.

## 3. Worked case from the reference material

For the OS side, send `SIGTERM` to a test process and record that its graceful path begins. Separately, explain why sending `SIGKILL` offers no opportunity to run a handler. Do not use a production process for destructive signal experiments.

For the Go side, show the context creation, propagation, cancellation observation, and returned error as a four-step trace.

## 4. Verification, failure, and evidence

**Verification.** Record the deadline, the goroutine's cancellation path, the returned error, and whether all started work has completed. For signal handling, record the signal requested and the bounded cleanup outcome.

**Failure mode.** Create a goroutine that ignores cancellation.

**Completion criterion.** You can explain cooperative cancellation, propagate a deadline, and distinguish the process-level `SIGTERM` path from uncatchable `SIGKILL`.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
