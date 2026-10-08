# Lesson 087 — RPC, Serialization, and Backpressure

## Threads
**Primary:** Distributed/Go/Networking

**Supporting:** RPC boundaries; encoding; overload

**Math companion:** queueing

**Time:** 3–4 hours. **Prerequisites:** Lesson 086.

**Time:** 3–4 hours. **Prerequisites:** Lesson 086.

## What you will be able to do

1. Define an RPC request/response contract, wire encoding, and timeout behavior.
2. Explain why a remote call is not semantically equivalent to a local function call.
3. Apply bounded queues and backpressure so overload does not create unbounded memory growth.
4. Use request IDs/idempotency where retries could repeat a side effect, without claiming generic exactly-once delivery.

## Retrieve first (no notes)

1. What does a timeout tell the caller, and what does it leave unknown?
2. Why can a retry repeat a side effect?
3. What does an idempotency key allow the server to recognize?
4. How does a network partition differ from a definite server error from the client's view?

## 1. Problem and model

**What problem this solves.** Build a minimal RPC (Remote Procedure Call) system with proper timeouts and retries, and solve Lesson 1's exact duplicate-request problem with idempotency — turning the previous lesson's abstract danger into a concretely handled case.

**The contract.** An RPC call serializes a request, sends it over a network, waits for a response, and decodes it. The call may time out after the server has performed its side effect; transport failure cannot by itself reveal whether processing occurred. Define method/version compatibility and explicit error semantics.

**Worked example.** Start with the smallest concrete case described by the source material: Build a minimal RPC client/server (Go's built-in `net/rpc` package is a reasonable, quick way to get a working RPC mechanism without hand-rolling wire-format parsing, though you're welcome to build directly on networking Lesson 1's raw sockets/framing if you want the deeper exercise) with at least one method that takes an argument and returns a result.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Use a context/deadline for bounded client waiting, but remember cancellation does not guarantee the remote server stops. The server must observe cancellation where possible, and side-effecting methods need an idempotency/reconciliation strategy if callers retry.

**Second point.** Serialization is a protocol contract: field names/numbers, types, optional values, versioning, and size limits must be compatible across sender and receiver. Unknown/optional fields and schema evolution should have documented behavior.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Add a deadline to the RPC client, bound the server's work queue, and define what happens when the queue is full: reject, shed load, or wait for a bounded duration. Expose queue depth/rejection metrics without allowing unbounded accumulation.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Build a minimal RPC client/server (Go's built-in `net/rpc` package is a reasonable, quick way to get a working RPC mechanism without hand-rolling wire-format parsing, though you're welcome to build directly on networking Lesson 1's raw sockets/framing if you want the deeper exercise) with at least one method that takes an argument and returns a result.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Test normal response, timeout, server error, malformed payload, full queue, and duplicate request ID. Record client-visible outcomes and server-side side-effect count; specify the conditions under which a retry is safe.

**Failure mode.** Disable request-ID deduplication and simulate a processed request whose response is lost. Show that a retry can duplicate the effect, then restore deduplication and explain its retention/state assumptions.

**Completion criterion.** You can state the RPC wire/error contract, bound waiting and queue growth, and explain why idempotency makes selected retries safe without promising exactly-once transport delivery.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 087 — RPC, Serialization, and Backpressure case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Build a minimal RPC client/server (Go's built-in `net/rpc` package is a reasonable, quick way to get a working RPC mechanism without hand-rolling wire-format parsing, though you're welcome to build directly on networking Lesson 1's raw sockets/framing if you want the deeper exercise) with at least one method that takes an argument and returns a result.

2. Extend it with one explicit boundary case: Add a client-side timeout: if the server doesn't respond within a configured duration, the client call returns an error rather than blocking indefinitely.

3. Rebuild the same behavior with one changed constraint: Test it by adding an artificial delay to the server's handler (longer than your configured timeout) and confirming the client correctly times out rather than hanging.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Disable your step 3 idempotency mechanism (revert to the naive retry from Lesson 1's failure drill — no request-ID tracking) and rerun step 4's exact scenario.

2. Confirm the item is now appended twice, reproducing the duplicate-write bug concretely, with your own server's state as the evidence, rather than only as an abstract concern from Lesson 1.

3. Then explain, in your own words, precisely what specific piece of state (the server's request-ID history) is what separates the working version from the broken one — the retry logic itself didn't change between the two versions, only whether the server could recognize a retry as a retry.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 087 — RPC, Serialization, and Backpressure: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 087 — RPC, Serialization, and Backpressure?
2. What two concrete practices are required in the assigned material for # 087 — RPC, Serialization, and Backpressure?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson087/EVIDENCE.md` with exactly this structure:

```text
# Lesson 087 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 087 — RPC, Serialization, and Backpressure.
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

You pass when you can explain the central model of # 087 — RPC, Serialization, and Backpressure without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your RPC client correctly times out rather than hanging indefinitely on an unresponsive server, and you've directly demonstrated — with your own server's observable state, not just a warning read in this lesson — both the duplicate-write bug from naive retries and its fix via request-ID-based idempotency, and can explain precisely what state the fix depends. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
