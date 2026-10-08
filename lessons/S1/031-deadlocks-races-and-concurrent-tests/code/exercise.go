// STUDENT PROBLEM SET — Lesson 031.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Write the two-thread/two-lock schedule and draw its wait-for cycle.
// 2. Change the implementation to acquire locks in one global order and explain which Coffman condition is broken.
// 3. Design a concurrency test with explicit barriers/channels to control the relevant schedule and a timeout only as a bounded failure guard. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 031: Deadlocks, Races, and Concurrent Tests
// Your task: replace the global gate with a stable lock ordering to permit safe concurrency.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main
import "sync"
var pairGate sync.Mutex
func withLocks(a,b *sync.Mutex, work func()) {
    // TODO: implement the assigned behavior.
}
func deliberatelyBroken(){ var a,b sync.Mutex; withLocks(&a,&b,func(){}) }
func main(){deliberatelyBroken()}

// Keep starter imports available while the student implementation is unfinished.
var _ sync.Mutex
