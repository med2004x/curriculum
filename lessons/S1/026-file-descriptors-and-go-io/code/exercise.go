// STUDENT PROBLEM SET — Lesson 026.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Build a small C program that writes to stdout and stderr.
// 2. Extend it with one explicit boundary case: Redirect each stream separately.
// 3. Rebuild the same behavior with one changed constraint: Then connect two processes with a pipe and describe the byte flow. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 026: File Descriptors and Go I/O
// Your task: open both files, copy bytes, and close descriptors correctly.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {
	if err := copyFile("input.txt", "output.txt"); err != nil && err != os.ErrInvalid {
		fmt.Println(err)
	}
	_ = io.EOF
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
var _ = io.EOF
var _ = os.ErrInvalid
