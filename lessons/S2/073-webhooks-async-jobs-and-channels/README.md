# Lesson 073 — Webhooks, Async Jobs, and Channels

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Design and implement outbound webhooks correctly — signed payloads for authenticity, retry logic for delivery failures, and idempotent receiver-side handling — understanding webhooks as essentially an inverted API call where your server becomes the client and must handle all the same reliability problems from the other side.

**The contract.** A webhook sender delivers an event to a receiver over HTTP. Delivery commonly uses retries, so the receiver must tolerate duplicate event IDs. The sender signs the exact body bytes with a shared-secret HMAC; the receiver verifies the MAC before parsing/processing the payload and stores the event ID atomically with its effect.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.

## 2. Mechanism and vocabulary

**Mechanism.** The sender is an HTTP client and must set timeouts, classify retryable outcomes, use bounded exponential backoff with jitter, and persist delivery attempts if delivery must survive process restarts. A successful HTTP response is defined by the receiver contract; retries can still duplicate delivery.

**Second point.** HMAC authenticates possession of a shared secret and detects body modification; it does not encrypt the payload. Use the standard library's HMAC verification (`hmac.Equal` in Go), compare the MAC for the exact raw body, and protect/rotate the secret.

**Worked example.** Implement a webhook receiver (a separate small server) that verifies the incoming signature against its own copy of the shared secret before processing the payload, rejecting (with an appropriate 4xx status) any request with a missing or incorrect signature.

## 3. Worked case from the reference material

Worked example: Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.

## 4. Verification, failure, and evidence

**Verification.** For step 2, show both the accepted (valid signature) and rejected (invalid/tampered signature) request/response pairs, confirming signature verification is genuinely functioning.

**Failure mode.** Change one byte in the body while keeping the original signature and verify rejection. Then replay the exact valid event ID and verify no second side effect occurs. Do not attempt to demonstrate a practical timing attack in this lab; use the documented constant-time comparison API.

**Completion criterion.** You can sign/verify exact payload bytes, distinguish tampering from transient delivery failure, apply bounded retries, and make receiver processing idempotent by event ID.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
