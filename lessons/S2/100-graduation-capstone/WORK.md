# Lesson 100 — Problem Set

**Topic:** Graduation: Design, Build, Measure, Explain.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:

2. Extend it with one explicit boundary case: **Contract-first design** (API engineering Lessons 1-3): a written requirements document and OpenAPI spec before implementation.

3. Rebuild the same behavior with one changed constraint: **Correct networking fundamentals** (networking Lessons 1-4): proper framing, correct HTTP semantics, connection reuse.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/100-graduation-capstone/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Graduation capstone

Design, build, measure, secure, test, document, and defend a substantial backend or systems project. Treat the result as a technical review, not a demo.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
