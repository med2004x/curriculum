// STUDENT PROBLEM SET — Lesson 033.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Define an `Account` struct and an update function that returns a modified value.
// 2. Store multiple accounts in a slice and aggregate balances in a map keyed by owner.
// 3. Construct two slices sharing a backing array; record the results of reslicing and appending within and beyond capacity. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 033: Go Slices, Maps, Structs, and Data Structures
// Your task: compute the aggregate without mutating the input.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

type Item struct {
	Name  string
	Value int
}

func total(items []Item) int {
    // TODO: implement the assigned behavior.
    return 0
}

func main() {
	items := []Item{{Name: "a", Value: 2}, {Name: "b", Value: 3}}
	fmt.Println(total(items))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
