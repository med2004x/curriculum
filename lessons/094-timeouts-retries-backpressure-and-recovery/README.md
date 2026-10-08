# Lesson 094 — Timeouts, Retries, Backpressure, and Recovery

## Threads
**Primary:** System Design/Go/Distributed

**Supporting:** bounded work; retry storms; recovery

**Math companion:** reliability math

**Time:** 3–4 hours. **Prerequisites:** Lesson 093.

**Time:** 3–4 hours. **Prerequisites:** Lesson 093.

## What you will be able to do

1. Propagate an end-to-end deadline through a multi-service call chain.
2. Retry only suitable failures with bounded exponential backoff, jitter, and an overall retry budget.
3. Apply backpressure through bounded queues, admission control, or load shedding.
4. Explain recovery behavior after overload without creating a retry storm.

## Retrieve first (no notes)

1. What assumptions underlie the M/M/1 response-time formula?
2. Why does mean delay grow rapidly as utilization approaches one?
3. Why are p95/p99 latency objectives distinct from mean latency?
4. How can a 50% demand underestimate affect a near-saturated service?

## 1. Problem and model

**What problem this solves.** Set timeouts and retry policies deliberately, based on the actual failure modes and cost tradeoffs involved, rather than copying default values — and understand cascading failure, the specific danger poorly-tuned timeouts/retries can cause across a system of interdependent services.

**The contract.** An operation has one end-to-end deadline shared across its dependency calls; nested services should not each restart a fresh full timeout. Retries consume the remaining deadline and a bounded budget. Retry only operations that are idempotent or carry a stable idempotency key, and distinguish transient failures from permanent errors.

**Worked example.** Start with the smallest concrete case described by the source material: For a real call chain in TARDOC or Mahall (or a plausible one — e.g. an API endpoint that calls a database, which itself might be under load), identify what timeout is currently configured (or would be, by a common default) at each hop, and reason about whether that value seems deliberately chosen or just a framework default.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Exponential backoff increases delay between attempts; jitter spreads synchronized clients; a retry budget caps extra work. A timeout bounds waiting but does not prove the remote operation did not complete.

**Second point.** Backpressure prevents unbounded work accumulation: bounded queues block or reject producers, admission control limits in-flight requests, and load shedding rejects lower-priority work. Recovery should ramp traffic gradually so a recovering dependency is not immediately overwhelmed.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Simulate A → B → C with a shared deadline. Let C slow down and compare bounded queue/rejection plus backoff against immediate unbounded retries. Track in-flight work, attempt count, queue depth, and final success/error.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: For a real call chain in TARDOC or Mahall (or a plausible one — e.g. an API endpoint that calls a database, which itself might be under load), identify what timeout is currently configured (or would be, by a common default) at each hop, and reason about whether that value seems deliberately chosen or just a framework default.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Under the same downstream slowdown, record deadline propagation, attempts, backoff delays, queue depth, in-flight work, and rejected requests for bounded versus naive retry behavior.

**Failure mode.** Configure a timeout shorter than healthy downstream latency and measure false timeouts/retries. Then set a realistic budget and explain the latency-versus-resource trade-off.

**Completion criterion.** You can propagate deadlines, bound retries/queues, reject overload deliberately, and explain how the combined policy limits cascading work and supports controlled recovery.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict how a child deadline relates to its caller's deadline.
2. Predict the maximum attempt count under a retry budget.
3. Predict how jitter changes synchronized retry traffic.
4. Predict whether retrying a non-idempotent operation is safe without a key.
5. Predict queue behavior when capacity is reached.
6. Predict what happens to latency when a timeout is too short.
7. Predict the recovery risk if all clients retry immediately after an outage.
8. State which overload metric should trigger admission control or shedding.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Draw a call chain and assign one end-to-end deadline across its hops.

2. Add bounded exponential retries with jitter only for a transient/idempotent operation.

3. Add a bounded queue or admission limit and define the overload response.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Configure your step 3 timeout to be deliberately too short (much shorter than C's *normal*, healthy response time, not just its slowed-down time) and rerun the simulation with C responding at its normal, healthy speed.

2. Confirm B now frequently times out on requests that would have succeeded, generating unnecessary failures and retries even though nothing is actually wrong with C — directly demonstrating the "too short" failure mode from Learn, using your own measured false-timeout rate as evidence, and connect this concretely to why timeout values need to be chosen based o.

3. Take your step 2 bounded-queue-with-rejection service and deliberately configure the queue capacity far too small (e.g. capacity of 1, when your actual expected burst size is much larger) relative to normal, healthy traffic patterns (not just overload conditions).

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one dependency or retry policy: alter the overall deadline, queue capacity, or maximum attempts. Explain how resource use, success probability, and overload response change.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Why propagate one deadline through a call chain?
2. Which operations are safe to retry?
3. What do exponential backoff and jitter each accomplish?
4. How does a retry budget limit amplification?
5. What is backpressure and how can a bounded queue enforce it?
6. Why can a timeout be too short or too long?

## Evidence to record

Create `lesson094/EVIDENCE.md` with exactly this structure:

```text
# Lesson 094 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Giving each downstream call a fresh timeout beyond the caller's overall deadline.
- Retrying non-idempotent operations without a stable idempotency mechanism.
- Retrying immediately or without a bounded attempt/budget policy.
- Allowing an unbounded queue to turn overload into memory exhaustion.
- Restoring full traffic instantly when a dependency begins recovering.

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

You pass when you can explain the central model of # 094 — Timeouts, Retries, Backpressure, and Recovery without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've directly demonstrated cascading resource exhaustion from an unbounded downstream slowdown, then directly demonstrated a correctly-set timeout bounding that exhaustion, using the same before/after simulation, and you've directly measured the concrete reduction in request volume that exponential backoff with jitter provides over naive immediate retries. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
