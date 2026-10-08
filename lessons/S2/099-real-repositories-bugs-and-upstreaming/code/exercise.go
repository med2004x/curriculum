// STUDENT PROBLEM SET — Lesson 099.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before.
// 2. Extend it with one explicit boundary case: A project related to something covered elsewhere in this curriculum (a database engine, an HTTP framework, a CLI tool you use) is a reasonable, motivated choice.
// 3. Rebuild the same behavior with one changed constraint: Identify and trace the project's entry point(s): find `main()` (or equivalent), and trace the first few layers of what it calls, building an initial map of the major subsystems/modules it wires together. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 099: Real Repositories, Bug Reproduction, and Upstreaming
// Your task: minimize the failing behavior into a deterministic regression test.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "testing"

func TestRegression(t *testing.T) {
    // TODO: implement the assigned behavior.
}

// Keep starter imports available while the student implementation is unfinished.
var _ testing.T
