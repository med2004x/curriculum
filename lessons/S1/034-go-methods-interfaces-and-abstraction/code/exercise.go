// STUDENT PROBLEM SET — Lesson 034.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Define a small `Payer` interface and a value-receiver implementation.
// 2. Add a second implementation and make one caller depend only on `Payer`.
// 3. Change the receiver to a pointer receiver and predict interface satisfaction for `T` and `*T`. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 034: Go Methods, Interfaces, and Abstraction
// Your task: define the representation contract.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

type Stringer interface {
	String() string
}

type User struct {
	Name string
}

func (u User) String() string {
    // TODO: implement the assigned behavior.
}

func describe(s Stringer) string {
	return s.String()
}

func main() {
	fmt.Println(describe(User{Name: "Ada"}))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
