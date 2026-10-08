# Lesson 055 — Problem Set

**Topic:** Go PostgreSQL Repositories and Tests.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Replace a memory repository with PostgreSQL while keeping the service interface unchanged.

2. Extend it with one explicit boundary case: Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

3. Rebuild the same behavior with one changed constraint: Write a second case with one changed constraint and verify it.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/055-go-postgresql-repository-and-tests/code
go test -race exercise.go exercise_test.go
```

## Stretch problem

### Go repository layer

Implement a Go repository against PostgreSQL with transaction-aware methods and integration tests. Include one test for a constraint failure.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
