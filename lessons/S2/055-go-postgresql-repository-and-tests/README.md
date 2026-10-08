# Lesson 055 — Go PostgreSQL Repositories and Tests

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Connect the application layer to a real database without leaking database details through the whole program.

**The contract.** A repository owns persistence operations behind a domain-facing interface. It accepts `context.Context`, uses parameterized SQL, closes `Rows`, checks `Rows.Err()` after iteration, and translates `sql.ErrNoRows` into a deliberate domain result. It should not hide cancellation, constraint, or commit failures as successful zero values.

**Worked example.** Start with the smallest concrete case described by the source material: Replace a memory repository with PostgreSQL while keeping the service interface unchanged.

## 2. Mechanism and vocabulary

**Mechanism.** `database/sql` manages a pool through `*sql.DB`; query methods can use contexts to propagate cancellation. A `*sql.Tx` pins a transaction context for its statements and must be committed or rolled back. Connection configuration and pool sizing belong at the application boundary, not scattered through business logic.

**Second point.** Unit tests can use a fake repository to test service rules; integration tests should exercise SQL, migrations, constraints, and driver behavior against PostgreSQL. Use deterministic fixtures and isolate test data so tests do not depend on run order.

Avoid tests that silently depend on a developer's local database contents. A test must control schema and data setup, use a dedicated database/schema or transaction where appropriate, and reliably clean up even after a failure.

**Worked example.** Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```bash
cd ./16-go-repository
$env:DATABASE_URL = "postgres://task_app:local-app-password@localhost:5432/go_course?sslmode=disable"
go mod tidy
go run .
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** Test repository success, missing-row behavior, constraint violations, canceled context, row scan/iteration failure where practical, and transaction rollback. Record which cases use a fake and which require a real PostgreSQL server.

**Failure mode.** Return `sql.ErrNoRows` directly from the repository and observe how the service/API incorrectly exposes a storage detail; then map it to a domain-level not-found error. Separately cancel a context and ensure the operation does not report success.

**Completion criterion.** You can explain the repository boundary, preserve cancellation and error meaning, and demonstrate repeatable integration setup/cleanup.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
