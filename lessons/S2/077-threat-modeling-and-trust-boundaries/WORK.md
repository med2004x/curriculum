# Lesson 077 — Threat Modeling and Trust Boundaries Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Draw the data-flow diagram for browser/client, API, database, object store, and operator.

2. Identify assets, actors, entry points, trust boundaries, and sensitive flows.

3. Apply STRIDE prompts to flows and record prioritized threats, controls, and residual risks.

## Run and check

```powershell
Set-Location lessons/S2/077-threat-modeling-and-trust-boundaries/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Threat model

Threat-model a small backend with assets, trust boundaries, attacker capabilities, threats, mitigations, and residual risk. Verify two controls in code or configuration.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
