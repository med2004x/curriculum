# Lesson 080 — Problem Set

**Topic:** Isolation, Containers, and Supply Chain.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Inspect process visibility and PID identity inside a container and compare it with the host view.

2. Compare isolated networking with host networking and record visible interfaces/routes.

3. Compare root/default capabilities with a non-root, capability-dropped, read-only configuration; test one operation expected to be denied.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/080-isolation-containers-and-supply-chain/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Supply-chain boundary

Harden a containerized build against dependency and privilege risks. Record the changed attack surface and verify the build remains reproducible.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
