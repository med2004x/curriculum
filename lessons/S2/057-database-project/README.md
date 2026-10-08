# Lesson 057 — Integrated Lab 2: Durable Data Service

## Threads
**Primary:** Go/PostgreSQL/Linux

**Supporting:** schema + queries + repository + tests

**Math companion:** measurement and reporting

**Time:** 3–4 hours. **Prerequisites:** Lesson 056.

**Time:** 3–4 hours. **Prerequisites:** Lesson 056.

## What you will be able to do

1. Deliver a Go task API backed by PostgreSQL through a repository boundary.
2. Apply versioned migrations and enforce task invariants with database constraints.
3. Implement create/read/update/delete operations with explicit HTTP and storage errors.
4. Verify integration behavior with isolated database fixtures and document backup/restore evidence.

## Retrieve first (no notes)

1. Why separate migration and runtime roles?
2. What compatibility problem does expand-migrate-contract address?
3. Why is a successful dump not proof of recoverability?
4. What must be checked after restoring to an isolated target?

## 1. Problem and model

**What problem this solves.** Integrate schema design, migrations, database access, HTTP behavior, tests, and operations into one durable service with clear ownership boundaries.

**The contract.** A task has a stable identifier, required title, status from an allowed domain, and timestamps. HTTP handlers validate requests and map outcomes to status codes; a service owns business rules; a repository owns SQL and persistence. Migrations create/upgrade schema. Tests cover valid operations, not-found, invalid input, database constraint failures, and cancellation.

**Worked example.** Implement `POST /tasks`, `GET /tasks/{id}`, `PATCH /tasks/{id}`, and `DELETE /tasks/{id}`. Ensure one invalid status is rejected consistently by both request validation and the database constraint, and that a missing ID is not reported as an internal server error.

Self-check: draw the request path from HTTP handler to service to repository to PostgreSQL, and state which layer owns validation, transaction boundaries, and SQL.

## 2. Mechanism and vocabulary

**Mechanism.** Each dependency crosses an explicit interface: handler parses/serializes HTTP, service enforces domain rules, repository maps domain operations to parameterized SQL. Context cancellation reaches database calls; transactions group multi-statement state changes.

**Second point.** Integration tests require a known schema and data state. Migrations should run against a clean test database, each test should use isolated identifiers or cleanup, and failure paths must not leave persistent partial state.

Return a stable API error contract: malformed input is a client error; missing resources have a defined not-found response; unexpected database failures are server errors with details logged internally, not leaked to the caller.

**Worked example.** Use the assigned request path as a review checklist: HTTP handler → `TaskService` → `TaskRepository` interface → `PostgresTaskRepository` → `database/sql` pool → PostgreSQL. Keep the service contract independent of SQL driver types.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Check migration from an empty database, CRUD status/body behavior, repository no-row mapping, transaction rollback, and an integration fixture's cleanup. Distinguish tests using a fake repository from tests requiring PostgreSQL.

**Failure mode.** Cause a database constraint violation during a request and verify that it does not become a success response or leak SQL internals. Cancel a request context and verify the repository operation stops or returns an error.

**Completion criterion.** You can rebuild the service from the API contract, trace one request through each layer, verify migrations and persistent state, and explain the test and recovery evidence.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the response for creating a valid task.
2. Predict the response when the title is missing or status is invalid.
3. Predict the result of reading an unknown task ID.
4. Predict which layer maps `sql.ErrNoRows` to an API not-found response.
5. Predict the database state if a multi-statement operation rolls back.
6. Predict how a migration failure should affect startup/deployment.
7. Predict which checks require a real PostgreSQL instance.
8. State what restore evidence is required before claiming the service is recoverable.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Define task schema, migration sequence, API routes, and repository/service boundaries.

2. Implement handlers and a PostgreSQL repository with parameterized SQL and explicit error mapping.

3. Add integration tests for CRUD, invalid/not-found cases, transaction rollback, and fixture isolation.

4. Document the migration, least-privilege role, and a tested restore path.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Force a database constraint failure during a request and inspect API status/body plus stored state.

2. Force a context cancellation or transaction rollback and confirm the API does not claim a committed mutation.

3. Inspect the database as a least-privilege runtime role and distinguish expected denied operations from service defects.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one API or schema constraint, such as adding a due date or making status mandatory. Update migration, validation, repository mapping, and tests while preserving the request path boundary.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What responsibilities belong to handler, service, and repository?
2. How are invalid input and missing resource mapped to HTTP outcomes?
3. Which changes must share one transaction?
4. What does the integration fixture reset or isolate?
5. What evidence shows a migration works from an empty database?
6. What evidence supports the backup/restore claim?

## Evidence to record

Create `lesson057/EVIDENCE.md` with exactly this structure:

```text
# Lesson 057 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Putting SQL or driver types in the HTTP handler.
- Returning success before the database transaction commits.
- Leaking SQL errors or credentials in HTTP responses.
- Running integration tests against uncontrolled developer data.
- Claiming backup safety without restoring and verifying data.

## Precision checks

1. Draw the request path across all layers.
2. Explain the data invariant and database constraint enforcing it.
3. State the API contract for success, invalid input, not-found, and internal failure.
4. Describe deterministic database integration setup and cleanup.
5. List migration and restore evidence needed for operational confidence.

Durable service behavior spans code, schema, and operations; review the invariant and failure handling across all three.
## University revision bridge

For the LIG track, integrate database constraints and transactions with Go repository design, HTTP semantics, and repeatable operational procedures.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 057 — Integrated Lab 2: Durable Data Service without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can write specific, content-conveying subject lines on demand, you've correctly applied formal opening/closing convention to a real first-contact email, and you've correctly reasoned through a real or realistic CC/To decision and can explain the distinction clearly. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
