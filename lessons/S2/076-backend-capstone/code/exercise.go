// STUDENT PROBLEM SET — Lesson 076.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Design and build a complete API for a real resource type from TARDOC or Mahall's actual domain (or a plausible extension of it) with this minimum required scope:
// 2. Extend it with one explicit boundary case: **Contract-first**: a complete OpenAPI spec (Lesson 3) written before implementation, covering all endpoints below, with reusable error-schema components (Lesson 3).
// 3. Rebuild the same behavior with one changed constraint: **Resource model** (Lesson 2): at least one top-level resource and one genuinely nested sub-resource, with a justified nesting decision. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 076: Integrated Project 3: Production-Grade Go Service
// Your task: implement the complete service path: validation, persistence, HTTP, observability.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "context"

type Service struct{}

func (Service) Create(ctx context.Context, name string) error {
    // TODO: implement the assigned behavior.
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ context.Context
