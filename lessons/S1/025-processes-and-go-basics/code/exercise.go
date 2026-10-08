// STUDENT PROBLEM SET — Lesson 025.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Write a C program (or use `fork()` via a small wrapper if working in Go, which exposes `os/exec` instead — pick whichever language makes the fork/exec split easiest to observe directly; C is more instructive here since it exposes both syscalls explicitly) that calls `fork()`, and in the child, calls `execve` (or `execvp`) to run `/bin/ls`.
// 2. Extend it with one explicit boundary case: In the parent, call `wait()` to retrieve the child's exit status and print it.
// 3. Rebuild the same behavior with one changed constraint: Confirm the parent process's PID is unchanged before and after, while the child's PID (printed from inside the child, before `exec`) differs from the parent's. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 025: Processes and Your First Go Program
// Your task: implement the input transformation from the lesson.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"fmt"
	"strings"
)

func normalize(input string) string {
    // TODO: implement the assigned behavior.
    return ""
}

func main() {
	fmt.Println(normalize("  process  "))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
var _ = strings.TrimSpace
