# Lesson 075 — Integration Testing and Containers Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Write a test for creating and retrieving a task through HTTP against a real PostgreSQL instance.

2. Containerize one backend service and document which state is ephemeral and which state is persistent.

3. Replace the service container and verify persistent database state remains only when the database volume is retained.

## Run and check

```powershell
Set-Location lessons/S2/075-integration-testing-and-containers/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Container integration

Containerize a Go/PostgreSQL service and create integration tests that start dependencies, run migrations, test the service, and clean up deterministically.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
