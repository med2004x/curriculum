# Lesson 071 — Idempotency, Transactions, and Retries

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Implement idempotency keys for a real, non-naturally-idempotent write endpoint (a `POST` that creates something with a real side effect), directly applying distributed Lesson 2's mechanism at the API-design level, where it's most commonly needed in practice.

**The contract.** Scope an idempotency key to the authenticated principal and operation. Store a fingerprint of the request and a state such as in-progress/completed. Reusing the same key and same request returns or waits for the original result; reusing it with a different request is rejected. The side effect and durable idempotency result must be coordinated so a crash cannot create an unrecorded duplicate.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a `POST` endpoint with a real (simulated) side effect — e.g. "create an order" that increments a counter or appends to a list representing "orders placed," standing in for a real payment charge.

## 2. Mechanism and vocabulary

**Mechanism.** A unique constraint on `(principal_id, operation, idempotency_key)` can serialize duplicate reservations. A transaction can record the key/result alongside a database-local side effect. For external effects, use an outbox or provider-supported idempotency; a database transaction cannot atomically commit an unrelated external payment network action.

**Second point.** Retrying is safe only when the operation is idempotent or protected by a stable idempotency key. Use bounded retries, backoff with jitter where appropriate, and retry only outcomes that may be transient; do not assume all 5xx or all network errors have identical semantics.

**Worked example.** Confirm, without any idempotency protection, that sending the same logical request twice (two separate POST calls with identical body) produces two side effects (two orders created) — reproducing the base problem concretely before fixing it.

## 3. Worked case from the reference material

Worked example: Implement a `POST` endpoint with a real (simulated) side effect — e.g. "create an order" that increments a counter or appends to a list representing "orders placed," standing in for a real payment charge.

## 4. Verification, failure, and evidence

**Verification.** Send the same principal/key/body twice and confirm one durable side effect and replayed response. Send the same key with a different body and confirm rejection. Send concurrent duplicates and verify they cannot both execute the side effect.

**Failure mode.** Reserve only after performing the side effect. Simulate a crash between side effect and record persistence and explain why a later retry duplicates work. Then move reservation/side-effect coordination to a durable transaction or outbox/provider idempotency boundary.

**Completion criterion.** You can explain sequential and concurrent duplicate protection, request fingerprint matching, response replay, crash windows, and the limits of database transactions around external effects.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
