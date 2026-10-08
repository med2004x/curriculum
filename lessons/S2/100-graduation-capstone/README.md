# Lesson 100 — Graduation: Design, Build, Measure, Explain

## Threads
**Primary:** All threads

**Supporting:** independent system; evidence; architecture review; communication

**Math companion:** choose and justify required math

**Time:** 3–4 hours. **Prerequisites:** Lesson 099.


## What you will be able to do

1. Deliver one deployable service connecting contract, network, storage, and operations.
2. Measure a declared workload and compare it with a transparent capacity estimate.
3. Diagnose a controlled incident from telemetry and state evidence.
4. Explain design tradeoffs, limits, and operating steps to another engineer.

## Retrieve first (no notes)

1. How do you preserve a pre-fix reproduction?
2. What makes a regression test useful?
3. Which metadata traces an artifact to source?
4. What questions do logs, metrics, and traces answer?

## 1. Problem and model

**What problem this solves.** Build a complete, production-shaped network service integrating correct networking, persistence, observability, and failure handling — the first of four terminal capstones drawing on nearly every track in this curriculum, proving the pieces compose into something a real team could actually operate.

**The contract.** The capstone adds no isolated technology; it tests integration. A small, complete vertical path with explicit limits is stronger evidence than a broad feature list that has not been run end to end.

**Worked example.** Start with the smallest concrete case described by the source material: Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Contract-first design fixes the user-visible interface before implementation. Persistence invariants, idempotency, bounded work, timeouts, telemetry, migrations, and recovery must agree at the boundaries between components. A load test states offered rate, concurrency, data shape, duration, and the point where latency or errors show saturation.

**Second point.** Production-shaped means that the service has repeatable setup, health checks, bounded behavior under failure, traceable releases, and documented recovery. Its scope is intentionally modest; evidence and operability matter more than feature count.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** **Contract-first design** (API engineering Lessons 1-3): a written requirements document and OpenAPI spec before implementation.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** **Load test it** (production Lesson 6): a real, open-loop load test finding your actual capacity and bottleneck, compared against a theoretical capacity estimate (system-engineering Lesson 2).

**Failure mode.** Break a cross-component assumption that isolated tests miss: for example, let a retry repeat a committed write without an idempotency key, or make a circuit-breaker rejection invisible to its metric. Diagnose from user-path evidence, repair it, and add an integration test that fails on the broken behavior.

**Completion criterion.** Another engineer can build and run the service from documented steps, reproduce its measured workload and synthetic incident, and identify remaining risks without relying on unstated assumptions.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 100 — Graduation: Design, Build, Measure, Explain case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:

2. Extend it with one explicit boundary case: **Contract-first design** (API engineering Lessons 1-3): a written requirements document and OpenAPI spec before implementation.

3. Rebuild the same behavior with one changed constraint: **Correct networking fundamentals** (networking Lessons 1-4): proper framing, correct HTTP semantics, connection reuse.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

Break a cross-component behavior that unit tests miss, such as a retry after commit without idempotency or a breaker rejection missing its metric. Diagnose, fix, and add an integration regression test.

Save the baseline, command, observed failure, and repaired result. Confirm the original success case still passes after the repair.

## Transfer

Change exactly one constraint in the main case for # 100 — Graduation: Design, Build, Measure, Explain: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Which user path defines the service objective?
2. What invariant must hold after retry?
3. How should one dependency failure degrade?
4. What workload reveals the capacity limit?
5. Which telemetry should expose the injected failure?
6. What risks remain after the demo?

## Evidence to record

Create `lesson100/EVIDENCE.md` with exactly this structure:

```text
# Lesson 100 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 100 — Graduation: Design, Build, Measure, Explain.
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

1. Record requirements, explicit exclusions, and API contract.
2. Identify the main durable-state invariant.
3. Document clean build, migration, start, and health commands.
4. State offered load, data shape, and test duration.
5. Save throughput, latency percentiles, and error rate.
6. Preserve telemetry from the controlled incident.
7. List remaining risks and the next operational action.

## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the lesson contract, rebuild its core case from a clean workspace, reproduce and repair its failure drill, solve the transfer constraint, and meet this evidence bar: Another engineer can build and run the service from documented steps, reproduce its measured workload and synthetic incident, and identify remaining risks without relying on unstated assumptions. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
