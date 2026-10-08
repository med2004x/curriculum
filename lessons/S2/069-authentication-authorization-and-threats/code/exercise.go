// STUDENT PROBLEM SET — Lesson 069.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement a minimal token-based authentication scheme: a login endpoint that verifies credentials and issues a token (a signed JWT is a reasonable, realistic choice, using a Go library like `golang-jwt/jwt` or equivalent), and middleware that validates the token on subsequent requests, rejecting requests with missing or invalid tokens with `401 Unauthorized`.
// 2. Extend it with one explicit boundary case: Implement role-based authorization: attach a role claim to the issued token, and add middleware or handler-level logic that checks the required role for a given endpoint, rejecting requests from an authenticated-but-wrong-role user with `403 Forbidden` (distinct from the 401 in step 1 — test both cases and confirm the status codes differ correctly).
// 3. Rebuild the same behavior with one changed constraint: Implement resource-level (ownership) authorization on top of role-based checks: for an endpoint like `GET /clinics/{id}/records`, in addition to checking the user has an appropriate role, explicitly check that the requested clinic ID actually belongs to (or is otherwise legitimately accessible by) the authenticated user — and test the specific vulnerability. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 069: Authentication, Authorization, and Threats
// Your task: separate authentication from authorization and enforce least privilege.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

type User struct {
	ID   int
	Role string
}

func canAccess(u User, resourceOwner int) bool {
    // TODO: implement the assigned behavior.
    return false
}

func main() { fmt.Println(canAccess(User{ID: 1, Role: "user"}, 1)) }

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
