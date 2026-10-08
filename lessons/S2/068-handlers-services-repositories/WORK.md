# Lesson 068 — Handlers, Services, Repositories, and Boundaries Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Define one create-task use case and a narrow repository interface.

2. Implement service validation and a fake repository for isolated tests.

3. Add an HTTP handler that maps typed errors to documented responses; add a separate database-backed repository test.

## Run and check

```powershell
Set-Location lessons/S2/068-handlers-services-repositories/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Service boundaries

Split a backend feature into handler, service, and repository layers. Test each layer at its appropriate scope and document one boundary that prevents coupling.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
