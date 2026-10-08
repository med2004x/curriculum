// STUDENT PROBLEM SET — Lesson 087.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Build a minimal RPC client/server (Go's built-in `net/rpc` package is a reasonable, quick way to get a working RPC mechanism without hand-rolling wire-format parsing, though you're welcome to build directly on networking Lesson 1's raw sockets/framing if you want the deeper exercise) with at least one method that takes an argument and returns a result.
// 2. Extend it with one explicit boundary case: Add a client-side timeout: if the server doesn't respond within a configured duration, the client call returns an error rather than blocking indefinitely.
// 3. Rebuild the same behavior with one changed constraint: Test it by adding an artificial delay to the server's handler (longer than your configured timeout) and confirming the client correctly times out rather than hanging. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 087: RPC, Serialization, and Backpressure
// Your task: make admission/backpressure behavior explicit and testable.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

type Request struct {
	ID      int
	Payload []byte
}

func send(queue chan<- Request, req Request) bool {
    // TODO: implement the assigned behavior.
    return false
}

func main() {}
