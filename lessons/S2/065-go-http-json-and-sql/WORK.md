# Lesson 065 — Go HTTP, JSON, and SQL Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Define typed create-request and response structs and a bounded JSON decoder.

2. Validate the request and call a context-aware repository method using parameters.

3. Map success, malformed input, invalid fields, not-found, and internal error to distinct outcomes.

## Run and check

```powershell
Set-Location lessons/S2/065-go-http-json-and-sql/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### JSON + SQL service

Build a Go HTTP endpoint that validates JSON input, performs one SQL operation, and returns a stable JSON contract with error cases tested.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
