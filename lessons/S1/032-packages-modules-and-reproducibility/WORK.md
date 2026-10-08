# Lesson 032 — Packages, Modules, and Reproducible Builds Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Split a small program into an application package and a lower-level parsing or data package.

2. Export only the identifiers required by the caller and keep the package dependency graph acyclic.

3. Declare the module path and required Go version, inspect dependency metadata, and explain how module-wide tests enumerate packages.

## Run and check

```powershell
Set-Location lessons/S1/032-packages-modules-and-reproducibility/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Reproducible module

Create a small Go module with two packages and deterministic tests. Pin dependencies, run from a clean directory, and record the exact build/test commands.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
