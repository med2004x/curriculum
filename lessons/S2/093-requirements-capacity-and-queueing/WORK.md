# Lesson 093 — Problem Set

**Topic:** Requirements, Capacity, and Queueing.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.

2. Extend it with one explicit boundary case: Using the math-for-engineering queueing lesson's M/M/1 formulas, and your step 1 peak `λ` estimate, solve for the minimum service rate `μ` required to keep average time-in-system `W` under a target (e.g.

3. Rebuild the same behavior with one changed constraint: 2 seconds) — show your work using the actual formulas, not just a rough guess.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/093-requirements-capacity-and-queueing/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Capacity worksheet + simulator

Translate a set of product requirements into throughput, latency, storage, and capacity estimates. Validate the assumptions with a small simulator.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
