# Lesson 100 — Graduation: Design, Build, Measure, Explain Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:

2. **Contract-first design** (API engineering Lessons 1-3): a written requirements document and OpenAPI spec before implementation.

3. **Correct networking fundamentals** (networking Lessons 1-4): proper framing, correct HTTP semantics, connection reuse.

## Run and check

```powershell
Set-Location lessons/S2/100-graduation-capstone/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Graduation capstone

Design, build, measure, secure, test, document, and defend a substantial backend or systems project. Treat the result as a technical review, not a demo.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
