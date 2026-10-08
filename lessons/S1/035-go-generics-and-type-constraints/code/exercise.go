// STUDENT PROBLEM SET — Lesson 035.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement `contains` for strings and integers first.
// 2. Test empty, singleton, duplicate, found, and missing-value cases for both types.
// 3. Write `Contains[T comparable]` and explain why the constraint is necessary for equality. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 035: Generics and Type Constraints
// Your task: implement the empty-input contract and generic comparison.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

func Max[T ~int | ~float64](values []T) (T, bool) {
    // TODO: implement the assigned behavior.
    var zero T
    return zero, false
}

func main() {
	v, ok := Max([]int{3, 7, 2})
	fmt.Println(v, ok)
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
