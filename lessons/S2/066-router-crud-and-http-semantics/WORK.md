# Lesson 066 — Problem Set

**Topic:** Routing, CRUD, and HTTP Semantics.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Audit one of your own existing API endpoints (TARDOC or Mahall) and check whether its HTTP method usage matches this lesson's safety/idempotency conventions — specifically, confirm no GET endpoint has side effects, and check whether your DELETE endpoints behave idempotently (calling delete on an already-deleted resource — what does it currently return, and i.

2. Extend it with one explicit boundary case: For the same endpoint (or a small set of related endpoints), produce a status-code table: every distinct outcome (success and each failure mode) mapped to a specific status code, and confirm you're not reusing the same status code for two semantically different outcomes, or using different codes for what's actually the same kind of outcome across similar end.

3. Rebuild the same behavior with one changed constraint: Deliberately construct a bad example: an endpoint that returns `200 OK` even when the operation actually failed (with the actual error only discoverable by inspecting the response body), and rewrite it to return an appropriate 4xx or 5xx status code instead, with the error detail in the body.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/066-router-crud-and-http-semantics/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### CRUD API

Implement a complete resource API with routing, status semantics, pagination or filtering, and deterministic error responses. Test the public contract rather than internals.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
