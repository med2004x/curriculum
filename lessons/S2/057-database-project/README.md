# Lesson 057 — Integrated Lab 2: Durable Data Service

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Integrate schema design, migrations, database access, HTTP behavior, tests, and operations into one durable service with clear ownership boundaries.

**The contract.** A task has a stable identifier, required title, status from an allowed domain, and timestamps. HTTP handlers validate requests and map outcomes to status codes; a service owns business rules; a repository owns SQL and persistence. Migrations create/upgrade schema. Tests cover valid operations, not-found, invalid input, database constraint failures, and cancellation.

**Worked example.** Implement `POST /tasks`, `GET /tasks/{id}`, `PATCH /tasks/{id}`, and `DELETE /tasks/{id}`. Ensure one invalid status is rejected consistently by both request validation and the database constraint, and that a missing ID is not reported as an internal server error.

## 2. Mechanism and vocabulary

**Mechanism.** Each dependency crosses an explicit interface: handler parses/serializes HTTP, service enforces domain rules, repository maps domain operations to parameterized SQL. Context cancellation reaches database calls; transactions group multi-statement state changes.

**Second point.** Integration tests require a known schema and data state. Migrations should run against a clean test database, each test should use isolated identifiers or cleanup, and failure paths must not leave persistent partial state.

Return a stable API error contract: malformed input is a client error; missing resources have a defined not-found response; unexpected database failures are server errors with details logged internally, not leaked to the caller.

**Worked example.** Use the assigned request path as a review checklist: HTTP handler → `TaskService` → `TaskRepository` interface → `PostgresTaskRepository` → `database/sql` pool → PostgreSQL. Keep the service contract independent of SQL driver types.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
HTTP handler
-> TaskService
-> TaskRepository interface
<- PostgresTaskRepository
-> database/sql pool
-> PostgreSQL
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** Check migration from an empty database, CRUD status/body behavior, repository no-row mapping, transaction rollback, and an integration fixture's cleanup. Distinguish tests using a fake repository from tests requiring PostgreSQL.

**Failure mode.** Cause a database constraint violation during a request and verify that it does not become a success response or leak SQL internals. Cancel a request context and verify the repository operation stops or returns an error.

**Completion criterion.** You can rebuild the service from the API contract, trace one request through each layer, verify migrations and persistent state, and explain the test and recovery evidence.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
