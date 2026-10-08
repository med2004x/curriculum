# Lesson 067 — Problem Set

**Topic:** Validation, Error Contracts, and API Design.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Pick a real feature from TARDOC or Mahall that involves an API endpoint you've already built (or plan to build) — for example, TARDOC's clinic subscription activation, or Mahall's product listing creation.

2. Extend it with one explicit boundary case: Before looking at (or without regard to) your existing implementation, write a plain-language requirements document answering: what resource is being acted on, what does the client send, what's returned on success, and what are at least 3 distinct ways this specific operation could fail, with a plain-English description of what each failure means.

3. Rebuild the same behavior with one changed constraint: For the same operation, explicitly list every field you'd include in the success response, and for each field, justify in one sentence why the client actually needs it (not "it was easy to include" or "it's just there internally") — if you can't justify a field this way, mark it as a candidate for removal from the contract.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/067-validation-errors-and-contracts/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Contract validation

Design a validation layer for a JSON API. Return machine-readable errors, distinguish client misuse from domain conflict, and test boundary conditions.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
