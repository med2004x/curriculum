# Lesson 077 — Problem Set

**Topic:** Threat Modeling and Trust Boundaries.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Draw the data-flow diagram for browser/client, API, database, object store, and operator.

2. Identify assets, actors, entry points, trust boundaries, and sensitive flows.

3. Apply STRIDE prompts to flows and record prioritized threats, controls, and residual risks.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/077-threat-modeling-and-trust-boundaries/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Threat model

Threat-model a small backend with assets, trust boundaries, attacker capabilities, threats, mitigations, and residual risk. Verify two controls in code or configuration.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
