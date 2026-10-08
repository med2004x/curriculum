// STUDENT PROBLEM SET — Lesson 094.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Draw a call chain and assign one end-to-end deadline across its hops.
// 2. Add bounded exponential retries with jitter only for a transient/idempotent operation.
// 3. Add a bounded queue or admission limit and define the overload response. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 094: Timeouts, Retries, Backpressure, and Recovery
// Your task: apply bounded retries, per-attempt deadlines, and backoff without retrying unsafe operations.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"context"
	"net/http"
	"time"
)

func request(ctx context.Context, client *http.Client, url string) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ context.Context
var _ http.Client
var _ = time.Sleep
