# Lesson 099 — Real Repositories, Bug Reproduction, and Upstreaming

## Threads
**Primary:** Open Source/Systems/Communication

**Supporting:** navigate unfamiliar code; reproduce; patch; review

**Math companion:** experimental evidence

**Time:** 3–4 hours. **Prerequisites:** Lesson 098.


## What you will be able to do

1. Map an unfamiliar repository using entry points, tests, and history.
2. Reduce a reported defect to a minimal reproducible case.
3. Add a regression test and make the smallest responsible patch.
4. Prepare a reviewable contribution with reproducible evidence.

## Retrieve first (no notes)

1. How can you trace an artifact to source?
2. Why do rolling versions need compatibility?
3. What signal halts a canary?
4. How can schema changes block rollback?

## 1. Problem and model

**What problem this solves.** Build an accurate mental model of a large, unfamiliar codebase before making any changes to it — a distinct, learnable skill from writing new code, and one this entire curriculum has been implicitly building toward through every "read before modifying" instruction in earlier tracks (OS Lesson 8's xv6 labs, database-internals' BusTub capstone, and others).

**The contract.** A useful patch starts with a reproducible difference between expected and actual behavior. Tests expose existing contracts; entry points and history reveal ownership and constraints that directory names alone cannot establish.

**Worked example.** Start with the smallest concrete case described by the source material: Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Reading someone else's large, mature codebase requires reconstructing that understanding *after the fact*, from evidence (the code itself, tests, commit history, documentation) rather than lived decision-making — a skill that doesn't automatically transfer from being good at writing code, and one that's genuinely necessary for almost any real engineering job.

**Second point.** **Entry points: where to actually start reading.** A large codebase has no single obvious starting point, but it does have identifiable entry points — the `main()` function or equivalent, the top-level HTTP route registrations, the primary public API surface.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** A project related to something covered elsewhere in this curriculum (a database engine, an HTTP framework, a CLI tool you use) is a reasonable, motivated choice.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Save a one-page map of modules, dependencies, entry points, and one relevant historical decision. Include the minimal pre-fix reproduction, the regression test, focused and neighboring test results, and the exact revision used. Separate facts observed in the repository from assumptions in the map.

**Failure mode.** First write a ten-minute sketch based only on README and file names. After tracing code, tests, and history, identify one specific incorrect or incomplete assumption in that sketch and cite the evidence that changed it.

**Completion criterion.** You can trace an operation into its implementation and tests, preserve a pre-fix failure, and submit a narrow patch whose regression test passes afterward.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 099 — Real Repositories, Bug Reproduction, and Upstreaming case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before.

2. Extend it with one explicit boundary case: A project related to something covered elsewhere in this curriculum (a database engine, an HTTP framework, a CLI tool you use) is a reasonable, motivated choice.

3. Rebuild the same behavior with one changed constraint: Identify and trace the project's entry point(s): find `main()` (or equivalent), and trace the first few layers of what it calls, building an initial map of the major subsystems/modules it wires together.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

Write a ten-minute first-impression architecture sketch. Then inspect entry points, tests, and history, identify a missed assumption, and compare the evidence-backed map with the sketch.

Save the baseline, command, observed failure, and repaired result. Confirm the original success case still passes after the repair.

## Transfer

Change exactly one constraint in the main case for # 099 — Real Repositories, Bug Reproduction, and Upstreaming: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Which tests define the affected public behavior?
2. What is the smallest input that reproduces the defect?
3. Which command proves the pre-fix failure?
4. Which boundary case shares the code path?
5. What can history reveal about a constraint?
6. What checks follow the focused regression test?

## Evidence to record

Create `lesson099/EVIDENCE.md` with exactly this structure:

```text
# Lesson 099 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 099 — Real Repositories, Bug Reproduction, and Upstreaming.
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

1. Record repository URL and exact revision.
2. Save the initial README-and-tree impression.
3. Trace one entry point to implementation and tests.
4. Save the pre-fix reproduction command and output.
5. Add one regression assertion for expected behavior.
6. Run focused and neighboring tests after the patch.
7. Inspect the final diff for unrelated edits.

## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the lesson contract, rebuild its core case from a clean workspace, reproduce and repair its failure drill, solve the transfer constraint, and meet this evidence bar: You can trace an operation into its implementation and tests, preserve a pre-fix failure, and submit a narrow patch whose regression test passes afterward. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
