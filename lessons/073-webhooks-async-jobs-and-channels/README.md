# Lesson 073 — Webhooks, Async Jobs, and Channels

## Threads
**Primary:** API/Go/Distributed

**Supporting:** event delivery; queues; worker lifecycles

**Math companion:** random variables and latency

**Time:** 3–4 hours. **Prerequisites:** Lesson 072.

**Time:** 3–4 hours. **Prerequisites:** Lesson 072.

## What you will be able to do

1. Explain the central model of # 073 — Webhooks, Async Jobs, and Channels and name the boundary where it applies.
2. Trace one concrete case of # 073 — Webhooks, Async Jobs, and Channels before executing it.
3. Rebuild the core exercise from an empty workspace without copying the reference implementation.
4. Produce terminal evidence for one success case and one failure case.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 072, # Lesson 072 — Concurrent Updates and Consistency.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Design and implement outbound webhooks correctly — signed payloads for authenticity, retry logic for delivery failures, and idempotent receiver-side handling — understanding webhooks as essentially an inverted API call where your server becomes the client and must handle all the same reliability problems from the other side.

**The contract.** A webhook sender delivers an event to a receiver over HTTP. Delivery commonly uses retries, so the receiver must tolerate duplicate event IDs. The sender signs the exact body bytes with a shared-secret HMAC; the receiver verifies the MAC before parsing/processing the payload and stores the event ID atomically with its effect.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** The sender is an HTTP client and must set timeouts, classify retryable outcomes, use bounded exponential backoff with jitter, and persist delivery attempts if delivery must survive process restarts. A successful HTTP response is defined by the receiver contract; retries can still duplicate delivery.

**Second point.** HMAC authenticates possession of a shared secret and detects body modification; it does not encrypt the payload. Use the standard library's HMAC verification (`hmac.Equal` in Go), compare the MAC for the exact raw body, and protect/rotate the secret.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Implement a webhook receiver (a separate small server) that verifies the incoming signature against its own copy of the shared secret before processing the payload, rejecting (with an appropriate 4xx status) any request with a missing or incorrect signature.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** For step 2, show both the accepted (valid signature) and rejected (invalid/tampered signature) request/response pairs, confirming signature verification is genuinely functioning.

**Failure mode.** Change one byte in the body while keeping the original signature and verify rejection. Then replay the exact valid event ID and verify no second side effect occurs. Do not attempt to demonstrate a practical timing attack in this lab; use the documented constant-time comparison API.

**Completion criterion.** You can sign/verify exact payload bytes, distinguish tampering from transient delivery failure, apply bounded retries, and make receiver processing idempotent by event ID.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the receiver result for a correct signature and an altered body.
2. Predict the result of delivering the same event ID twice.
3. Predict which failures merit bounded retry and which are permanent configuration errors.
4. Predict what an HTTP timeout leaves unknown about receiver processing.
5. Predict why signature verification must use raw body bytes.
6. Predict the effect of an expired delivery deadline.
7. Predict how an outbox improves durability of event publication.
8. State which database transaction protects deduplication and receiver-side effects.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.

2. Extend it with one explicit boundary case: Implement a webhook receiver (a separate small server) that verifies the incoming signature against its own copy of the shared secret before processing the payload, rejecting (with an appropriate 4xx status) any request with a missing or incorrect signature.

3. Rebuild the same behavior with one changed constraint: Test both a correctly signed request (accepted) and a tampered payload with a stale signature (rejected) to confirm signature verification is actually functioning, not just present in code but unused.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Implement a receiver that verifies the payload signature using a naive, non-constant-time string comparison (e.g. a simple `==` or `strcmp`-style comparison rather than a timing-safe comparison function) and research (via documentation on your language's cryptography library, e.g.

2. Go's `crypto/hmac.Equal` or `subtle.ConstantTimeCompare`) why this specific implementation detail matters — a naive comparison can leak timing information about how many leading characters of the signature matched before it found a mismatch, which is (in principle, given a very large number of attempts) an exploitable timing side-channel for an attacker atte.

3. You don't need to demonstrate an actual successful timing attack (genuinely difficult to do reliably in a simple lab exercise), but explain in your own words why "use a constant-time comparison for secret verification" is a real, specific, non-obvious security requirement — the difference between the naive and safe version is invisible in normal functional t.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 073 — Webhooks, Async Jobs, and Channels: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does HMAC provide, and what does it not provide?
2. Why sign the exact raw request body?
3. Why can a webhook receiver see duplicate deliveries?
4. How should the receiver make event processing idempotent?
5. What properties should delivery retries have?
6. Why use a constant-time MAC comparison API?

## Evidence to record

Create `lesson073/EVIDENCE.md` with exactly this structure:

```text
# Lesson 073 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Signing reserialized JSON instead of the received raw bytes.
- Treating HMAC as encryption.
- Retrying forever or retrying permanent failures without a bound.
- Processing duplicate event IDs as new side effects.
- Checking a signature with a hand-written comparison rather than the crypto library's constant-time API.

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

You pass when you can explain the central model of # 073 — Webhooks, Async Jobs, and Channels without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your webhook sender correctly signs payloads and your receiver correctly verifies and rejects tampered ones, your sender's retry logic correctly persists through transient failures with proper exponential backoff, you've implemented and tested idempotent receiver-side processing using a unique event ID, and you can explain why constant-time signature compari. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
