# Lesson 066 — Routing, CRUD, and HTTP Semantics Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Audit one of your own existing API endpoints (TARDOC or Mahall) and check whether its HTTP method usage matches this lesson's safety/idempotency conventions — specifically, confirm no GET endpoint has side effects, and check whether your DELETE endpoints behave idempotently (calling delete on an already-deleted resource — what does it currently return, and i.

2. For the same endpoint (or a small set of related endpoints), produce a status-code table: every distinct outcome (success and each failure mode) mapped to a specific status code, and confirm you're not reusing the same status code for two semantically different outcomes, or using different codes for what's actually the same kind of outcome across similar end.

3. Deliberately construct a bad example: an endpoint that returns `200 OK` even when the operation actually failed (with the actual error only discoverable by inspecting the response body), and rewrite it to return an appropriate 4xx or 5xx status code instead, with the error detail in the body.

## Run and check

```powershell
Set-Location lessons/S2/066-router-crud-and-http-semantics/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### CRUD API

Implement a complete resource API with routing, status semantics, pagination or filtering, and deterministic error responses. Test the public contract rather than internals.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
