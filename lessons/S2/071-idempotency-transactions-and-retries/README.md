# Lesson 071 — Idempotency, Transactions, and Retries

## Threads
**Primary:** API/DB/Distributed

**Supporting:** safe repetition; retry boundaries

**Math companion:** probability of duplicate work

**Time:** 3–4 hours. **Prerequisites:** Lesson 070.

## What you will be able to do

1. Explain why a client cannot know whether a timed-out write was processed.
2. Implement an idempotency-key state machine with a request fingerprint and stored response.
3. Make key reservation and side-effect recording safe under concurrent duplicate requests.
4. Define retention, replay, and mismatched-body behavior for an idempotency key.

## Retrieve first (no notes)

1. Why does a fixed-window limiter allow bursts at the boundary?
2. What does token-bucket capacity control?
3. Why might users behind one NAT share an IP-based limit?
4. How can a client use `Retry-After`?

## 1. Problem and model

**What problem this solves.** Implement idempotency keys for a real, non-naturally-idempotent write endpoint (a `POST` that creates something with a real side effect), directly applying distributed Lesson 2's mechanism at the API-design level, where it's most commonly needed in practice.

**The contract.** Scope an idempotency key to the authenticated principal and operation. Store a fingerprint of the request and a state such as in-progress/completed. Reusing the same key and same request returns or waits for the original result; reusing it with a different request is rejected. The side effect and durable idempotency result must be coordinated so a crash cannot create an unrecorded duplicate.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a `POST` endpoint with a real (simulated) side effect — e.g. "create an order" that increments a counter or appends to a list representing "orders placed," standing in for a real payment charge.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** A unique constraint on `(principal_id, operation, idempotency_key)` can serialize duplicate reservations. A transaction can record the key/result alongside a database-local side effect. For external effects, use an outbox or provider-supported idempotency; a database transaction cannot atomically commit an unrelated external payment network action.

**Second point.** Retrying is safe only when the operation is idempotent or protected by a stable idempotency key. Use bounded retries, backoff with jitter where appropriate, and retry only outcomes that may be transient; do not assume all 5xx or all network errors have identical semantics.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Confirm, without any idempotency protection, that sending the same logical request twice (two separate POST calls with identical body) produces two side effects (two orders created) — reproducing the base problem concretely before fixing it.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a `POST` endpoint with a real (simulated) side effect — e.g. "create an order" that increments a counter or appends to a list representing "orders placed," standing in for a real payment charge.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Send the same principal/key/body twice and confirm one durable side effect and replayed response. Send the same key with a different body and confirm rejection. Send concurrent duplicates and verify they cannot both execute the side effect.

**Failure mode.** Reserve only after performing the side effect. Simulate a crash between side effect and record persistence and explain why a later retry duplicates work. Then move reservation/side-effect coordination to a durable transaction or outbox/provider idempotency boundary.

**Completion criterion.** You can explain sequential and concurrent duplicate protection, request fingerprint matching, response replay, crash windows, and the limits of database transactions around external effects.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the state after a request completes but its response is lost.
2. Predict the result for a repeated key with the same request fingerprint.
3. Predict the result for a repeated key with a different body.
4. Predict what two simultaneous first requests can do without atomic reservation.
5. Predict the crash outcome if the external effect occurs before key/result storage.
6. Predict which operations a database transaction can atomically cover.
7. Predict how long an idempotency record should be retained for the API contract.
8. State which error outcomes are safe to retry.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement a `POST` endpoint with a durable side-effect record and a unique idempotency key scoped to the caller.

2. Return the stored response for a repeated same-key/same-fingerprint request; reject same-key/different-fingerprint reuse.

3. Add a concurrent reservation path and document how in-progress duplicates wait, receive a conflict, or receive the eventual stored result.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Persist the idempotency result only after performing the side effect and model a crash in between.

2. Explain why the retry can repeat the side effect and identify the atomicity gap.

3. Add atomic key reservation for a database-local operation, or an outbox/provider idempotency boundary for an external side effect.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one retention or replay policy: expire keys after a documented interval, or return the original status/body for all retries. Explain the effect on late retries and client guarantees.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What ambiguity causes clients to retry a write after timeout?
2. How is an idempotency key scoped?
3. Why store a request fingerprint?
4. How does concurrent reservation differ from checking only completed keys?
5. What does an idempotency record need to replay?
6. Why can a database transaction not cover an unrelated external side effect?

## Evidence to record

Create `lesson071/EVIDENCE.md` with exactly this structure:

```text
# Lesson 071 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Checking only completed keys while two concurrent requests are both in progress.
- Accepting the same key with a different request body.
- Performing an external side effect before durably reserving/recording the key.
- Retrying a non-idempotent request without a stable key.
- Assuming a database transaction makes an external payment atomic.

## Precision checks

1. State key scope, fingerprint rule, retention, and replay behavior.
2. Explain reservation and duplicate handling while a request is in progress.
3. Identify the side effect and durable record transaction boundary.
4. Show same-request replay and changed-body rejection.
5. State which retry outcomes are safe and bounded.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 071 — Idempotency, Transactions, and Retries without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've reproduced the duplicate-side-effect problem without protection, then implemented and verified idempotency-key-based deduplication actually prevents it across a real retry scenario, you've implemented and tested mismatched-body detection as a distinct case from a legitimate retry, and you've demonstrated — via the failure drill — the specific concurre. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
