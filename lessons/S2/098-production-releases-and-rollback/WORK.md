# Lesson 098 — Production Releases, CI/CD, Containers, and Rollback Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Create a pipeline that runs formatting checks, tests, static analysis, a build, and an artifact verification step.

2. Pin dependencies and toolchain inputs; build once and record the artifact digest so later promotion can verify artifact identity.

3. GitHub Actions, or a local script if CI isn't set up yet) that: checks out a specific Git commit, builds the artifact (a Go binary, a Docker image, or equivalent), and embeds the commit hash and build timestamp into the artifact itself (e.g. via a version endpoint the running application exposes, or embedded build-time constants).

## Run and check

```powershell
Set-Location lessons/S2/098-production-releases-and-rollback/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Release pipeline

Create a reproducible build and deployment pipeline with versioned artifacts, health checks, rollback, and a documented failed deployment drill.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
