# Lesson 093 — Requirements, Capacity, and Queueing

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Translate a vague performance goal ("it should be fast," "it should handle our users") into concrete numbers — expected request rate, latency targets, data volume — using the mathematics track's queueing theory to reason about what capacity is actually required, rather than guessing.

**The contract.** **Why "it should be fast" isn't a specification.** Without a number, you can't know whether a design meets the requirement, can't compare two design options objectively, and can't tell whether a load test result is good or bad.

**Worked example.** Start with the smallest concrete case described by the source material: For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.

## 2. Mechanism and vocabulary

**Mechanism.** For an M/M/1 queue, assume Poisson arrivals at rate `λ`, exponential service at rate `μ`, one server, and `ρ=λ/μ<1`. Mean time in system is `W=1/(μ-λ)` and mean queue wait is `Wq=λ/(μ(μ-λ))`. These are model outputs, not p95/p99 guarantees.

**Second point.** As `ρ` approaches 1, the denominator `μ-λ` shrinks and mean delay rises sharply. Real services have non-Poisson bursts, multiple servers, variable service times, caches, and downstream limits; validate estimates with representative load tests and observed latency distributions.

**Worked example.** Convert 500 clinics × 20 requests per 8-hour day into an average rate, state a peak-to-average factor, and calculate peak `λ`. Given a mean response target `W`, solve `μ > λ + 1/W` for the M/M/1 model; state that this is a simplified mean-latency estimate, not a tail-latency SLO.

## 3. Worked case from the reference material

Worked example: For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.

## 4. Verification, failure, and evidence

**Verification.** Show the arithmetic from business volume to average/peak `λ`, compute `ρ`, `W`, and `Wq` for the stated `μ`, and compare with a higher load. State model assumptions and identify what must be measured to validate them.

**Failure mode.** Take your step 2 capacity plan (provisioned for your best-estimate peak `λ`) and recompute `W` if actual peak load turns out to be 50% higher than estimated (a realistic margin of estimation error) — using the same `μ` you provisioned for.

**Completion criterion.** You can derive a transparent demand estimate, apply the M/M/1 formulas within their assumptions, and explain why headroom and measurement are necessary near saturation.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
