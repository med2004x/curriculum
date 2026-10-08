// STUDENT PROBLEM SET — Lesson 066.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Audit one of your own existing API endpoints (TARDOC or Mahall) and check whether its HTTP method usage matches this lesson's safety/idempotency conventions — specifically, confirm no GET endpoint has side effects, and check whether your DELETE endpoints behave idempotently (calling delete on an already-deleted resource — what does it currently return, and i.
// 2. Extend it with one explicit boundary case: For the same endpoint (or a small set of related endpoints), produce a status-code table: every distinct outcome (success and each failure mode) mapped to a specific status code, and confirm you're not reusing the same status code for two semantically different outcomes, or using different codes for what's actually the same kind of outcome across similar end.
// 3. Rebuild the same behavior with one changed constraint: Deliberately construct a bad example: an endpoint that returns `200 OK` even when the operation actually failed (with the actual error only discoverable by inspecting the response body), and rewrite it to return an appropriate 4xx or 5xx status code instead, with the error detail in the body. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 066: Routing, CRUD, and HTTP Semantics
// Your task: implement the POST contract and status semantics. fmt.Fprintln(w, "implement the required behavior")
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"fmt"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
    // TODO: implement the assigned behavior.
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", createProduct)
	_ = http.ListenAndServe(":8080", mux)
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Fprintln
var _ = http.ResponseWriter
