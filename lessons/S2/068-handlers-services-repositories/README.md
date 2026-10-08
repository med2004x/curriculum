# Lesson 068 — Handlers, Services, Repositories, and Boundaries

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Explicit layer boundaries keep transport details, business policy, and storage technology from becoming tangled, while making each contract independently testable.

**The contract.** A handler decodes and validates HTTP-level input, calls a service use case, and maps results/errors to HTTP. A service enforces domain rules and coordinates work. A repository abstracts persistence operations and returns domain-relevant results/errors; it should not decide HTTP status codes.

**Worked example.** `CreateTaskHandler` parses JSON and request context; `TaskService.Create` enforces title/status rules; `TaskRepository.Create` persists the task. The service can be tested with a fake repository that records the call and returns selected domain outcomes.

## 2. Mechanism and vocabulary

**Mechanism.** Dependencies point inward: handlers depend on service contracts; services depend on repository abstractions; concrete PostgreSQL repositories implement those abstractions. The service should receive its dependencies explicitly rather than constructing a database pool inside a request handler.

**Second point.** Use typed domain errors such as `ErrNotFound`, `ErrConflict`, and `ErrInvalidInput`. The handler maps these to stable HTTP status/body contracts; unknown errors become a generic server response while details are logged internally.

Keep interfaces narrow and owned by the consuming service package where practical. Avoid an interface per struct or a generic repository API that hides the actual use cases.

**Worked example.** Add one boundary case to the # 068 — Handlers, Services, Repositories, and Boundaries exercise.

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

## 4. Verification, failure, and evidence

**Verification.** Use a fake repository to test valid create, invalid title, not-found, and repository failure mapping at the service/handler boundary. Use a PostgreSQL integration test for SQL constraints, migrations, and query behavior.

**Failure mode.** Put SQL and HTTP status selection in the service, then explain which dependencies make the service hard to test or reuse. Move each decision to its owning boundary.

**Completion criterion.** You can explain dependency direction, error mapping, and the distinct unit/integration test boundary.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
