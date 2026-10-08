# Lesson 068 — Handlers, Services, Repositories, and Boundaries

## Threads
**Primary:** Go/Design

**Supporting:** layer boundaries; interfaces; tests

**Math companion:** functions and composition

**Time:** 3–4 hours. **Prerequisites:** Lesson 067.

**Time:** 3–4 hours. **Prerequisites:** Lesson 067.

## What you will be able to do

1. Separate HTTP parsing/response work from application rules and persistence details.
2. Define a service interface around a use case rather than around SQL statements.
3. Map domain errors to HTTP responses at the handler boundary.
4. Test service behavior with a fake repository and persistence behavior with a database integration test.

## Retrieve first (no notes)

1. What does an API contract specify before implementation?
2. Why should a client-facing error not expose an internal SQL message?
3. What is one distinct failure result for a create operation?
4. Why can a status code and response body both be part of the contract?

## 1. Problem and model

**What problem this solves.** Explicit layer boundaries keep transport details, business policy, and storage technology from becoming tangled, while making each contract independently testable.

**The contract.** A handler decodes and validates HTTP-level input, calls a service use case, and maps results/errors to HTTP. A service enforces domain rules and coordinates work. A repository abstracts persistence operations and returns domain-relevant results/errors; it should not decide HTTP status codes.

**Worked example.** `CreateTaskHandler` parses JSON and request context; `TaskService.Create` enforces title/status rules; `TaskRepository.Create` persists the task. The service can be tested with a fake repository that records the call and returns selected domain outcomes.

Self-check: which layer knows about `http.ResponseWriter`? Which layer knows the SQL table name? Keep those dependencies at their boundaries.

## 2. Mechanism and vocabulary

**Mechanism.** Dependencies point inward: handlers depend on service contracts; services depend on repository abstractions; concrete PostgreSQL repositories implement those abstractions. The service should receive its dependencies explicitly rather than constructing a database pool inside a request handler.

**Second point.** Use typed domain errors such as `ErrNotFound`, `ErrConflict`, and `ErrInvalidInput`. The handler maps these to stable HTTP status/body contracts; unknown errors become a generic server response while details are logged internally.

Keep interfaces narrow and owned by the consuming service package where practical. Avoid an interface per struct or a generic repository API that hides the actual use cases.

**Worked example.** Add one boundary case to the # 068 — Handlers, Services, Repositories, and Boundaries exercise.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
type TaskRepository interface {
	Create(context.Context, NewTask) (Task, error)
	Get(context.Context, int64) (Task, error)
}

type TaskService struct {
	repo TaskRepository
}

func (s *TaskService) Create(ctx context.Context, input NewTask) (Task, error) {
	if strings.TrimSpace(input.Title) == "" {
		return Task{}, ErrInvalidInput
	}
	return s.repo.Create(ctx, input)
}
```

The handler maps `ErrInvalidInput` to a client error and `ErrNotFound` to 404. The repository does not know about `http.ResponseWriter`; the service does not need to know SQL table names.

Worked example check: identify which behavior belongs in the service and which result is translated into HTTP by the handler.

Self-check: if the repository returns a connection error, should the handler expose the raw driver message?

## 4. Verification, failure, and evidence

**Verification.** Use a fake repository to test valid create, invalid title, not-found, and repository failure mapping at the service/handler boundary. Use a PostgreSQL integration test for SQL constraints, migrations, and query behavior.

**Failure mode.** Put SQL and HTTP status selection in the service, then explain which dependencies make the service hard to test or reuse. Move each decision to its owning boundary.

**Completion criterion.** You can explain dependency direction, error mapping, and the distinct unit/integration test boundary.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the handler response for `ErrInvalidInput` and `ErrNotFound`.
2. Predict whether a repository should return raw HTTP status codes.
3. Predict which test can run with a fake repository.
4. Predict which behavior requires PostgreSQL integration.
5. Predict where SQL query text belongs.
6. Predict how request context reaches persistence.
7. Predict which layer logs internal database diagnostics.
8. Draw dependency direction before adding an interface.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Define one create-task use case and a narrow repository interface.

2. Implement service validation and a fake repository for isolated tests.

3. Add an HTTP handler that maps typed errors to documented responses; add a separate database-backed repository test.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one storage implementation behind `TaskRepository` or one HTTP transport (for example, JSON handler to another adapter). State which service contract should remain unchanged.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does the handler own?
2. What does the service own?
3. What does the repository own?
4. Why should dependencies point inward?
5. How are typed domain errors mapped to HTTP?
6. Which tests need a real database?

## Evidence to record

Create `lesson068/EVIDENCE.md` with exactly this structure:

```text
# Lesson 068 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Returning HTTP status codes from a repository.
- Embedding SQL or driver types in HTTP handlers.
- Constructing infrastructure dependencies inside the service method.
- Testing SQL constraints only with a fake repository.
- Exposing internal database diagnostics as public API errors.

## Precision checks

1. Draw handler → service → repository dependency flow.
2. State the service invariant and typed errors.
3. Explain the handler's HTTP mapping.
4. Identify fake-backed unit cases and real-database integration cases.
5. Explain how a storage adapter can change without rewriting business rules.

Layer boundaries are useful when each boundary owns a real policy or translation, not merely because a project has three folders.
## University revision bridge

For the LIG track, connect interface boundaries to dependency inversion and API error contracts, then verify the layers at their appropriate test scope.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 068 — Handlers, Services, Repositories, and Boundaries without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can reproduce the main behavior and diagnose one failure from evidence. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
