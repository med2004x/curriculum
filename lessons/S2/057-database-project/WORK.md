# Lesson 057 — Problem Set

**Topic:** Integrated Lab 2: Durable Data Service.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define task schema, migration sequence, API routes, and repository/service boundaries.

2. Implement handlers and a PostgreSQL repository with parameterized SQL and explicit error mapping.

3. Add integration tests for CRUD, invalid/not-found cases, transaction rollback, and fixture isolation.

4. Document the migration, least-privilege role, and a tested restore path.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/057-database-project/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Durable service lab

Build a small Go data service backed by PostgreSQL with migrations, transactions, tests, and a reproducible local setup. Diagnose one intentional failure.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
