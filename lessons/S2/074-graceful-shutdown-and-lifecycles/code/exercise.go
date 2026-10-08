// STUDENT PROBLEM SET — Lesson 074.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Implement a server that receives shutdown cancellation and stops accepting new work.
// 2. Add readiness transition, worker cancellation/join, and a bounded shutdown deadline.
// 3. Write a test that observes worker completion using synchronization rather than an arbitrary sleep. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 074: Graceful Shutdown and Lifecycles
// Your task: coordinate listener, shutdown, deadlines, and returned errors.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"context"
	"net/http"
)

func runServer(ctx context.Context, srv *http.Server) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ context.Context
var _ http.Server
