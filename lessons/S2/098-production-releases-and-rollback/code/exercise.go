// STUDENT PROBLEM SET — Lesson 098.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Create a pipeline that runs formatting checks, tests, static analysis, a build, and an artifact verification step.
// 2. Pin dependencies and toolchain inputs; build once and record the artifact digest so later promotion can verify artifact identity.
// 3. Rebuild the same behavior with one changed constraint: GitHub Actions, or a local script if CI isn't set up yet) that: checks out a specific Git commit, builds the artifact (a Go binary, a Docker image, or equivalent), and embeds the commit hash and build timestamp into the artifact itself (e.g. via a version endpoint the running application exposes, or embedded build-time constants). Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 098: Production Releases, CI/CD, Containers, and Rollback
// Your task: add build metadata, graceful shutdown, and a release-safe health check.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "net/http"

func main() {
    // TODO: implement the assigned behavior.
}

// Keep starter imports available while the student implementation is unfinished.
var _ = http.NewServeMux
