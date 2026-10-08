# Lesson 071 — Idempotency, Transactions, and Retries Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a `POST` endpoint with a durable side-effect record and a unique idempotency key scoped to the caller.

2. Return the stored response for a repeated same-key/same-fingerprint request; reject same-key/different-fingerprint reuse.

3. Add a concurrent reservation path and document how in-progress duplicates wait, receive a conflict, or receive the eventual stored result.

## Run and check

```powershell
Set-Location lessons/S2/071-idempotency-transactions-and-retries/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Idempotent write

Design an idempotent command endpoint with retries. Prove that duplicate delivery cannot create duplicate durable effects.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
