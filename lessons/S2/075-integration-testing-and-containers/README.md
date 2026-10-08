# Lesson 075 — Integration Testing and Containers

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Verify the boundaries between HTTP, service logic, and PostgreSQL without turning every test into an opaque end-to-end test.

**The contract.** Cover test fixtures, isolated databases, HTTP test servers, transaction cleanup, deterministic data, and the distinction between unit, integration, and end-to-end tests.

**Worked example.** Start with the smallest concrete case described by the source material: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

## 2. Mechanism and vocabulary

**Mechanism.** Containers package a filesystem/user-space environment while relying on host-kernel isolation mechanisms. Image layers are immutable build inputs; runtime environment, secrets, network endpoints, CPU/memory limits, and mounted volumes are separate configuration. A container restart does not make its writable layer durable.

**Second point.** A reliable integration test controls startup/readiness, unique test data, schema migration, and cleanup. A running container is not necessarily a ready database; wait for a health check or explicit connection readiness rather than sleeping a fixed interval.

**Worked example.** Run the service and PostgreSQL in separate containers on a private network. Apply migrations before API tests, use a disposable test database or schema, and mount persistent storage only when the scenario is explicitly testing persistence across container replacement.

## 3. Worked case from the reference material

Worked example: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

## 4. Verification, failure, and evidence

**Verification.** Record image build inputs, health/readiness state, migration result, integration-test database identity, and cleanup outcome. Restart or replace the service container and state which data is expected to persist and why.

**Failure mode.** Start tests before PostgreSQL is ready, or reuse a shared database without isolating fixtures. Identify whether the failure is startup ordering, leaked test state, or application behavior; fix readiness and cleanup explicitly.

**Completion criterion.** You can choose the smallest test scope that still proves a specific boundary.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
