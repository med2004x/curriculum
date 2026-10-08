# Lesson 072 — Concurrent Updates and Consistency Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement `GET /resource/{id}` returning an `ETag` header (a simple version integer, or a hash of the resource's serialized content, is fine) alongside the resource body.

2. Implement `PUT /resource/{id}` requiring an `If-Match` header matching the resource's current ETag.

3. On a match, apply the update and return a new ETag; on a mismatch, return `412 Precondition Failed` with a clear error body explaining the resource has changed.

## Run and check

```powershell
Set-Location lessons/S2/072-concurrent-updates-and-consistency/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Concurrent inventory

Implement an inventory update path and test concurrent decrements. Show the race or lost-update failure before selecting a consistency mechanism.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
