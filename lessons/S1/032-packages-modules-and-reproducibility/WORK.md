# Lesson 032 — Problem Set

**Topic:** Packages, Modules, and Reproducible Builds.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Split a small program into an application package and a lower-level parsing or data package.

2. Export only the identifiers required by the caller and keep the package dependency graph acyclic.

3. Declare the module path and required Go version, inspect dependency metadata, and explain how module-wide tests enumerate packages.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/032-packages-modules-and-reproducibility/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Reproducible module

Create a small Go module with two packages and deterministic tests. Pin dependencies, run from a clean directory, and record the exact build/test commands.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
