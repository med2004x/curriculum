# Lesson 083 — Query Execution and Cost Models

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build a minimal query execution pipeline (scan → filter → project) using the iterator pattern real database engines use, and learn to read a real `EXPLAIN` plan as a description of exactly this kind of operator tree.

**The contract.** **The iterator model (Volcano-style execution, the standard architecture).** A query plan is a tree of operators (scan, filter, project, join, sort, aggregate...), and each operator implements a uniform interface: typically `open()`, `next()` (returns the next row, or signals "no more rows"), and `close()`.

**Worked example.** Implement `SeqScan`, `Filter`, and `Project` operators. For a fixed input, record rows consumed by the scan, rows emitted by the filter, and columns emitted by projection. These counts make the pipeline's cardinality changes explicit.

## 2. Mechanism and vocabulary

**Mechanism.** In a pull iterator, a parent requests the next row from its child; filters may consume several child rows before producing one output. Blocking operators such as sort or some aggregates may need to consume/materialize substantial input before returning results.

**Second point.** Cost estimates combine estimated rows and per-operation work; selectivity estimates can change join order and scan choice. `EXPLAIN` reports planned estimates; `EXPLAIN ANALYZE` executes and adds actual measurements, with instrumentation overhead and workload/cache dependence.

**Worked example.** Implement a `Filter` operator that wraps any child iterator and a predicate function, calling the child's `Next()` and only returning rows where the predicate is true — confirm it correctly stops (returns `hasMore = false`) once the child is exhausted, not just once it happens to find no more matching rows in a finite lookahead.

## 3. Worked case from the reference material

Worked example: Implement the iterator interface (`Open()`, `Next() (row, hasMore)`, `Close()`, or the Go-idiomatic equivalent) and a `SeqScan` operator that iterates over an in-memory slice of rows (standing in for reading table pages, per Lesson 1) one at a time.

## 4. Verification, failure, and evidence

**Verification.** Verify iterator output against a hand-computed result. For a database query, record estimated/actual rows and chosen operators for selective and nonselective predicates; explain the planner choice rather than requiring a particular scan method.

**Failure mode.** Use a highly nonselective predicate or stale statistics and compare estimated versus actual row counts. Explain why an index plan may no longer be cheapest and how estimate error can affect downstream plan choices.

**Completion criterion.** Your operators preserve row semantics, and you can interpret cardinality, selectivity, and plan estimates while distinguishing actual execution measurements from cost-model predictions.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
