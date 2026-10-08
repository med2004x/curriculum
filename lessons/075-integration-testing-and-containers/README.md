# Lesson 075 — Integration Testing and Containers

## Threads
**Primary:** Go/Production

**Supporting:** test environments; dependencies; reproducibility

**Math companion:** experimental design

**Time:** 3–4 hours. **Prerequisites:** Lesson 074.

**Time:** 3–4 hours. **Prerequisites:** Lesson 074.

## What you will be able to do

1. Choose unit, integration, or end-to-end scope for a stated behavior.
2. Start a disposable PostgreSQL dependency and run integration tests against an isolated database.
3. Build a container image with explicit runtime configuration and health/readiness expectations.
4. Distinguish ephemeral container filesystem state from persistent volume/database state.

## Retrieve first (no notes)

1. What order should a service follow when it shuts down?
2. What does a shutdown deadline bound?
3. Why should background work receive cancellation and be joined?
4. Why close the database only after in-flight users are finished?

## 1. Problem and model

**What problem this solves.** Verify the boundaries between HTTP, service logic, and PostgreSQL without turning every test into an opaque end-to-end test.

**The contract.** Cover test fixtures, isolated databases, HTTP test servers, transaction cleanup, deterministic data, and the distinction between unit, integration, and end-to-end tests.

**Worked example.** Start with the smallest concrete case described by the source material: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Containers package a filesystem/user-space environment while relying on host-kernel isolation mechanisms. Image layers are immutable build inputs; runtime environment, secrets, network endpoints, CPU/memory limits, and mounted volumes are separate configuration. A container restart does not make its writable layer durable.

**Second point.** A reliable integration test controls startup/readiness, unique test data, schema migration, and cleanup. A running container is not necessarily a ready database; wait for a health check or explicit connection readiness rather than sleeping a fixed interval.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Run the service and PostgreSQL in separate containers on a private network. Apply migrations before API tests, use a disposable test database or schema, and mount persistent storage only when the scenario is explicitly testing persistence across container replacement.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record image build inputs, health/readiness state, migration result, integration-test database identity, and cleanup outcome. Restart or replace the service container and state which data is expected to persist and why.

**Failure mode.** Start tests before PostgreSQL is ready, or reuse a shared database without isolating fixtures. Identify whether the failure is startup ordering, leaked test state, or application behavior; fix readiness and cleanup explicitly.

**Completion criterion.** You can choose the smallest test scope that still proves a specific boundary.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether a container's writable layer persists after replacement.
2. Predict what happens if integration tests begin before database readiness.
3. Predict which test scope can verify SQL constraints.
4. Predict how two parallel tests can interfere through shared fixture IDs.
5. Predict what state a mounted volume preserves.
6. Predict which environment variables or secrets must be supplied at runtime.
7. Predict how a migration failure should affect test startup.
8. State the cleanup mechanism before creating test data.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

2. Extend it with one explicit boundary case: Containerize one backend service and document which state is ephemeral and which state is persistent.

3. Replace the service container and verify persistent database state remains only when the database volume is retained.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Remove the database health/readiness gate and observe startup/test behavior when the dependency is unavailable.

2. Reuse fixture IDs across parallel tests and identify the resulting cross-test interference.

3. Restore explicit readiness and fixture isolation; explain why containerization alone does not make tests deterministic.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one test environment condition: use a fresh database per test suite or one shared database with isolated schemas. Explain the isolation and cleanup trade-off.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What behavior belongs in a unit test versus an integration test?
2. Why is container running distinct from dependency readiness?
3. Which container state is ephemeral by default?
4. How do volumes change persistence?
5. How are test fixtures isolated and cleaned up?
6. Which runtime configuration should not be baked into the image?

## Evidence to record

Create `lesson075/EVIDENCE.md` with exactly this structure:

```text
# Lesson 075 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Assuming a started database container is ready to accept connections.
- Sharing persistent test state between cases without isolation.
- Assuming a container's writable layer is durable.
- Baking secrets or environment-specific endpoints into an image.
- Calling a test end-to-end when it does not cross the relevant system boundary.

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

You pass when you can explain the central model of # 075 — Integration Testing and Containers without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can choose the smallest test scope that still proves a specific boundary. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
