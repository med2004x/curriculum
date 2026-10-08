# Lesson 068 — Problem Set

**Topic:** Handlers, Services, Repositories, and Boundaries.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define one create-task use case and a narrow repository interface.

2. Implement service validation and a fake repository for isolated tests.

3. Add an HTTP handler that maps typed errors to documented responses; add a separate database-backed repository test.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/068-handlers-services-repositories/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Service boundaries

Split a backend feature into handler, service, and repository layers. Test each layer at its appropriate scope and document one boundary that prevents coupling.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
