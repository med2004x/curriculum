# Lesson 076 — Integrated Project 3: Production-Grade Go Service

## Threads
**Primary:** Go/DB/API/Operations

**Supporting:** complete service from requirements to tests

**Math companion:** capacity calculation

**Time:** 3–4 hours. **Prerequisites:** Lesson 075.

**Time:** 3–4 hours. **Prerequisites:** Lesson 075.

## What you will be able to do

1. Design and implement one production-shaped Go API over PostgreSQL from an explicit contract.
2. Integrate authentication/authorization, validation, transactions, idempotency or version checks, and observability.
3. Verify isolated integration behavior, migrations, container configuration, and graceful shutdown.
4. Explain architecture trade-offs and one cross-feature failure that tests expose.

## Retrieve first (no notes)

1. What readiness signal should gate integration tests?
2. How should integration fixtures avoid cross-test interference?
3. Which state is ephemeral in a container and which can persist through a volume?
4. Why should secrets be supplied at runtime rather than baked into an image?

## 1. Problem and model

**What problem this solves.** Integrate the curriculum's API, database, security, testing, and operations practices into one coherent service whose contracts and failure behavior remain understandable under real boundaries.

**The contract.** Select one bounded resource domain and document routes, request/response schemas, invariants, authentication/authorization policy, error outcomes, migration plan, retry/idempotency behavior, operational endpoints, and shutdown behavior. Implement only cross-cutting features that the domain needs; do not add nominal checklist features without a clear requirement.

**Worked example.** Start with one resource such as a task or appointment. Define its schema and lifecycle, then trace create/read/update through handler → service → repository → PostgreSQL. Add one conflict-sensitive update and one protected route; state how each failure maps to the public API.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Cross-feature composition is the assessment: e.g. an idempotent request retried after a concurrent update, authorization checked before resource disclosure, and request cancellation during a transaction. State the invariant, interleaving, expected response, and durable state for each interaction test.

**Second point.** Keep the system operable: configuration comes from the environment, health/readiness semantics are explicit, logs avoid secrets, dependencies are closed in order, and the service exits within a bounded shutdown deadline.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Write the API contract before implementation, including one success and at least three distinct failure outcomes. Add one integration test that exercises more than one feature boundary, not merely one unit test per package.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
handler -> service -> repository interface -> PostgreSQL repository -> database/sql pool
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the contract at each arrow and one error that must be translated rather than passed through unchanged.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Exercise migration from an empty database, happy-path routes, validation/not-found/authorization failures, one transaction conflict or idempotent retry, dependency unavailability, and bounded shutdown. Record which tests use fakes and which require PostgreSQL/container dependencies.

**Failure mode.** Choose one cross-feature interaction and break one ordering/invariant deliberately: skip ownership authorization after authentication, lose idempotency reservation under concurrency, or close the database before handlers drain. Show the observable failure, then repair the boundary.

**Completion criterion.** The implemented scope matches the written contract, integration tests cover cross-feature behavior, operational dependencies are reproducible, and you can explain one deliberately found integration bug and its correction.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict each layer's responsibility for one create request.
2. Predict the behavior for invalid input and unauthorized resource access.
3. Predict which integration tests need PostgreSQL.
4. Predict the stored state after a failed transaction.
5. Predict the response to a duplicate idempotency key.
6. Predict what happens when a dependency is unavailable at startup.
7. Predict the shutdown order for server, workers, and database pool.
8. Identify a cross-feature interleaving that can violate an invariant.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Design and build a complete API for a real resource type from TARDOC or Mahall's actual domain (or a plausible extension of it) with this minimum required scope:

2. Extend it with one explicit boundary case: **Contract-first**: a complete OpenAPI spec (Lesson 3) written before implementation, covering all endpoints below, with reusable error-schema components (Lesson 3).

3. Rebuild the same behavior with one changed constraint: **Resource model** (Lesson 2): at least one top-level resource and one genuinely nested sub-resource, with a justified nesting decision.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Pick one specific cross-feature interaction from your capstone and deliberately break it — for example, make your rate limiter apply *before* authentication (checking IP-based limits before knowing who the authenticated user is, contradicting Lesson 10's failure drill about per-identity limiting being preferable to per-IP), or make your async job endpoint (i.

2. Demonstrate the resulting vulnerability or incorrect behavior with a real test, then fix it, and explain in your own words why integration points between separately-correct features are disproportionately likely to hide bugs, compared to bugs within any single feature's own isolated logic.

3. Deliver a portion of your talk from memorized wording, then have someone interrupt you mid-sentence with an unrelated question, then ask you to continue from where you left off.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one contract requirement, such as adding a status transition or a new access policy. Update schema/migration, service behavior, API error contract, and tests together while preserving unrelated interfaces.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the selected resource and its invariant?
2. Which API errors are part of the public contract?
3. Which features must compose in the cross-feature integration test?
4. What runtime configuration and dependencies are required?
5. How does the service handle cancellation and shutdown?
6. What failure did the deliberate interaction test reveal?

## Evidence to record

Create `lesson076/EVIDENCE.md` with exactly this structure:

```text
# Lesson 076 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Implementing checklist features without a use case or contract.
- Testing authentication without testing resource ownership.
- Testing idempotency sequentially but not under concurrent first requests.
- Running integration tests against non-isolated data.
- Reporting a successful shutdown while workers or database users remain active.

## Precision checks

1. State the resource contract and invariant.
2. Trace one request through each layer.
3. Show cross-feature integration evidence for one concurrency/security/retry interaction.
4. Explain test isolation and dependency readiness.
5. List operational shutdown and recovery assumptions.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 076 — Integrated Project 3: Production-Grade Go Service without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your capstone API implements all 11 required elements with real, passing tests — not just each element in isolation, but including at least one deliberately-tested cross-feature interaction — and you've found, demonstrated, and fixed at least one genuine integration bug via the failure drill, with an honest explanation of why it specifically occurred at the. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
