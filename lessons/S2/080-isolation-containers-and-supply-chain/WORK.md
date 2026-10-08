# Lesson 080 — Isolation, Containers, and Supply Chain Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Inspect process visibility and PID identity inside a container and compare it with the host view.

2. Compare isolated networking with host networking and record visible interfaces/routes.

3. Compare root/default capabilities with a non-root, capability-dropped, read-only configuration; test one operation expected to be denied.

## Run and check

```powershell
Set-Location lessons/S2/080-isolation-containers-and-supply-chain/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Supply-chain boundary

Harden a containerized build against dependency and privilege risks. Record the changed attack surface and verify the build remains reproducible.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
