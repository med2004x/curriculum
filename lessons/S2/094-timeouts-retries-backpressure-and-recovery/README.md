# Lesson 094 — Timeouts, Retries, Backpressure, and Recovery

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Set timeouts and retry policies deliberately, based on the actual failure modes and cost tradeoffs involved, rather than copying default values — and understand cascading failure, the specific danger poorly-tuned timeouts/retries can cause across a system of interdependent services.

**The contract.** An operation has one end-to-end deadline shared across its dependency calls; nested services should not each restart a fresh full timeout. Retries consume the remaining deadline and a bounded budget. Retry only operations that are idempotent or carry a stable idempotency key, and distinguish transient failures from permanent errors.

**Worked example.** Start with the smallest concrete case described by the source material: For a real call chain in TARDOC or Mahall (or a plausible one — e.g. an API endpoint that calls a database, which itself might be under load), identify what timeout is currently configured (or would be, by a common default) at each hop, and reason about whether that value seems deliberately chosen or just a framework default.

## 2. Mechanism and vocabulary

**Mechanism.** Exponential backoff increases delay between attempts; jitter spreads synchronized clients; a retry budget caps extra work. A timeout bounds waiting but does not prove the remote operation did not complete.

**Second point.** Backpressure prevents unbounded work accumulation: bounded queues block or reject producers, admission control limits in-flight requests, and load shedding rejects lower-priority work. Recovery should ramp traffic gradually so a recovering dependency is not immediately overwhelmed.

**Worked example.** Simulate A → B → C with a shared deadline. Let C slow down and compare bounded queue/rejection plus backoff against immediate unbounded retries. Track in-flight work, attempt count, queue depth, and final success/error.

## 3. Worked case from the reference material

Worked example: For a real call chain in TARDOC or Mahall (or a plausible one — e.g. an API endpoint that calls a database, which itself might be under load), identify what timeout is currently configured (or would be, by a common default) at each hop, and reason about whether that value seems deliberately chosen or just a framework default.

## 4. Verification, failure, and evidence

**Verification.** Under the same downstream slowdown, record deadline propagation, attempts, backoff delays, queue depth, in-flight work, and rejected requests for bounded versus naive retry behavior.

**Failure mode.** Configure a timeout shorter than healthy downstream latency and measure false timeouts/retries. Then set a realistic budget and explain the latency-versus-resource trade-off.

**Completion criterion.** You can propagate deadlines, bound retries/queues, reject overload deliberately, and explain how the combined policy limits cascading work and supports controlled recovery.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
