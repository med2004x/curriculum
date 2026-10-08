# Lesson 093 — Requirements, Capacity, and Queueing Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.

2. Using the math-for-engineering queueing lesson's M/M/1 formulas, and your step 1 peak `λ` estimate, solve for the minimum service rate `μ` required to keep average time-in-system `W` under a target (e.g.

3. 2 seconds) — show your work using the actual formulas, not just a rough guess.

## Run and check

```powershell
Set-Location lessons/S2/093-requirements-capacity-and-queueing/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Capacity worksheet + simulator

Translate a set of product requirements into throughput, latency, storage, and capacity estimates. Validate the assumptions with a small simulator.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
