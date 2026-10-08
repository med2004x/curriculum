// STUDENT PROBLEM SET — Lesson 072.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement `GET /resource/{id}` returning an `ETag` header (a simple version integer, or a hash of the resource's serialized content, is fine) alongside the resource body.
// 2. Extend it with one explicit boundary case: Implement `PUT /resource/{id}` requiring an `If-Match` header matching the resource's current ETag.
// 3. Rebuild the same behavior with one changed constraint: On a match, apply the update and return a new ETag; on a mismatch, return `412 Precondition Failed` with a clear error body explaining the resource has changed. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 072: Concurrent Updates and Consistency
// Your task: reject stale versions and preserve the account invariant.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "errors"

type Account struct {
	Balance int
	Version int
}

func update(a *Account, expectedVersion, delta int) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ = errors.New
