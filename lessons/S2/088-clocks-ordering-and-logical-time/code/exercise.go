// STUDENT PROBLEM SET — Lesson 088.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement a Lamport clock as a small class/struct with a `tick()` method (increment for a local event, return new value) and a `receive(remoteTimestamp)` method (implementing the max-plus-one rule from Learn).
// 2. Extend it with one explicit boundary case: Write a unit test confirming the basic increment behavior in isolation.
// 3. Rebuild the same behavior with one changed constraint: Simulate 3 independent "processes" (just 3 instances of your Lamport clock in one program, or 3 actual goroutines/threads communicating via channels/sockets) that occasionally send messages to each other carrying their current Lamport timestamp, with the receiver updating its own clock per the receive rule. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 088: Clocks, Ordering, and Logical Time
// Your task: implement Lamport's receive rule.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

type Clock struct {
	ID    int
	Value int
}

func (c *Clock) Tick() { c.Value++ }
func (c *Clock) Receive(remote int) {
    // TODO: implement the assigned behavior.
}

func main() {}
