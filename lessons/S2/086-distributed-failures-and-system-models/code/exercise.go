// STUDENT PROBLEM SET — Lesson 086.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself.
// 2. Extend it with one explicit boundary case: For each point, state whether the client, from a plain timeout alone, can distinguish it from any of the other points.
// 3. Rebuild the same behavior with one changed constraint: Using `tc netem` (networking Lesson 7's tool) or a similar mechanism, simulate a network partition between a client and server (block all traffic between them for a period, e.g. via a firewall rule or `tc` loss configuration set to 100%), and observe what your client-side code actually experiences — a timeout, a connection refused, or a hang, depending on yo. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 086: Distributed Failures and System Models
// Your task: classify failures using the system model from the lesson.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

type Outcome int

const (
	Success Outcome = iota
	RetryableFailure
	PermanentFailure
)

func classify(err error) Outcome {
    // TODO: implement the assigned behavior.
    return Outcome{}
}

func main() {}
