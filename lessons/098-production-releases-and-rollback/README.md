# Lesson 098 — Production Releases, CI/CD, Containers, and Rollback

## Threads
**Primary:** Production/Go/Linux

**Supporting:** pipelines; versioning; migrations; rollback

**Math companion:** risk and probability

**Time:** 3–4 hours. **Prerequisites:** Lesson 097.


## What you will be able to do

1. Create a reproducible CI path from source to release artifact.
2. Trace a deployed artifact to its revision and build inputs.
3. Plan a backward-compatible schema migration and rollout.
4. Practice rollback using health signals and explicit stop criteria.

## Retrieve first (no notes)

1. Which rule prevents mutable aliasing in Rust?
2. Why does compilation not prove behavior?
3. What does owned data provide over a borrow?
4. Why preserve a baseline before changing code?

## 1. Problem and model

**What problem this solves.** Make correctness checks reproducible outside your laptop.

**The contract.** Cover CI stages, dependency caching, lint/test/build separation, artifact generation, secret injection, and promotion between environments.

**Worked example.** For a small service, make the pipeline run formatting, tests, static analysis, build, artifact checksum, and a version endpoint. Promote the same artifact to staging and production; use a canary health threshold and a documented rollback command.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Non-reproducibility can come from unpinned dependencies, embedded timestamps, or nondeterministic build steps. Record the source revision, toolchain, dependency lock, and artifact digest so an operator can identify exactly what is running.

**Second point.** Non-reproducibility usually stems from unpinned dependencies (security-engineering Lesson 6's exact concern, now applied to build-time rather than runtime risk), embedded timestamps, or non-deterministic build steps — and it matters because a non-reproducible build makes it genuinely hard to verify "the artifact currently running in production actually corre.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Build a release pipeline that formats, tests, runs static analysis, creates one artifact, and records its source revision and checksum. Promote that artifact to staging and production without rebuilding it.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Create a pipeline that runs formatting checks, tests, static analysis, a build, and an artifact verification step.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Show your pipeline's actual configuration/script, your step 2 reproducibility comparison (with any legitimate, expected differences explicitly noted and explained), and your step 3-4 rollback demonstration with the version endpoint correctly confirming which specific build is actually running after the rollback.

**Failure mode.** Deliberately leave one dependency unpinned or embed a build timestamp, then compare two builds of the same revision. Record which inputs or digests differ and explain why a commit alone does not identify the deployed artifact.

**Completion criterion.** You can reproduce a build from pinned inputs, trace the artifact through promotion, state the canary rollback trigger, and show that the prior version remains compatible with the migrated schema.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 098 — Production Releases, CI/CD, Containers, and Rollback case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Create a pipeline that runs formatting checks, tests, static analysis, a build, and an artifact verification step.

2. Pin dependencies and toolchain inputs; build once and record the artifact digest so later promotion can verify artifact identity.

3. Rebuild the same behavior with one changed constraint: GitHub Actions, or a local script if CI isn't set up yet) that: checks out a specific Git commit, builds the artifact (a Go binary, a Docker image, or equivalent), and embeds the commit hash and build timestamp into the artifact itself (e.g. via a version endpoint the running application exposes, or embedded build-time constants).

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

Unpin one dependency and compare resolved inputs or artifact digests. Simulate a failing canary and verify promotion stops; retain the expanded schema while restoring the prior application.

Save the baseline, command, observed failure, and repaired result. Confirm the original success case still passes after the repair.

## Transfer

Change exactly one constraint in the main case for # 098 — Production Releases, CI/CD, Containers, and Rollback: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Why promote one artifact across environments?
2. What metadata identifies the deployed artifact?
3. What must overlap versions support?
4. Why can a migration block code rollback?
5. What signal should halt a canary?
6. How does rollback differ from roll-forward?

## Evidence to record

Create `lesson098/EVIDENCE.md` with exactly this structure:

```text
# Lesson 098 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 098 — Production Releases, CI/CD, Containers, and Rollback.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## Lab checklist

Use this checklist to decide whether the exercise is reproducible and reviewable:

1. Record the source revision and toolchain.
2. Pin dependencies and retain the lockfile.
3. Save the artifact digest from the build.
4. Promote the same artifact to each environment.
5. Write down the canary health threshold.
6. Keep schema compatible with old and new binaries.
7. Verify service health after rollback.

## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the lesson contract, rebuild its core case from a clean workspace, reproduce and repair its failure drill, solve the transfer constraint, and meet this evidence bar: You can reproduce a build from pinned inputs, trace the artifact through promotion, state the canary rollback trigger, and show that the prior version remains compatible with the migrated schema. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
