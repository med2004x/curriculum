# Lesson 099 — Real Repositories, Bug Reproduction, and Upstreaming

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build an accurate mental model of a large, unfamiliar codebase before making any changes to it — a distinct, learnable skill from writing new code, and one this entire curriculum has been implicitly building toward through every "read before modifying" instruction in earlier tracks (OS Lesson 8's xv6 labs, database-internals' BusTub capstone, and others).

**The contract.** A useful patch starts with a reproducible difference between expected and actual behavior. Tests expose existing contracts; entry points and history reveal ownership and constraints that directory names alone cannot establish.

**Worked example.** Start with the smallest concrete case described by the source material: Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before.

## 2. Mechanism and vocabulary

**Mechanism.** Reading someone else's large, mature codebase requires reconstructing that understanding *after the fact*, from evidence (the code itself, tests, commit history, documentation) rather than lived decision-making — a skill that doesn't automatically transfer from being good at writing code, and one that's genuinely necessary for almost any real engineering job.

**Second point.** **Entry points: where to actually start reading.** A large codebase has no single obvious starting point, but it does have identifiable entry points — the `main()` function or equivalent, the top-level HTTP route registrations, the primary public API surface.

**Worked example.** A project related to something covered elsewhere in this curriculum (a database engine, an HTTP framework, a CLI tool you use) is a reasonable, motivated choice.

## 3. Worked case from the reference material

Worked example: Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before.

## 4. Verification, failure, and evidence

**Verification.** Save a one-page map of modules, dependencies, entry points, and one relevant historical decision. Include the minimal pre-fix reproduction, the regression test, focused and neighboring test results, and the exact revision used. Separate facts observed in the repository from assumptions in the map.

**Failure mode.** First write a ten-minute sketch based only on README and file names. After tracing code, tests, and history, identify one specific incorrect or incomplete assumption in that sketch and cite the evidence that changed it.

**Completion criterion.** You can trace an operation into its implementation and tests, preserve a pre-fix failure, and submit a narrow patch whose regression test passes afterward.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
