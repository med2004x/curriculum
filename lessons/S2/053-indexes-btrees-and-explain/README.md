# Lesson 053 — Indexes, B-Trees, and EXPLAIN

## Threads
**Primary:** Databases/Algorithms

**Supporting:** index access; B-tree structure; query plans

**Math companion:** logarithms and complexity

**Time:** 3–4 hours. **Prerequisites:** Lesson 052.

**Time:** 3–4 hours. **Prerequisites:** Lesson 052.

## What you will be able to do

1. Explain how a B-tree-family index narrows the search space and supports ordered access.
2. Create a PostgreSQL index that matches a stated filter/order workload.
3. Read `EXPLAIN` output to distinguish sequential scans, index scans, and bitmap plans.
4. Explain planner estimates, selectivity, and why an index can be slower than a sequential scan for some queries.

## Retrieve first (no notes)

1. What is the difference between `ROW_NUMBER` and `RANK`?
2. Why does a window function preserve rows while `GROUP BY` combines them?
3. What does `PARTITION BY` define?
4. Why should a ranking query state its tie policy?

## 1. Problem and model

**What problem this solves.** Understand how ordered indexes reduce data access work, then use PostgreSQL's planner to test whether an index helps a real query.

**The contract.** B-tree-family indexes keep keys ordered in a balanced, high-fanout page structure, supporting equality/range lookup and ordered scans. PostgreSQL's planner compares estimated costs and may choose a sequential scan when it expects that to be cheaper; index presence does not imply index use.

**Worked example.** For `SELECT id, title FROM app.tasks WHERE project_id = $1 ORDER BY created_at DESC LIMIT 20`, consider a composite index beginning with `project_id` and followed by `created_at DESC`. Compare plans and result ordering before and after adding the index; account for whether the query is selective and whether the index covers requested columns.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** A B-tree's high fanout reduces tree height, and index pages can be cached. PostgreSQL may use index scans, index-only scans when visibility and included data permit, bitmap scans, or sequential scans depending on the query and estimates.

**Second point.** Indexes cost disk space and add work to inserts, updates, and deletes. A query returning a large fraction of a table can be cheaper as a sequential scan, so compare actual workload and plans rather than assuming indexes are always faster.

`EXPLAIN` shows a plan and estimates. `EXPLAIN ANALYZE` executes the statement and reports observed rows/timing; use caution with statements that modify data, and remember runtime observations vary with cache state and workload.

**Worked example.** Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
open tasks for one project
newest first
limit 20
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Capture the plan for the target query with and without the index, along with table size, predicate selectivity, estimated/actual rows when available, and sort work. Explain whether the planner's choice is reasonable; do not require it to use the index on a tiny table.

**Failure mode.** Create an index that does not match the leading filter/order columns or query expression. Inspect why the plan does not use it; consider type, operator, selectivity, and statistics rather than concluding the planner is broken.

**Completion criterion.** You can relate index structure to workload, create an appropriate index, read a plan, and explain the trade-off in write/storage cost and query execution.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict which index columns support the project filter and descending time order.
2. Predict why a tiny table may use a sequential scan.
3. Predict the write/storage cost of adding another index.
4. Predict what changes when the query returns most rows rather than a selective subset.
5. Distinguish `EXPLAIN` estimates from `EXPLAIN ANALYZE` observations.
6. Predict when an index-only scan might still need heap visibility checks.
7. Predict how stale statistics can affect planner estimates.
8. State what evidence would justify retaining or dropping an index.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Capture a baseline plan for the project/task query and record the relevant table/index context.

2. Add a composite B-tree index matching the filter and ordering pattern; capture the new plan.

3. Compare selective and nonselective variants and explain the planner's cost trade-off.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Deliberately implement `rangeSearch` *without* the leaf-chain optimization — instead, for every key you want in the range, perform a fresh, independent tree search from the root.

2. Measure and compare the total number of node visits for a range query spanning, say, 20 keys using this naive re-search approach versus your step 3 leaf-chain approach.

3. Report the actual difference in node-visit count, and explain why the leaf-chain approach's advantage grows with the size of the range being queried — a range of 2 keys barely matters, but a range of thousands does, which is exactly the shape of query B+ tree indexes are specifically optimized to make cheap.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one query property: remove the `LIMIT`, change the sort order, or use a different leading filter. Predict whether the same index remains useful and explain from its key order.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Why does B-tree fanout matter for an index?
2. What query workload can a composite index support?
3. Why might PostgreSQL choose a sequential scan despite an index?
4. What does `EXPLAIN ANALYZE` add beyond `EXPLAIN`?
5. What are the write and storage costs of an index?
6. What information should accompany a plan comparison?

## Evidence to record

Create `lesson053/EVIDENCE.md` with exactly this structure:

```text
# Lesson 053 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 053 — Indexes, B-Trees, and EXPLAIN.
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

You pass when you can explain the central model of # 053 — Indexes, B-Trees, and EXPLAIN without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your B+ tree correctly implements search, insert-with-splitting, and leaf-chain range search, verified against a real multi-level tree (not just a single-node toy case), and you've directly measured and compared node-visit counts for both point lookups (versus a theoretical binary tree) and range queries (versus a naive re-search-per-key approach), connectin. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
