# Lesson 071 — Problem Set

**Topic:** Idempotency, Transactions, and Retries.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a `POST` endpoint with a durable side-effect record and a unique idempotency key scoped to the caller.

2. Return the stored response for a repeated same-key/same-fingerprint request; reject same-key/different-fingerprint reuse.

3. Add a concurrent reservation path and document how in-progress duplicates wait, receive a conflict, or receive the eventual stored result.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/071-idempotency-transactions-and-retries/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Idempotent write

Design an idempotent command endpoint with retries. Prove that duplicate delivery cannot create duplicate durable effects.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
