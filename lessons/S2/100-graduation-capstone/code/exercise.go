// STUDENT PROBLEM SET — Lesson 100.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:
// 2. Extend it with one explicit boundary case: **Contract-first design** (API engineering Lessons 1-3): a written requirements document and OpenAPI spec before implementation.
// 3. Rebuild the same behavior with one changed constraint: **Correct networking fundamentals** (networking Lessons 1-4): proper framing, correct HTTP semantics, connection reuse. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 100: Graduation: Design, Build, Measure, Explain
// Your task: your graduation system belongs here. State the contract before implementation.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "context"

type Application struct{}

func (Application) Run(ctx context.Context) error {
    // TODO: implement the assigned behavior.
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ context.Context
