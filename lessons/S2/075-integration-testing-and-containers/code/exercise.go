// STUDENT PROBLEM SET — Lesson 075.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.
// 2. Extend it with one explicit boundary case: Containerize one backend service and document which state is ephemeral and which state is persistent.
// 3. Replace the service container and verify persistent database state remains only when the database volume is retained. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 075: Integration Testing and Containers
// Your task: replace the placeholder with a real integration-style assertion.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
    // TODO: implement the assigned behavior.
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ = http.MethodGet
var _ = httptest.NewRequest
var _ testing.T
