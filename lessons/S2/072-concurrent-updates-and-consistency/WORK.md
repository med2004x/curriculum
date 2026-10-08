# Lesson 072 — Problem Set

**Topic:** Concurrent Updates and Consistency.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement `GET /resource/{id}` returning an `ETag` header (a simple version integer, or a hash of the resource's serialized content, is fine) alongside the resource body.

2. Extend it with one explicit boundary case: Implement `PUT /resource/{id}` requiring an `If-Match` header matching the resource's current ETag.

3. Rebuild the same behavior with one changed constraint: On a match, apply the update and return a new ETag; on a mismatch, return `412 Precondition Failed` with a clear error body explaining the resource has changed.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/072-concurrent-updates-and-consistency/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Concurrent inventory

Implement an inventory update path and test concurrent decrements. Show the race or lost-update failure before selecting a consistency mechanism.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
