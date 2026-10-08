# Lesson 076 — Problem Set

**Topic:** Integrated Project 3: Production-Grade Go Service.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Design and build a complete API for a real resource type from TARDOC or Mahall's actual domain (or a plausible extension of it) with this minimum required scope:

2. Extend it with one explicit boundary case: **Contract-first**: a complete OpenAPI spec (Lesson 3) written before implementation, covering all endpoints below, with reusable error-schema components (Lesson 3).

3. Rebuild the same behavior with one changed constraint: **Resource model** (Lesson 2): at least one top-level resource and one genuinely nested sub-resource, with a justified nesting decision.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/076-backend-capstone/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Production backend capstone

Build a production-shaped Go service combining routing, validation, persistence, authentication, idempotency, async work, observability, integration tests, and deployment notes.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
