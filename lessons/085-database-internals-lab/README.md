# Lesson 085 — Database Internals Lab

## Threads
**Primary:** DB Internals/C/Algorithms

**Supporting:** implement or instrument a DB component

**Math companion:** measurement

**Time:** 3–4 hours. **Prerequisites:** Lesson 084.

**Time:** 3–4 hours. **Prerequisites:** Lesson 084.

## What you will be able to do

1. Select one database component boundary: page storage, buffer pool, B+ tree, iterator execution, or WAL/recovery.
2. State invariants and interfaces before implementing or instrumenting that component.
3. Integrate it with at least one neighboring component and verify their contract together.
4. Measure one workload and explain the observed bottleneck without generalizing beyond the evidence.

## Retrieve first (no notes)

1. What must be durable before a database acknowledges commit under its WAL policy?
2. What does a checkpoint do?
3. What state can crash recovery reconstruct from the log?
4. How does a serialization failure differ from recovery after a crash?

## 1. Problem and model

**What problem this solves.** Integrate Lessons 1-6 into one substantial, tested database component with documented invariants and measured performance — the database-internals track's equivalent of the integration capstones in earlier tracks.

**The contract.** Choose one component with an observable invariant. For example, a buffer pool must not evict a pinned frame; a B+ tree must route searches according to separators and maintain equal leaf depth; an iterator must emit the same rows as its logical query; a recovery component must preserve acknowledged committed state under its stated assumptions.

**Worked example.** Start with the smallest concrete case described by the source material: If using BusTub directly, complete at least the buffer pool manager project and the B+ tree index project from the official course materials (`https://github.com/cmu-db/bustub`), using its existing test infrastructure to verify correctness — this is the recommended default path, since BusTub's tests are considerably more rigorous than what's practical to con.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Use the repository's chosen teaching system or a bounded component built from prior lessons. Keep the component boundary explicit: page format, buffer-pool lifecycle, index invariant, execution operator, or recovery record. Do not claim a full DBMS implementation when only one component is implemented.

**Second point.** Integration risk lives at assumptions between components: page IDs versus frame IDs, pin lifetime versus eviction, index keys versus table records, iterator end-of-stream versus errors, and log sequence numbers versus persistence order.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** If building your own integrated system instead (a reasonable alternative if you want the components to be ones you fully wrote yourself, end to end), integrate your own Lessons 1-6 implementations into one coherent system with this minimum scope:

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: If using BusTub directly, complete at least the buffer pool manager project and the B+ tree index project from the official course materials (`https://github.com/cmu-db/bustub`), using its existing test infrastructure to verify correctness — this is the recommended default path, since BusTub's tests are considerably more rigorous than what's practical to con.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record the selected invariant, the cross-component test that checks it, the exact test scope, and a reproducible measurement with workload size and environment. A passing test covers its cases, not all possible states.

**Failure mode.** Break one interface assumption in a controlled branch—for example, unpin a page before its consumer finishes using the frame, or return an iterator row after end-of-stream—and show which invariant/test fails.

**Completion criterion.** You can explain one integrated component path, state its invariants, show a cross-component test and a deliberate failure, and support one performance claim with measured evidence.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. State the chosen component's invariant.
2. Predict what fails if a pinned frame is evicted.
3. Predict the observable behavior of an iterator after exhaustion.
4. Predict which component owns page serialization and which owns frame lifetime.
5. Predict how the test verifies a split or recovery boundary.
6. Predict one measurement affected by the chosen workload.
7. Distinguish a correctness test from a performance experiment.
8. State the environment assumptions before interpreting results.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: If using BusTub directly, complete at least the buffer pool manager project and the B+ tree index project from the official course materials (`https://github.com/cmu-db/bustub`), using its existing test infrastructure to verify correctness — this is the recommended default path, since BusTub's tests are considerably more rigorous than what's practical to con.

2. Extend it with one explicit boundary case: If building your own integrated system instead (a reasonable alternative if you want the components to be ones you fully wrote yourself, end to end), integrate your own Lessons 1-6 implementations into one coherent system with this minimum scope:

3. Rebuild the same behavior with one changed constraint: **Storage layer**: your Lesson 1 page format, with records actually read/written through your Lesson 2 buffer pool (not bypassing it) — every page access in the rest of the system should go through `fetchPage`/`unpinPage`, with correct pin discipline throughout.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take your integrated system (or the BusTub project you completed) and deliberately break one cross-component assumption — for example, if you built your own system, bypass the buffer pool's pin discipline in exactly one code path (directly access a page without pinning it first) while the rest of the system correctly uses pins.

2. Run your workload under concurrent access (multiple goroutines/threads performing operations simultaneously) and attempt to reproduce a corruption or crash caused specifically by this one broken assumption — e.g. a page being evicted while the unprotected code path is still reading it, per Lesson 2's pin/unpin discussion.

3. If you cannot reliably reproduce a visible failure (this specific race can be timing-dependent and hard to force), document your reasoning for why the assumption violation is still a genuine bug even if your specific test run didn't happen to trigger a visible symptom — directly echoing C track Lesson 5's core point about undefined behavior not requiring a v.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 085 — Database Internals Lab: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 085 — Database Internals Lab?
2. What two concrete practices are required in the assigned material for # 085 — Database Internals Lab?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson085/EVIDENCE.md` with exactly this structure:

```text
# Lesson 085 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 085 — Database Internals Lab.
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
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 085 — Database Internals Lab without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've completed a genuinely integrated system (or the equivalent BusTub projects) where storage, buffer pool, indexing, execution, and durability components actually depend on and correctly use each other — not separate, disconnected pieces — your cross-component invariant tests pass consistently, and you've profiled a real bottleneck and connected at least. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
