# Lesson 095 — Observability, SLOs, and Measurement

## Threads
**Primary:** Production/Performance/Math

**Supporting:** logs/metrics/traces; SLOs; error budgets

**Math companion:** statistics and quantiles

**Time:** 3–4 hours. **Prerequisites:** Lesson 094.


## What you will be able to do

1. Track requests with correlated logs, metrics, and traces.
2. Define an event-based SLI, an SLO target, and its error budget.
3. Use telemetry to locate latency and distinguish observation from inference.
4. Explain sampling and metric-cardinality limits.

## Retrieve first (no notes)

1. How does a deadline constrain downstream work?
2. When is a retry safe?
3. What does a bounded queue do at capacity?
4. Why can tail latency differ from the mean?

## 1. Problem and model

**What problem this solves.** Implement the three pillars of observability — logs, metrics, and traces — in a way that lets you answer "why is this specific request slow/failing" after the fact, rather than only knowing *that* something is wrong in aggregate.

**The contract.** Monitoring checks known signals against expected thresholds; observability supplies enough linked evidence to investigate questions that were not anticipated. An SLI defines an eligible event and a good event; an SLO sets a target over a window, and its error budget is the permitted bad-event share.

**Worked example.** Start with the smallest concrete case described by the source material: Implement structured logging for a real endpoint (JSON-formatted, with at minimum a request ID, timestamp, and severity level on every log line), and generate a request ID at the start of each request, propagating it through every subsequent log line for that request (via context, in Go's idiomatic style, or an equivalent mechanism in your language).

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Correlated logs explain discrete events, metrics show aggregate behavior, and traces connect spans across service boundaries. A trace or request ID belongs in logs and trace context; metric labels should stay bounded so unique requests do not create unbounded time series.

**Second point.** **Logs: discrete, timestamped events, and why structure matters.** Unstructured log lines (`fmt.Println("processing request")`) are hard to search, filter, or aggregate at scale.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Implement basic metrics: a counter for total requests, a counter for errors, and a histogram for request latency (using a metrics library like Prometheus's client library, or a simple hand-rolled equivalent for this exercise).

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement structured logging for a real endpoint (JSON-formatted, with at minimum a request ID, timestamp, and severity level on every log line), and generate a request ID at the start of each request, propagating it through every subsequent log line for that request (via context, in Go's idiomatic style, or an equivalent mechanism in your language).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Send a known mix of fast, slow, successful, and failed requests. Show one request ID across logs and spans, show request and error counts, inspect latency histogram buckets, and calculate the SLI from the generated events. State whether the sample is large enough to support any SLO conclusion.

**Failure mode.** Attempt the same step 4 diagnosis using *only* unstructured, non-request-correlated logs (revert to plain `fmt.Println`-style logging with no request ID propagation, mixing log lines from many concurrent requests together with no way to tell which lines belong to which request).

**Completion criterion.** You can follow a request across logs and spans, calculate the stated SLI from raw counts, compare it with its SLO, and explain what the telemetry cannot establish.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 095 — Observability, SLOs, and Measurement case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement structured logging for a real endpoint (JSON-formatted, with at minimum a request ID, timestamp, and severity level on every log line), and generate a request ID at the start of each request, propagating it through every subsequent log line for that request (via context, in Go's idiomatic style, or an equivalent mechanism in your language).

2. Extend it with one explicit boundary case: Implement basic metrics: a counter for total requests, a counter for errors, and a histogram for request latency (using a metrics library like Prometheus's client library, or a simple hand-rolled equivalent for this exercise).

3. Rebuild the same behavior with one changed constraint: Expose them via a `/metrics` endpoint and confirm you can compute p50/p95/p99 latency from your histogram data.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

Release one trace ID during a dependency call and observe the broken trace; then add a high-cardinality request-ID metric label and inspect series growth. Restore propagation and remove the label.

Save the baseline, command, observed failure, and repaired result. Confirm the original success case still passes after the repair.

## Transfer

Change exactly one constraint in the main case for # 095 — Observability, SLOs, and Measurement: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Which events belong in the availability denominator?
2. How does event volume affect the error budget?
3. Why can p99 rise while the mean remains stable?
4. What information is lost through trace sampling?
5. Why are request IDs poor metric labels?
6. Which metric type represents in-flight requests?

## Evidence to record

Create `lesson095/EVIDENCE.md` with exactly this structure:

```text
# Lesson 095 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 095 — Observability, SLOs, and Measurement.
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

1. Write the SLI numerator and denominator before collecting data.
2. Record the request volume and the time window.
3. Show the trace ID shared by request logs and spans.
4. State histogram bucket boundaries and latency units.
5. Show metric labels and explain their bounded value sets.
6. Calculate the target error budget from the chosen volume.
7. Mark one conclusion that sampling prevents you from making.

## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the lesson contract, rebuild its core case from a clean workspace, reproduce and repair its failure drill, solve the transfer constraint, and meet this evidence bar: You can follow a request across logs and spans, calculate the stated SLI from raw counts, compare it with its SLO, and explain what the telemetry cannot establish. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
