# Lesson 095 — Observability, SLOs, and Measurement Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement structured logging for a real endpoint (JSON-formatted, with at minimum a request ID, timestamp, and severity level on every log line), and generate a request ID at the start of each request, propagating it through every subsequent log line for that request (via context, in Go's idiomatic style, or an equivalent mechanism in your language).

2. Implement basic metrics: a counter for total requests, a counter for errors, and a histogram for request latency (using a metrics library like Prometheus's client library, or a simple hand-rolled equivalent for this exercise).

3. Expose them via a `/metrics` endpoint and confirm you can compute p50/p95/p99 latency from your histogram data.

## Run and check

```powershell
Set-Location lessons/S2/095-observability-slos-and-measurement/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Observability + SLO

Instrument a service with request rate, errors, latency quantiles, and a defined SLO. Trigger a failure and use the evidence to explain what happened.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
