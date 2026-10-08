# Lesson 096 — Benchmarking, Profiling, and Locality

## Threads
**Primary:** Performance/Go/Linux

**Supporting:** benchmarks; pprof; perf; memory locality

**Math companion:** measurement and distributions

**Time:** 3–4 hours. **Prerequisites:** Lesson 095.


## What you will be able to do

1. Define a representative workload and repeatable benchmark.
2. Report repeated measurements with units and variability.
3. Use CPU and allocation profiles to locate likely costs.
4. Explain how working-set size and layout affect locality.

## Retrieve first (no notes)

1. What does a histogram reveal about latency?
2. Why keep metric label values bounded?
3. How does sampling limit a conclusion?
4. How do throughput and latency differ?

## 1. Problem and model

**What problem this solves.** Produce measurements that can distinguish a real performance change from noise.

**The contract.** Keep workload, build, and environment explicit. Warm-up, repeated runs, variance, throughput, latency, dead-code elimination, and workload drift all affect what a benchmark can establish.

**Worked example.** Start with the smallest concrete case described by the source material: Benchmark two implementations of the same operation.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** **Why profiling, not intuition, should drive optimization effort.** Linux tools Lesson 5 established the core principle: intuition about where time is spent is often wrong.

**Second point.** Go `pprof` can sample CPU call paths and attribute allocations to sites; Linux profiling tools can also expose hardware behavior such as cache misses. Use profiles to select a candidate cost, then benchmark the same correct workload before and after one change.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Repeat enough times to report a distribution, not one favorite number.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Benchmark two implementations of the same operation.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Save repeated baseline and candidate measurements with units, input size, compiler flags, and environment. Include a CPU or allocation profile that identifies a candidate hotspot, then show correctness checks and whether the measured change exceeds run-to-run variation.

**Failure mode.** Make a *second*, unrelated change to your function at the same time as your intended fix (e.g. also reorder some unrelated logic, or add an unrelated micro-optimization you assumed was harmless) and re-benchmark.

**Completion criterion.** You can reproduce a stated baseline, profile a representative workload, make one change, and report whether the repeated results support a bounded performance claim.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 096 — Benchmarking, Profiling, and Locality case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Benchmark two implementations of the same operation.

2. Extend it with one explicit boundary case: Repeat enough times to report a distribution, not one favorite number.

3. Rebuild the same behavior with one changed constraint: Write (or use) a deliberately inefficient Go function — e.g. one doing unnecessary repeated string concatenation in a loop (a classic Go performance anti-pattern, since naive `+=` string concatenation reallocates on every iteration) — and benchmark it using Lesson 1's methodology to establish a baseline.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

Discard the benchmark result or return a constant, then inspect for eliminated work. Add a result sink and correctness assertion. Compare a hot small working set with a larger one.

Save the baseline, command, observed failure, and repaired result. Confirm the original success case still passes after the repair.

## Transfer

Change exactly one constraint in the main case for # 096 — Benchmarking, Profiling, and Locality: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What work is included when setup is timed?
2. How can a larger working set change results?
3. Why is one fast run weak evidence?
4. What does a CPU profile sample?
5. When is an allocation profile useful?
6. How can dead-code elimination corrupt a benchmark?

## Evidence to record

Create `lesson096/EVIDENCE.md` with exactly this structure:

```text
# Lesson 096 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 096 — Benchmarking, Profiling, and Locality.
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
## Lab checklist

Use this checklist to decide whether the exercise is reproducible and reviewable:

1. Record compiler, flags, machine, and input distribution.
2. Run correctness checks before timing.
3. Separate one-time setup from the timed operation.
4. Save raw repeated results with units.
5. Name the profile tool and profile type.
6. Compare hot-cache and large-working-set behavior.
7. Limit the conclusion to the workload measured.

## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the lesson contract, rebuild its core case from a clean workspace, reproduce and repair its failure drill, solve the transfer constraint, and meet this evidence bar: You can reproduce a stated baseline, profile a representative workload, make one change, and report whether the repeated results support a bounded performance claim. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
