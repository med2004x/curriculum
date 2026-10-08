// STUDENT PROBLEM SET — Lesson 032.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Split a small program into an application package and a lower-level parsing or data package.
// 2. Export only the identifiers required by the caller and keep the package dependency graph acyclic.
// 3. Declare the module path and required Go version, inspect dependency metadata, and explain how module-wide tests enumerate packages. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 032: Packages, Modules, and Reproducible Builds
// Your task: move reusable logic into a package boundary and make the build reproducible.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

func makeLabel(name string, n int) string {
    // TODO: implement the assigned behavior.
    return ""
}

func main() {
	fmt.Println(makeLabel("lesson", 32))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Sprintf
