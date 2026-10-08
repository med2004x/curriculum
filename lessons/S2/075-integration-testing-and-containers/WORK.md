# Lesson 075 — Problem Set

**Topic:** Integration Testing and Containers.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

2. Extend it with one explicit boundary case: Containerize one backend service and document which state is ephemeral and which state is persistent.

3. Replace the service container and verify persistent database state remains only when the database volume is retained.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/075-integration-testing-and-containers/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Container integration

Containerize a Go/PostgreSQL service and create integration tests that start dependencies, run migrations, test the service, and clean up deterministically.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
