# Lesson 065 — Problem Set

**Topic:** Go HTTP, JSON, and SQL.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define typed create-request and response structs and a bounded JSON decoder.

2. Validate the request and call a context-aware repository method using parameters.

3. Map success, malformed input, invalid fields, not-found, and internal error to distinct outcomes.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/065-go-http-json-and-sql/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### JSON + SQL service

Build a Go HTTP endpoint that validates JSON input, performs one SQL operation, and returns a stable JSON contract with error cases tested.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
