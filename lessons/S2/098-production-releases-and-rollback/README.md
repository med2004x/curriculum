# Lesson 098 — Production Releases, CI/CD, Containers, and Rollback

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Make correctness checks reproducible outside your laptop.

**The contract.** Cover CI stages, dependency caching, lint/test/build separation, artifact generation, secret injection, and promotion between environments.

**Worked example.** For a small service, make the pipeline run formatting, tests, static analysis, build, artifact checksum, and a version endpoint. Promote the same artifact to staging and production; use a canary health threshold and a documented rollback command.

## 2. Mechanism and vocabulary

**Mechanism.** Non-reproducibility can come from unpinned dependencies, embedded timestamps, or nondeterministic build steps. Record the source revision, toolchain, dependency lock, and artifact digest so an operator can identify exactly what is running.

**Second point.** Non-reproducibility usually stems from unpinned dependencies (security-engineering Lesson 6's exact concern, now applied to build-time rather than runtime risk), embedded timestamps, or non-deterministic build steps — and it matters because a non-reproducible build makes it genuinely hard to verify "the artifact currently running in production actually corre.

**Worked example.** Build a release pipeline that formats, tests, runs static analysis, creates one artifact, and records its source revision and checksum. Promote that artifact to staging and production without rebuilding it.

## 3. Worked case from the reference material

Worked example: Create a pipeline that runs formatting checks, tests, static analysis, a build, and an artifact verification step.

## 4. Verification, failure, and evidence

**Verification.** Show your pipeline's actual configuration/script, your step 2 reproducibility comparison (with any legitimate, expected differences explicitly noted and explained), and your step 3-4 rollback demonstration with the version endpoint correctly confirming which specific build is actually running after the rollback.

**Failure mode.** Deliberately leave one dependency unpinned or embed a build timestamp, then compare two builds of the same revision. Record which inputs or digests differ and explain why a commit alone does not identify the deployed artifact.

**Completion criterion.** You can reproduce a build from pinned inputs, trace the artifact through promotion, state the canary rollback trigger, and show that the prior version remains compatible with the migrated schema.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
