// STUDENT PROBLEM SET — Lesson 028.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Implement a worker that receives a context and exits when `ctx.Done()` closes.
// 2. Add a deadline and propagate the context through at least two function layers.
// 3. Connect `SIGTERM` to cancellation with `signal.NotifyContext` and give cleanup a bounded wait. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 028: Signals, Context, and Cancellation
// Your task: stop promptly when ctx is cancelled.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {
	err := work(context.Background())
	fmt.Println(err)
}

// Keep starter imports available while the student implementation is unfinished.
var _ context.Context
var _ = fmt.Println
var _ = time.After
