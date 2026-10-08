# Lesson 067 — Validation, Error Contracts, and API Design Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Pick a real feature from TARDOC or Mahall that involves an API endpoint you've already built (or plan to build) — for example, TARDOC's clinic subscription activation, or Mahall's product listing creation.

2. Before looking at (or without regard to) your existing implementation, write a plain-language requirements document answering: what resource is being acted on, what does the client send, what's returned on success, and what are at least 3 distinct ways this specific operation could fail, with a plain-English description of what each failure means.

3. For the same operation, explicitly list every field you'd include in the success response, and for each field, justify in one sentence why the client actually needs it (not "it was easy to include" or "it's just there internally") — if you can't justify a field this way, mark it as a candidate for removal from the contract.

## Run and check

```powershell
Set-Location lessons/S2/067-validation-errors-and-contracts/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Contract validation

Design a validation layer for a JSON API. Return machine-readable errors, distinguish client misuse from domain conflict, and test boundary conditions.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
