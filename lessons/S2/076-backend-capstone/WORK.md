# Lesson 076 — Integrated Project 3: Production-Grade Go Service Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Design and build a complete API for a real resource type from TARDOC or Mahall's actual domain (or a plausible extension of it) with this minimum required scope:

2. **Contract-first**: a complete OpenAPI spec (Lesson 3) written before implementation, covering all endpoints below, with reusable error-schema components (Lesson 3).

3. **Resource model** (Lesson 2): at least one top-level resource and one genuinely nested sub-resource, with a justified nesting decision.

## Run and check

```powershell
Set-Location lessons/S2/076-backend-capstone/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Production backend capstone

Build a production-shaped Go service combining routing, validation, persistence, authentication, idempotency, async work, observability, integration tests, and deployment notes.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
