# Lesson 076 — Integrated Project 3: Production-Grade Go Service

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Integrate the curriculum's API, database, security, testing, and operations practices into one coherent service whose contracts and failure behavior remain understandable under real boundaries.

**The contract.** Select one bounded resource domain and document routes, request/response schemas, invariants, authentication/authorization policy, error outcomes, migration plan, retry/idempotency behavior, operational endpoints, and shutdown behavior. Implement only cross-cutting features that the domain needs; do not add nominal checklist features without a clear requirement.

**Worked example.** Start with one resource such as a task or appointment. Define its schema and lifecycle, then trace create/read/update through handler → service → repository → PostgreSQL. Add one conflict-sensitive update and one protected route; state how each failure maps to the public API.

## 2. Mechanism and vocabulary

**Mechanism.** Cross-feature composition is the assessment: e.g. an idempotent request retried after a concurrent update, authorization checked before resource disclosure, and request cancellation during a transaction. State the invariant, interleaving, expected response, and durable state for each interaction test.

**Second point.** Keep the system operable: configuration comes from the environment, health/readiness semantics are explicit, logs avoid secrets, dependencies are closed in order, and the service exits within a bounded shutdown deadline.

**Worked example.** Write the API contract before implementation, including one success and at least three distinct failure outcomes. Add one integration test that exercises more than one feature boundary, not merely one unit test per package.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
handler -> service -> repository interface -> PostgreSQL repository -> database/sql pool
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the contract at each arrow and one error that must be translated rather than passed through unchanged.

## 4. Verification, failure, and evidence

**Verification.** Exercise migration from an empty database, happy-path routes, validation/not-found/authorization failures, one transaction conflict or idempotent retry, dependency unavailability, and bounded shutdown. Record which tests use fakes and which require PostgreSQL/container dependencies.

**Failure mode.** Choose one cross-feature interaction and break one ordering/invariant deliberately: skip ownership authorization after authentication, lose idempotency reservation under concurrency, or close the database before handlers drain. Show the observable failure, then repair the boundary.

**Completion criterion.** The implemented scope matches the written contract, integration tests cover cross-feature behavior, operational dependencies are reproducible, and you can explain one deliberately found integration bug and its correction.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
