// STUDENT PROBLEM SET — Lesson 073.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.
// 2. Extend it with one explicit boundary case: Implement a webhook receiver (a separate small server) that verifies the incoming signature against its own copy of the shared secret before processing the payload, rejecting (with an appropriate 4xx status) any request with a missing or incorrect signature.
// 3. Rebuild the same behavior with one changed constraint: Test both a correctly signed request (accepted) and a tampered payload with a stale signature (rejected) to confirm signature verification is actually functioning, not just present in code but unused. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 073: Webhooks, Async Jobs, and Channels
// Your task: consume jobs, handle cancellation, and report completion exactly once.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "context"

func worker(ctx context.Context, jobs <-chan string, done chan<- string) {
    // TODO: implement the assigned behavior.
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ context.Context
