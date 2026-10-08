# Lesson 055 — Go PostgreSQL Repositories and Tests

## Threads
**Primary:** Go/DB/Testing

**Supporting:** driver boundaries; repository code; integration tests

**Math companion:** sampling and test evidence

**Time:** 3–4 hours. **Prerequisites:** Lesson 054.

**Time:** 3–4 hours. **Prerequisites:** Lesson 054.

## What you will be able to do

1. Implement a Go repository using `database/sql` and a PostgreSQL driver.
2. Pass request contexts to queries and map database outcomes to explicit domain errors.
3. Handle `sql.ErrNoRows`, row iteration errors, and transaction failures without leaking database details.
4. Design integration fixtures that isolate state and clean up reliably.

## Retrieve first (no notes)

1. How does MVCC affect transaction visibility?
2. What does `SELECT ... FOR UPDATE` lock?
3. Why should multiple rows be locked in a consistent order?
4. What should the application do after a serialization failure?

## 1. Problem and model

**What problem this solves.** Connect the application layer to a real database without leaking database details through the whole program.

**The contract.** A repository owns persistence operations behind a domain-facing interface. It accepts `context.Context`, uses parameterized SQL, closes `Rows`, checks `Rows.Err()` after iteration, and translates `sql.ErrNoRows` into a deliberate domain result. It should not hide cancellation, constraint, or commit failures as successful zero values.

**Worked example.** Start with the smallest concrete case described by the source material: Replace a memory repository with PostgreSQL while keeping the service interface unchanged.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** `database/sql` manages a pool through `*sql.DB`; query methods can use contexts to propagate cancellation. A `*sql.Tx` pins a transaction context for its statements and must be committed or rolled back. Connection configuration and pool sizing belong at the application boundary, not scattered through business logic.

**Second point.** Unit tests can use a fake repository to test service rules; integration tests should exercise SQL, migrations, constraints, and driver behavior against PostgreSQL. Use deterministic fixtures and isolate test data so tests do not depend on run order.

Avoid tests that silently depend on a developer's local database contents. A test must control schema and data setup, use a dedicated database/schema or transaction where appropriate, and reliably clean up even after a failure.

**Worked example.** Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

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

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Test repository success, missing-row behavior, constraint violations, canceled context, row scan/iteration failure where practical, and transaction rollback. Record which cases use a fake and which require a real PostgreSQL server.

**Failure mode.** Return `sql.ErrNoRows` directly from the repository and observe how the service/API incorrectly exposes a storage detail; then map it to a domain-level not-found error. Separately cancel a context and ensure the operation does not report success.

**Completion criterion.** You can explain the repository boundary, preserve cancellation and error meaning, and demonstrate repeatable integration setup/cleanup.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the repository result when an ID has no matching row.
2. Predict the domain/API result after mapping `sql.ErrNoRows`.
3. Predict what cancellation should do to a context-aware query.
4. Predict what happens if `Rows` is not closed.
5. Predict why `Rows.Err()` must be checked after iteration.
6. Predict which state is visible after transaction rollback.
7. Distinguish a unit test with a fake from an integration test with PostgreSQL.
8. State how test fixtures remain independent of order.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Replace a memory repository with PostgreSQL while keeping the service interface unchanged.

2. Extend it with one explicit boundary case: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

3. Rebuild the same behavior with one changed constraint: Write a second case with one changed constraint and verify it.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Handle connection failure, query cancellation, no-row results, and a transaction rollback.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one storage implementation behind the repository interface, or replace one fake-backed service test with a PostgreSQL integration test. Explain which contract remains stable and which behavior now requires the database.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What belongs behind a repository interface?
2. Why pass a context to database operations?
3. How should a missing row differ from a query failure?
4. Why close `Rows` and check `Rows.Err()`?
5. What does a transaction boundary guarantee for grouped writes?
6. How should integration tests isolate fixtures and cleanup?

## Evidence to record

Create `lesson055/EVIDENCE.md` with exactly this structure:

```text
# Lesson 055 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Returning `sql.ErrNoRows` as an opaque storage detail to the API layer.
- Ignoring a context cancellation or commit error.
- Forgetting `Rows.Close()` or `Rows.Err()`.
- Sharing mutable test fixtures across tests that may run independently.
- Calling a fake-backed test an integration test.

## Precision checks

1. State the repository's domain-facing contract.
2. Map no-row, canceled-context, and constraint errors explicitly.
3. Explain row iteration cleanup and terminal error checking.
4. Distinguish fake-based unit tests from database integration tests.
5. Describe fixture setup and cleanup that remain safe on failure.

The repository boundary isolates SQL details, but it must preserve meaningful error, cancellation, and transaction outcomes.
## University revision bridge

For the LIG track, connect repository abstractions to dependency inversion and database transaction semantics, then test the boundary at both unit and integration levels.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 055 — Go PostgreSQL Repositories and Tests without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can swap storage implementations without rewriting business rules. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
