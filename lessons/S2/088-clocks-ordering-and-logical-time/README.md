# Lesson 088 — Clocks, Ordering, and Logical Time

## Threads
**Primary:** Distributed/Math

**Supporting:** Lamport-style reasoning; ordering without global time

**Math companion:** discrete math and relations

**Time:** 3–4 hours. **Prerequisites:** Lesson 087.

**Time:** 3–4 hours. **Prerequisites:** Lesson 087.

## What you will be able to do

1. Define happened-before for process-local events and message send/receive.
2. Implement Lamport `tick` and `receive` updates and verify causal ordering.
3. Explain why Lamport timestamps do not identify concurrent events or measure physical time.
4. Use an event graph or vector-clock comparison to distinguish concurrency where needed.

## Retrieve first (no notes)

1. What does an RPC timeout tell the caller, and what remains unknown?
2. Why can a retry repeat a side effect?
3. What does an idempotency key allow the server to recognize?
4. How can a bounded queue respond to overload?

## 1. Problem and model

**What problem this solves.** Understand why wall-clock timestamps cannot reliably order events across different machines, and implement Lamport logical clocks to establish a causally-consistent ordering instead.

**The contract.** Happened-before (`→`) is a causal partial order: program order and message send-before-receive imply order, and the relation is transitive. If neither `a → b` nor `b → a`, the events are concurrent. Lamport clocks guarantee `a → b` implies `L(a) < L(b)`, but not the converse.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a Lamport clock as a small class/struct with a `tick()` method (increment for a local event, return new value) and a `receive(remoteTimestamp)` method (implementing the max-plus-one rule from Learn).

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Each process increments its local counter before a local event. On receiving a message with timestamp `t`, update local time to `max(local,t)+1`; include the resulting counter in outgoing messages.

**Second point.** Lamport timestamps can produce a deterministic total ordering when combined with a process ID, but that order does not prove events were causally related. Physical clocks estimate real time and require synchronization assumptions and uncertainty bounds.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Write a unit test confirming the basic increment behavior in isolation.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a Lamport clock as a small class/struct with a `tick()` method (increment for a local event, return new value) and a `receive(remoteTimestamp)` method (implementing the max-plus-one rule from Learn).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Produce an event table containing process, event, message relation, Lamport timestamp, and happened-before/concurrent classification. Confirm every causal edge has increasing Lamport time and include a concurrent pair whose timestamps are still numerically ordered.

**Failure mode.** Assign different synthetic wall-clock offsets to two processes and show that wall-clock order can contradict a known message-causality edge. Explain that Lamport time preserves that edge but does not measure elapsed time.

**Completion criterion.** You can implement Lamport updates, classify causality versus concurrency from an event graph, and explain the limits of logical and wall-clock timestamps.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the local timestamp after a local event.
2. Predict the timestamp after receiving remote time `t`.
3. Does `L(a) < L(b)` prove that `a` happened before `b`?
4. Identify concurrent events in two independent processes.
5. Predict how synthetic clock offset changes wall-clock ordering.
6. State the causal guarantee Lamport clocks preserve.
7. State what additional information vector clocks provide.
8. Distinguish logical order from elapsed-time measurement.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement a Lamport clock as a small class/struct with a `tick()` method (increment for a local event, return new value) and a `receive(remoteTimestamp)` method (implementing the max-plus-one rule from Learn).

2. Extend it with one explicit boundary case: Write a unit test confirming the basic increment behavior in isolation.

3. Rebuild the same behavior with one changed constraint: Simulate 3 independent "processes" (just 3 instances of your Lamport clock in one program, or 3 actual goroutines/threads communicating via channels/sockets) that occasionally send messages to each other carrying their current Lamport timestamp, with the receiver updating its own clock per the receive rule.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Modify your step 2 simulation to also record and print each event's *wall-clock* timestamp (using each process's own local system clock, `time.Now()` or equivalent) alongside its Lamport timestamp — and, to make the failure concrete rather than requiring naturally-occurring clock drift, deliberately introduce a synthetic clock offset for one of your simulate.

2. Construct a scenario where process A sends a message to process B (A's send genuinely, causally precedes B's receive), but due to your artificial clock skew, B's wall-clock-timestamped receive event shows an *earlier* wall-clock time than A's send event — while the Lamport timestamps for the same two events correctly still show the send before the receive.

3. Explain, using this concrete, deliberately-constructed contradiction, exactly why wall-clock ordering failed here specifically (clock skew) while Lamport ordering didn't (it never depends on wall-clock values at all, only on the message-passing structure itself).

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 088 — Clocks, Ordering, and Logical Time: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 088 — Clocks, Ordering, and Logical Time?
2. What two concrete practices are required in the assigned material for # 088 — Clocks, Ordering, and Logical Time?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson088/EVIDENCE.md` with exactly this structure:

```text
# Lesson 088 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 088 — Clocks, Ordering, and Logical Time.
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

You pass when you can explain the central model of # 088 — Clocks, Ordering, and Logical Time without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your Lamport clock implementation correctly maintains the happens-before-implies-smaller-timestamp guarantee across a real, traced multi-process simulation, you've identified and confirmed at least one pair of genuinely concurrent events where the timestamps provide no ordering information, and you've constructed a concrete, reproducible case (via the failur. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
