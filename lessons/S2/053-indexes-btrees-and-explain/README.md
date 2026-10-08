# Lesson 053 — Indexes, B-Trees, and EXPLAIN

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand how ordered indexes reduce data access work, then use PostgreSQL's planner to test whether an index helps a real query.

**The contract.** B-tree-family indexes keep keys ordered in a balanced, high-fanout page structure, supporting equality/range lookup and ordered scans. PostgreSQL's planner compares estimated costs and may choose a sequential scan when it expects that to be cheaper; index presence does not imply index use.

**Worked example.** For `SELECT id, title FROM app.tasks WHERE project_id = $1 ORDER BY created_at DESC LIMIT 20`, consider a composite index beginning with `project_id` and followed by `created_at DESC`. Compare plans and result ordering before and after adding the index; account for whether the query is selective and whether the index covers requested columns.

## 2. Mechanism and vocabulary

**Mechanism.** A B-tree's high fanout reduces tree height, and index pages can be cached. PostgreSQL may use index scans, index-only scans when visibility and included data permit, bitmap scans, or sequential scans depending on the query and estimates.

**Second point.** Indexes cost disk space and add work to inserts, updates, and deletes. A query returning a large fraction of a table can be cheaper as a sequential scan, so compare actual workload and plans rather than assuming indexes are always faster.

`EXPLAIN` shows a plan and estimates. `EXPLAIN ANALYZE` executes the statement and reports observed rows/timing; use caution with statements that modify data, and remember runtime observations vary with cache state and workload.

**Worked example.** Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
open tasks for one project
newest first
limit 20
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** Capture the plan for the target query with and without the index, along with table size, predicate selectivity, estimated/actual rows when available, and sort work. Explain whether the planner's choice is reasonable; do not require it to use the index on a tiny table.

**Failure mode.** Create an index that does not match the leading filter/order columns or query expression. Inspect why the plan does not use it; consider type, operator, selectivity, and statistics rather than concluding the planner is broken.

**Completion criterion.** You can relate index structure to workload, create an appropriate index, read a plan, and explain the trade-off in write/storage cost and query execution.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
