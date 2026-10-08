# Lesson 083 — Query Execution and Cost Models

## Threads
**Primary:** DB Internals/Algorithms

**Supporting:** operators; plans; selectivity; joins

**Math companion:** statistics

**Time:** 3–4 hours. **Prerequisites:** Lesson 082.

**Time:** 3–4 hours. **Prerequisites:** Lesson 082.

## What you will be able to do

1. Compose scan, filter, and projection operators using a pull-based iterator interface.
2. Estimate operator output cardinality and explain how selectivity affects plan cost.
3. Read a PostgreSQL `EXPLAIN` tree and distinguish estimates from measured execution.
4. Explain why index scans can lose to sequential scans for nonselective queries.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 082, # Lesson 082 — B+ Trees and Index Structures.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Build a minimal query execution pipeline (scan → filter → project) using the iterator pattern real database engines use, and learn to read a real `EXPLAIN` plan as a description of exactly this kind of operator tree.

**The contract.** **The iterator model (Volcano-style execution, the standard architecture).** A query plan is a tree of operators (scan, filter, project, join, sort, aggregate...), and each operator implements a uniform interface: typically `open()`, `next()` (returns the next row, or signals "no more rows"), and `close()`.

**Worked example.** Implement `SeqScan`, `Filter`, and `Project` operators. For a fixed input, record rows consumed by the scan, rows emitted by the filter, and columns emitted by projection. These counts make the pipeline's cardinality changes explicit.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** In a pull iterator, a parent requests the next row from its child; filters may consume several child rows before producing one output. Blocking operators such as sort or some aggregates may need to consume/materialize substantial input before returning results.

**Second point.** Cost estimates combine estimated rows and per-operation work; selectivity estimates can change join order and scan choice. `EXPLAIN` reports planned estimates; `EXPLAIN ANALYZE` executes and adds actual measurements, with instrumentation overhead and workload/cache dependence.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Implement a `Filter` operator that wraps any child iterator and a predicate function, calling the child's `Next()` and only returning rows where the predicate is true — confirm it correctly stops (returns `hasMore = false`) once the child is exhausted, not just once it happens to find no more matching rows in a finite lookahead.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement the iterator interface (`Open()`, `Next() (row, hasMore)`, `Close()`, or the Go-idiomatic equivalent) and a `SeqScan` operator that iterates over an in-memory slice of rows (standing in for reading table pages, per Lesson 1) one at a time.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Verify iterator output against a hand-computed result. For a database query, record estimated/actual rows and chosen operators for selective and nonselective predicates; explain the planner choice rather than requiring a particular scan method.

**Failure mode.** Use a highly nonselective predicate or stale statistics and compare estimated versus actual row counts. Explain why an index plan may no longer be cheapest and how estimate error can affect downstream plan choices.

**Completion criterion.** Your operators preserve row semantics, and you can interpret cardinality, selectivity, and plan estimates while distinguishing actual execution measurements from cost-model predictions.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 083 — Query Execution and Cost Models case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement the iterator interface (`Open()`, `Next() (row, hasMore)`, `Close()`, or the Go-idiomatic equivalent) and a `SeqScan` operator that iterates over an in-memory slice of rows (standing in for reading table pages, per Lesson 1) one at a time.

2. Extend it with one explicit boundary case: Implement a `Filter` operator that wraps any child iterator and a predicate function, calling the child's `Next()` and only returning rows where the predicate is true — confirm it correctly stops (returns `hasMore = false`) once the child is exhausted, not just once it happens to find no more matching rows in a finite lookahead.

3. Rebuild the same behavior with one changed constraint: Implement a `Project` operator that wraps a child iterator and a list of column indices/names to keep, transforming each row before passing it up.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Run your step 4 comparison again, but this time with a *non-selective* predicate — one matching most or all of the rows in the dataset (e.g.

2. `WHERE x > 0` against data where nearly every row satisfies it).

3. Confirm the index-scan approach no longer shows a clear advantage over the sequential scan — it may even be slower, since traversing the B+ tree structure plus following pointers back to full row data can cost more overhead than simply reading rows sequentially when nearly all of them are going to be returned anyway.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 083 — Query Execution and Cost Models: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 083 — Query Execution and Cost Models?
2. What two concrete practices are required in the assigned material for # 083 — Query Execution and Cost Models?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson083/EVIDENCE.md` with exactly this structure:

```text
# Lesson 083 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 083 — Query Execution and Cost Models.
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

You pass when you can explain the central model of # 083 — Query Execution and Cost Models without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your composed scan/filter/project pipeline produces correct results using the iterator pattern, you've built and correctly used an index-scan alternative to a sequential scan, and you've measured — not just asserted — that the index scan wins decisively on a selective query and loses or ties on a non-selective one, connecting this directly to how real query. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
