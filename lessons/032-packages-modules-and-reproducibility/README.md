# Lesson 032 — Packages, Modules, and Reproducible Builds

## Threads
**Primary:** Go/Linux

**Supporting:** package boundaries; modules; dependency hygiene

**Math companion:** sets and modular decomposition

**Time:** 3–4 hours. **Prerequisites:** Lesson 031.

**Time:** 3–4 hours. **Prerequisites:** Lesson 031.

## What you will be able to do

1. Distinguish a Go package from a Go module.
2. Split a small program into packages with a one-way dependency direction and an explicit exported API.
3. Explain the roles of `go.mod` and `go.sum` without treating either as a complete build lockfile.
4. Describe what `go test ./...` discovers and what inputs a reproducible build requires.

## Retrieve first (no notes)

1. What is a wait-for graph?
2. Which Coffman condition does a global lock order prevent?
3. Why should concurrent tests coordinate on events instead of guessed sleep times?
4. How do safety failures differ from liveness failures?

## 1. Problem and model

**What problem this solves.** Understand how a non-trivial Go program is divided into packages and how dependencies are declared, resolved, upgraded, and tested.

**The contract.** A package is a compilation and namespace unit; packages in a module share a module path and version/dependency context. An import path names a package. Identifiers beginning with an uppercase letter are exported from their package. Keep dependencies directed so lower-level packages do not import higher-level application packages and create cycles.

**Worked example.** Start with the smallest concrete case described by the source material: Split one previous capstone-sized exercise into two packages with a clear dependency direction.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** `go.mod` declares the module path, Go language version, and module requirements. Go's module selection resolves dependency versions according to module rules; `go.sum` records checksums used to authenticate downloaded module content. These files do not capture every toolchain, environment variable, generated file, or external system needed for bit-for-bit reproducibility.

**Second point.** `go test ./...` asks the Go tool to test packages in the current module beneath the current directory. It is not a recursive filesystem search over every nested module, and directory/package exclusions and build constraints still affect what is included.

An import cycle is a package dependency graph cycle and is rejected. An unexported identifier cannot be selected from another package. A testable package boundary makes ownership and dependency direction visible rather than merely moving files into folders.

**Worked example.** Add one external dependency only if the problem genuinely needs it.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Split one previous capstone-sized exercise into two packages with a clear dependency direction.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Build/test each package and the module-wide package set. Record the module path, package import paths, exported names, dependency requirements, Go version declaration, and any generated or platform-specific inputs needed to reproduce the build.

**Failure mode.** Create an import cycle, then separately reference an unexported identifier from another package. Explain the distinct compile-time boundary each case violates.

**Completion criterion.** You can explain package/module/import-path distinctions, module metadata, package discovery, and why an import graph should remain acyclic.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the import path for a package beneath the module root.
2. Predict what happens when package A imports B and B imports A.
3. Predict whether a lowercase package-level function can be called from another package.
4. Predict which project metadata belongs in `go.mod` and which checksums belong in `go.sum`.
5. Predict which packages `go test ./...` targets from the module root.
6. Predict what may differ when the Go toolchain version changes.
7. Predict how a second nested module affects a module-root test pattern.
8. State one environment input not fully captured by module metadata.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Split a small program into an application package and a lower-level parsing or data package.

2. Export only the identifiers required by the caller and keep the package dependency graph acyclic.

3. Declare the module path and required Go version, inspect dependency metadata, and explain how module-wide tests enumerate packages.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Make the lower-level package import the application package and observe why the cycle is invalid.

2. Refer to a lowercase identifier across the package boundary and explain why it is unavailable.

3. Change the module path in `go.mod` without updating imports and describe the resulting resolution failure.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one boundary: move a package into a subdirectory or extract it into a separate module. Update imports and explain which metadata and test commands now apply.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the difference between a package and a module?
2. What does an import path identify?
3. Which identifiers are exported from a Go package?
4. What information is represented by `go.mod` and `go.sum`?
5. Why is an import cycle invalid?
6. What scope does `go test ./...` target?

## Evidence to record

Create `lesson032/EVIDENCE.md` with exactly this structure:

```text
# Lesson 032 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating a package directory and a module as synonyms.
- Assuming a `go.sum` entry declares a direct dependency.
- Assuming module files capture the full compiler and host environment.
- Creating reverse dependencies that produce an import cycle.
- Assuming `go test ./...` crosses into every nested module.

## Precision checks

1. Draw the package import graph and verify it is acyclic.
2. Identify the module root and module path.
3. Explain the role of each dependency-related metadata file.
4. State which package tree a test pattern targets.
5. List the remaining toolchain/environment assumptions for a reproducible build.

Reproducibility is a declared set of inputs, not a property guaranteed by the presence of `go.mod` alone.
## University revision bridge

For the LIG track, connect package boundaries to modular decomposition and the dependency graph, while keeping Go module versioning distinct from package-level visibility.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can split a small program into packages, explain exported visibility and module metadata, diagnose an import cycle, and state the scope and limitations of `go test ./...` and module-based reproducibility.
