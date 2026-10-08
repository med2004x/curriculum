// STUDENT PROBLEM SET — Lesson 030.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Implement a fixed-capacity ring buffer with `count`, head, and tail state protected by one mutex.
// 2. Coordinate empty/full transitions using two condition variables and predicate loops, or implement the same bounded flow with a buffered Go channel.
// 3. Run multiple producers and consumers over uniquely identified finite items; verify conservation and termination.
// 4. Build a separate two-lock example with opposite acquisition order, draw the wait-for graph, then enforce a global order. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 030: Synchronization, Channels, and Mutexes
// Your task: use a bounded worker pool and preserve result order.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main
import ("fmt"; "sync")
func runWorkers(jobs []int, workers int) []int {
    // TODO: implement the assigned behavior.
    return nil
}
func main(){fmt.Println(runWorkers([]int{1,2,3,4},2))}
