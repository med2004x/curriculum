# Lesson 086 — Distributed Failures and System Models

## Threads
**Primary:** Distributed/Networking

**Supporting:** partial failure; timing models; partitions

**Math companion:** probability and sets

**Time:** 3–4 hours. **Prerequisites:** Lesson 085.

**Time:** 3–4 hours. **Prerequisites:** Lesson 085.

## What you will be able to do

1. Distinguish crash, omission, delay, and network-partition failures.
2. Explain why a timeout cannot establish whether a remote operation completed.
3. State the timing and failure assumptions of a distributed system model.
4. Explain why retrying a non-idempotent operation can duplicate effects.

## Retrieve first (no notes)

1. What invariant does the selected database component preserve?
2. Which neighboring component relies on that invariant?
3. Why does a passing component test not prove integration correctness?
4. What workload and environment are needed to interpret a measurement?

## 1. Problem and model

**What problem this solves.** Understand why "partial failure" — the defining problem of distributed systems — has no equivalent in single-machine programming, and build the vocabulary (crash, omission, partition, delay) to reason precisely about what can go wrong.

**The contract.** A distributed client observes messages, responses, and elapsed time, not the remote process's internal state. A timeout means no response was observed before a deadline; the request may have been lost, delayed, processed with its response lost, or still be running.

**Worked example.** Start with the smallest concrete case described by the source material: For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Define the system model: which components may crash, whether messages may be lost, duplicated, or reordered, what delay bounds exist, and how nodes recover. Safety and liveness claims depend on these assumptions.

**Second point.** A network partition separates communication paths without proving either endpoint has crashed. From one side, a sufficiently long delay can be indistinguishable from a failed or partitioned peer.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** For request lost, server crash before processing, server processes but response is lost, and response delayed: state the client-visible result under timeout-only behavior and identify what extra operation ID/status lookup could clarify.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Use a controlled client/server model to record request send, server-side effect, response loss/delay, and client timeout. Compare client symptoms and state which internal outcomes remain ambiguous.

**Failure mode.** Retry an append after the server committed but the response was lost. Trace the duplicate effect, then explain why an idempotency key or operation-status protocol is needed before retrying safely.

**Completion criterion.** You can classify failure modes, state the system model, explain timeout ambiguity, and identify when idempotency or reconciliation is needed to make retries safe.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 086 — Distributed Failures and System Models case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself.

2. Extend it with one explicit boundary case: For each point, state whether the client, from a plain timeout alone, can distinguish it from any of the other points.

3. Rebuild the same behavior with one changed constraint: Using `tc netem` (networking Lesson 7's tool) or a similar mechanism, simulate a network partition between a client and server (block all traffic between them for a period, e.g. via a firewall rule or `tc` loss configuration set to 100%), and observe what your client-side code actually experiences — a timeout, a connection refused, or a hang, depending on yo.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take your step 3 "naive retry" design and actually reason through what happens if the client retries an append operation after *not* receiving a response, in the specific case where the server had actually already succeeded (the response was lost, not the original request) — trace explicitly that this naive approach results in the item being appended twice,.

2. Explain why this single scenario is the entire motivating case for idempotency (which Lesson 2 and Lesson 7 both address with specific mechanisms) — the retry itself isn't wrong (it's often the only reasonable response to an ambiguous timeout), but a retry without some means of detecting "this was already done" turns ambiguity into an actual, guaranteed data.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 086 — Distributed Failures and System Models: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 086 — Distributed Failures and System Models?
2. What two concrete practices are required in the assigned material for # 086 — Distributed Failures and System Models?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson086/EVIDENCE.md` with exactly this structure:

```text
# Lesson 086 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 086 — Distributed Failures and System Models.
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

You pass when you can explain the central model of # 086 — Distributed Failures and System Models without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can name and distinguish the four failure modes covered here for a concrete scenario, you've directly observed that a network partition and a genuine crash are indistinguishable from a client's timeout-only perspective, and you can trace, using your own step 3/failure-drill reasoning, exactly why naive retry-on-timeout is unsafe without an additional mec. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
