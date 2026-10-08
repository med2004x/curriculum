// STUDENT PROBLEM SET — Lesson 038.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Implement naive recursive Fibonacci and draw its call tree for a small argument.
// 2. Add memoization and bottom-up tabulation; explain which repeated computation is removed.
// 3. Derive and solve the recurrences for binary search and merge sort, stating theorem assumptions. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 038: Recursion, Dynamic Programming, and Recurrences
// Your task: use the recurrence and memoization to avoid repeated work.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

func Fib(n int, memo map[int]int) int {
    // TODO: implement the assigned behavior.
    return 0
}

func main() {
	fmt.Println(Fib(10, map[int]int{}))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
