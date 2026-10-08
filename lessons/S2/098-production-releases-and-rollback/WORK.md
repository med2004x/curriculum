# Lesson 098 — Problem Set

**Topic:** Production Releases, CI/CD, Containers, and Rollback.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Create a pipeline that runs formatting checks, tests, static analysis, a build, and an artifact verification step.

2. Pin dependencies and toolchain inputs; build once and record the artifact digest so later promotion can verify artifact identity.

3. Rebuild the same behavior with one changed constraint: GitHub Actions, or a local script if CI isn't set up yet) that: checks out a specific Git commit, builds the artifact (a Go binary, a Docker image, or equivalent), and embeds the commit hash and build timestamp into the artifact itself (e.g. via a version endpoint the running application exposes, or embedded build-time constants).

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/098-production-releases-and-rollback/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Release pipeline

Create a reproducible build and deployment pipeline with versioned artifacts, health checks, rollback, and a documented failed deployment drill.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
