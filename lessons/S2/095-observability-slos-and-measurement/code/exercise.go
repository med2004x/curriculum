// STUDENT PROBLEM SET — Lesson 095.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement structured logging for a real endpoint (JSON-formatted, with at minimum a request ID, timestamp, and severity level on every log line), and generate a request ID at the start of each request, propagating it through every subsequent log line for that request (via context, in Go's idiomatic style, or an equivalent mechanism in your language).
// 2. Extend it with one explicit boundary case: Implement basic metrics: a counter for total requests, a counter for errors, and a histogram for request latency (using a metrics library like Prometheus's client library, or a simple hand-rolled equivalent for this exercise).
// 3. Rebuild the same behavior with one changed constraint: Expose them via a `/metrics` endpoint and confirm you can compute p50/p95/p99 latency from your histogram data. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 095: Observability, SLOs, and Measurement
// Your task: keep an SLI window and make the measurement semantics explicit.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "time"

type Window struct {
	Good  int
	Total int
	Start time.Time
}

func record(w *Window, ok bool) {
    // TODO: implement the assigned behavior.
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ = time.Time
