# Lesson 055 — Go PostgreSQL Repositories and Tests Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Replace a memory repository with PostgreSQL while keeping the service interface unchanged.

2. Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

3. Write a second case with one changed constraint and verify it.

## Run and check

```powershell
Set-Location lessons/S2/055-go-postgresql-repository-and-tests/code
go test -race exercise.go exercise_test.go
```

## Optional challenge

### Go repository layer

Implement a Go repository against PostgreSQL with transaction-aware methods and integration tests. Include one test for a constraint failure.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
