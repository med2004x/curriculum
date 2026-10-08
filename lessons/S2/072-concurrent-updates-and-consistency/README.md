# Lesson 072 — Concurrent Updates and Consistency

## Threads
**Primary:** API/DB/OS

**Supporting:** lost updates; preconditions; locking

**Math companion:** relations and invariants

**Time:** 3–4 hours. **Prerequisites:** Lesson 071.

**Time:** 3–4 hours. **Prerequisites:** Lesson 071.

## What you will be able to do

1. Explain the central model of # 072 — Concurrent Updates and Consistency and name the boundary where it applies.
2. Trace one concrete case of # 072 — Concurrent Updates and Consistency before executing it.
3. Rebuild the core exercise from an empty workspace without copying the reference implementation.
4. Produce terminal evidence for one success case and one failure case.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 071, # Lesson 071 — Idempotency, Transactions, and Retries.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Detect and correctly reject conflicting concurrent updates to the same resource using optimistic concurrency control (ETags/version numbers), rather than letting the last write silently win and destroy an earlier, unseen change.

**The contract.** **The lost-update problem, restated at the API level.** Two clients both `GET /clinics/{id}`, both see the same current state, both independently decide on a change, both `PUT` their updated version back — the second `PUT` overwrites the first, silently discarding the first client's change, with neither client aware anything went wrong.

**Worked example.** Start with the smallest concrete case described by the source material: Implement `GET /resource/{id}` returning an `ETag` header (a simple version integer, or a hash of the resource's serialized content, is fine) alongside the resource body.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** This is precisely database-internals Lesson 5's lost-update anomaly, except now it's happening across two separate HTTP requests (each internally a properly isolated database transaction) rather than within one database transaction — meaning the database's own isolation guarantees don't automatically protect against it, since each individual `PUT` request is.

**Second point.** **Optimistic concurrency control: detect conflicts instead of preventing them upfront.** Rather than locking the resource for the duration of a client's "think time" between GET and PUT (which would be both impractical over HTTP's stateless request model and a real availability cost, since one slow client would block others), optimistic concurrency lets both.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Implement `PUT /resource/{id}` requiring an `If-Match` header matching the resource's current ETag.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```sql
UPDATE app.clinics
SET name = $1,
		version = version + 1
WHERE id = $2
	AND version = $3
RETURNING id, name, version;
```

If no row is returned, the ID may not exist or the version is stale; resolve that distinction according to the API's disclosure policy. The version predicate and update occur as one database statement, so two clients with the same old version cannot both update successfully.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** For step 3, show both clients' actual requests (including their `If-Match` headers) and the server's actual responses, confirming the first succeeds (200 with new ETag) and the second is correctly rejected (412), with the resource's final state matching only the first client's intended change, not silently reflecting a merged or overwritten result.

**Failure mode.** Implement a version of `PUT` that *ignores* the `If-Match` header entirely (accepts and applies any update regardless of the client-provided version) and rerun step 3's exact concurrent-update scenario against this broken version.

**Completion criterion.** Your `PUT` endpoint correctly enforces `If-Match` and rejects stale updates with 412, you've directly reproduced the lost-update problem both with and without protection using the same concurrent-client test scenario, and you've implemented a working client-side retry-after-conflict flow that successfully reconciles and resubmits after a detected conflict ra.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 072 — Concurrent Updates and Consistency case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement `GET /resource/{id}` returning an `ETag` header (a simple version integer, or a hash of the resource's serialized content, is fine) alongside the resource body.

2. Extend it with one explicit boundary case: Implement `PUT /resource/{id}` requiring an `If-Match` header matching the resource's current ETag.

3. Rebuild the same behavior with one changed constraint: On a match, apply the update and return a new ETag; on a mismatch, return `412 Precondition Failed` with a clear error body explaining the resource has changed.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Implement a version of `PUT` that *ignores* the `If-Match` header entirely (accepts and applies any update regardless of the client-provided version) and rerun step 3's exact concurrent-update scenario against this broken version.

2. Confirm the second client's stale-based update now silently succeeds, overwriting the first client's change with no error or warning to either party — reproducing the original lost-update problem exactly, now with your own test as concrete evidence rather than an abstract description.

3. Explain why this specific bug — an endpoint that accepts `If-Match` syntactically but doesn't actually enforce it — would be easy to miss in casual testing (a single client updating a resource with no concurrent conflict would never trigger the difference) and would only surface under real concurrent usage, exactly the kind of bug that's invisible in typical.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 072 — Concurrent Updates and Consistency: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 072 — Concurrent Updates and Consistency?
2. What two concrete practices are required in the assigned material for # 072 — Concurrent Updates and Consistency?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson072/EVIDENCE.md` with exactly this structure:

```text
# Lesson 072 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 072 — Concurrent Updates and Consistency.
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

You pass when you can explain the central model of # 072 — Concurrent Updates and Consistency without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your `PUT` endpoint correctly enforces `If-Match` and rejects stale updates with 412, you've directly reproduced the lost-update problem both with and without protection using the same concurrent-client test scenario, and you've implemented a working client-side retry-after-conflict flow that successfully reconciles and resubmits after a detected conflict ra. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
