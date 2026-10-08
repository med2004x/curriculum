// STUDENT PROBLEM SET — Lesson 029.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Write a Go program with several goroutines incrementing a shared counter without synchronization.
// 2. Add the race detector and explain the reported conflicting read/write operations.
// 3. Protect the counter with a mutex and make every read of the protected state use the same mutex. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 029: Threads, Goroutines, and Data Races
// Your task: protect the shared update and wait for all workers.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main
import ("fmt"; "sync")
func incrementMany(total *int, workers, increments int) {
    // TODO: implement the assigned behavior.
}
func main(){ total:=0; incrementMany(&total,4,1000); fmt.Println(total) }
