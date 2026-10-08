# Lesson 093 — Requirements, Capacity, and Queueing

## Threads
**Primary:** System Design/Math

**Supporting:** load model; bottlenecks; capacity estimates

**Math companion:** queueing theory

**Time:** 3–4 hours. **Prerequisites:** Lesson 092.

**Time:** 3–4 hours. **Prerequisites:** Lesson 092.

## What you will be able to do

1. Translate business demand into average and peak request-rate assumptions.
2. Apply the M/M/1 utilization and response-time formulas under their stated assumptions.
3. Explain why latency rises nonlinearly as utilization approaches one.
4. State the limits of a queueing model and what measurements are needed for production capacity planning.

## Retrieve first (no notes)

1. What happens when a consumer crashes after processing but before acknowledgment?
2. How can deduplication and the effect be committed atomically in one database?
3. What ordering guarantee applies within a queue partition?
4. What scope can an exactly-once claim reasonably cover?

## 1. Problem and model

**What problem this solves.** Translate a vague performance goal ("it should be fast," "it should handle our users") into concrete numbers — expected request rate, latency targets, data volume — using the mathematics track's queueing theory to reason about what capacity is actually required, rather than guessing.

**The contract.** **Why "it should be fast" isn't a specification.** Without a number, you can't know whether a design meets the requirement, can't compare two design options objectively, and can't tell whether a load test result is good or bad.

**Worked example.** Start with the smallest concrete case described by the source material: For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** For an M/M/1 queue, assume Poisson arrivals at rate `λ`, exponential service at rate `μ`, one server, and `ρ=λ/μ<1`. Mean time in system is `W=1/(μ-λ)` and mean queue wait is `Wq=λ/(μ(μ-λ))`. These are model outputs, not p95/p99 guarantees.

**Second point.** As `ρ` approaches 1, the denominator `μ-λ` shrinks and mean delay rises sharply. Real services have non-Poisson bursts, multiple servers, variable service times, caches, and downstream limits; validate estimates with representative load tests and observed latency distributions.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Convert 500 clinics × 20 requests per 8-hour day into an average rate, state a peak-to-average factor, and calculate peak `λ`. Given a mean response target `W`, solve `μ > λ + 1/W` for the M/M/1 model; state that this is a simplified mean-latency estimate, not a tail-latency SLO.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Show the arithmetic from business volume to average/peak `λ`, compute `ρ`, `W`, and `Wq` for the stated `μ`, and compare with a higher load. State model assumptions and identify what must be measured to validate them.

**Failure mode.** Take your step 2 capacity plan (provisioned for your best-estimate peak `λ`) and recompute `W` if actual peak load turns out to be 50% higher than estimated (a realistic margin of estimation error) — using the same `μ` you provisioned for.

**Completion criterion.** You can derive a transparent demand estimate, apply the M/M/1 formulas within their assumptions, and explain why headroom and measurement are necessary near saturation.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Calculate average request rate from daily volume and service hours.
2. Predict peak rate under the chosen multiplier.
3. Calculate utilization `ρ` for stated `λ` and `μ`.
4. Predict mean system time as `ρ` approaches 1.
5. Recompute time under a 50% increase in `λ` and state if stability remains.
6. Identify which model assumptions may fail for bursty traffic.
7. Distinguish mean latency from p95/p99 latency.
8. State which load-test observation would validate or challenge the estimate.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: For a real or plausible scenario based on TARDOC (e.g. "500 clinics, each triggering roughly 20 transcription requests per business day, concentrated in an 8-hour window"), derive an estimated average requests-per-second, and then apply a reasonable peak-to-average ratio (state your assumption explicitly) to estimate peak requests-per-second.

2. Extend it with one explicit boundary case: Using the math-for-engineering queueing lesson's M/M/1 formulas, and your step 1 peak `λ` estimate, solve for the minimum service rate `μ` required to keep average time-in-system `W` under a target (e.g.

3. Rebuild the same behavior with one changed constraint: 2 seconds) — show your work using the actual formulas, not just a rough guess.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take your step 2 capacity plan (provisioned for your best-estimate peak `λ`) and recompute `W` if actual peak load turns out to be 50% higher than estimated (a realistic margin of estimation error) — using the same `μ` you provisioned for.

2. Report the new `ρ` and the resulting `W`, and confirm, using the queueing formula directly, that this modest underestimate produces a disproportionately large latency degradation, not a modest one — directly demonstrating why capacity planning needs deliberate headroom, not just point estimates, and why "close to my estimate" and "close to my target latency".

3. Push your simulation from step 3 to `ρ=0.99` (e.g.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 093 — Requirements, Capacity, and Queueing: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 093 — Requirements, Capacity, and Queueing?
2. What two concrete practices are required in the assigned material for # 093 — Requirements, Capacity, and Queueing?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson093/EVIDENCE.md` with exactly this structure:

```text
# Lesson 093 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 093 — Requirements, Capacity, and Queueing.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 093 — Requirements, Capacity, and Queueing without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've translated a vague performance goal into concrete numbers using real (or realistically estimated) business data, you've applied the queueing formulas correctly to derive a required capacity, and you've demonstrated — using your own recomputed numbers under a 50% traffic-estimate error — why headroom beyond your best-guess estimate is a genuine enginee. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
