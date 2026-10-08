# Lesson 069 — Authentication, Authorization, and Threats Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a minimal token-based authentication scheme: a login endpoint that verifies credentials and issues a token (a signed JWT is a reasonable, realistic choice, using a Go library like `golang-jwt/jwt` or equivalent), and middleware that validates the token on subsequent requests, rejecting requests with missing or invalid tokens with `401 Unauthorized`.

2. Implement role-based authorization: attach a role claim to the issued token, and add middleware or handler-level logic that checks the required role for a given endpoint, rejecting requests from an authenticated-but-wrong-role user with `403 Forbidden` (distinct from the 401 in step 1 — test both cases and confirm the status codes differ correctly).

3. Implement resource-level (ownership) authorization on top of role-based checks: for an endpoint like `GET /clinics/{id}/records`, in addition to checking the user has an appropriate role, explicitly check that the requested clinic ID actually belongs to (or is otherwise legitimately accessible by) the authenticated user — and test the specific vulnerability.

## Run and check

```powershell
Set-Location lessons/S2/069-authentication-authorization-and-threats/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### AuthN/AuthZ

Implement authenticated access to two resources with distinct permissions. Add tests for anonymous, authenticated-but-forbidden, and authorized requests.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
