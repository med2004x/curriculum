# Lesson 095 — Observability, SLOs, and Measurement

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Implement the three pillars of observability — logs, metrics, and traces — in a way that lets you answer "why is this specific request slow/failing" after the fact, rather than only knowing *that* something is wrong in aggregate.

**The contract.** Monitoring checks known signals against expected thresholds; observability supplies enough linked evidence to investigate questions that were not anticipated. An SLI defines an eligible event and a good event; an SLO sets a target over a window, and its error budget is the permitted bad-event share.

**Worked example.** Start with the smallest concrete case described by the source material: Implement structured logging for a real endpoint (JSON-formatted, with at minimum a request ID, timestamp, and severity level on every log line), and generate a request ID at the start of each request, propagating it through every subsequent log line for that request (via context, in Go's idiomatic style, or an equivalent mechanism in your language).

## 2. Mechanism and vocabulary

**Mechanism.** Correlated logs explain discrete events, metrics show aggregate behavior, and traces connect spans across service boundaries. A trace or request ID belongs in logs and trace context; metric labels should stay bounded so unique requests do not create unbounded time series.

**Second point.** **Logs: discrete, timestamped events, and why structure matters.** Unstructured log lines (`fmt.Println("processing request")`) are hard to search, filter, or aggregate at scale.

**Worked example.** Implement basic metrics: a counter for total requests, a counter for errors, and a histogram for request latency (using a metrics library like Prometheus's client library, or a simple hand-rolled equivalent for this exercise).

## 3. Worked case from the reference material

Worked example: Implement structured logging for a real endpoint (JSON-formatted, with at minimum a request ID, timestamp, and severity level on every log line), and generate a request ID at the start of each request, propagating it through every subsequent log line for that request (via context, in Go's idiomatic style, or an equivalent mechanism in your language).

## 4. Verification, failure, and evidence

**Verification.** Send a known mix of fast, slow, successful, and failed requests. Show one request ID across logs and spans, show request and error counts, inspect latency histogram buckets, and calculate the SLI from the generated events. State whether the sample is large enough to support any SLO conclusion.

**Failure mode.** Attempt the same step 4 diagnosis using *only* unstructured, non-request-correlated logs (revert to plain `fmt.Println`-style logging with no request ID propagation, mixing log lines from many concurrent requests together with no way to tell which lines belong to which request).

**Completion criterion.** You can follow a request across logs and spans, calculate the stated SLI from raw counts, compare it with its SLO, and explain what the telemetry cannot establish.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
