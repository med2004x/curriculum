// STUDENT PROBLEM SET — Lesson 027.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Run a sleeping process and inspect its PID, parent PID, open descriptors, and command line.
// 2. Extend it with one explicit boundary case: Send a signal and record the state transition.
// 3. Rebuild the same behavior with one changed constraint: Trace a program that opens and reads a file. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 027: Syscalls, Errors, and Explicit Control Flow
// Your task: parse, validate, and wrap failures with useful context.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parseCount(input string) (int, error) {
    // TODO: implement the assigned behavior.
    return 0, nil
}

func main() {
	n, err := parseCount("42")
	fmt.Println(n, err)
}

// Keep starter imports available while the student implementation is unfinished.
var _ = errors.New
var _ = fmt.Println
var _ = strconv.Atoi
