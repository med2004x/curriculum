# Lesson 057 — Integrated Lab 2: Durable Data Service Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Define task schema, migration sequence, API routes, and repository/service boundaries.

2. Implement handlers and a PostgreSQL repository with parameterized SQL and explicit error mapping.

3. Add integration tests for CRUD, invalid/not-found cases, transaction rollback, and fixture isolation.

4. Document the migration, least-privilege role, and a tested restore path.

## Run and check

```powershell
Set-Location lessons/S2/057-database-project/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Durable service lab

Build a small Go data service backed by PostgreSQL with migrations, transactions, tests, and a reproducible local setup. Diagnose one intentional failure.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
