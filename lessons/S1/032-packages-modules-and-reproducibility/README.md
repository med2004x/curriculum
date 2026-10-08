# Lesson 032 — Packages, Modules, and Reproducible Builds

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand how a non-trivial Go program is divided into packages and how dependencies are declared, resolved, upgraded, and tested.

**The contract.** A package is a compilation and namespace unit; packages in a module share a module path and version/dependency context. An import path names a package. Identifiers beginning with an uppercase letter are exported from their package. Keep dependencies directed so lower-level packages do not import higher-level application packages and create cycles.

**Worked example.** Start with the smallest concrete case described by the source material: Split one previous capstone-sized exercise into two packages with a clear dependency direction.

## 2. Mechanism and vocabulary

**Mechanism.** `go.mod` declares the module path, Go language version, and module requirements. Go's module selection resolves dependency versions according to module rules; `go.sum` records checksums used to authenticate downloaded module content. These files do not capture every toolchain, environment variable, generated file, or external system needed for bit-for-bit reproducibility.

**Second point.** `go test ./...` asks the Go tool to test packages in the current module beneath the current directory. It is not a recursive filesystem search over every nested module, and directory/package exclusions and build constraints still affect what is included.

An import cycle is a package dependency graph cycle and is rejected. An unexported identifier cannot be selected from another package. A testable package boundary makes ownership and dependency direction visible rather than merely moving files into folders.

**Worked example.** Add one external dependency only if the problem genuinely needs it.

## 3. Worked case from the reference material

Worked example: Split one previous capstone-sized exercise into two packages with a clear dependency direction.

## 4. Verification, failure, and evidence

**Verification.** Build/test each package and the module-wide package set. Record the module path, package import paths, exported names, dependency requirements, Go version declaration, and any generated or platform-specific inputs needed to reproduce the build.

**Failure mode.** Create an import cycle, then separately reference an unexported identifier from another package. Explain the distinct compile-time boundary each case violates.

**Completion criterion.** You can explain package/module/import-path distinctions, module metadata, package discovery, and why an import graph should remain acyclic.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
