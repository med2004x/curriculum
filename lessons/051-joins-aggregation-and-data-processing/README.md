# Lesson 051 — Joins, Aggregation, and Data Processing

## Threads
**Primary:** Databases/Algorithms

**Supporting:** joins; GROUP BY; reports

**Math companion:** descriptive statistics

**Time:** 3–4 hours. **Prerequisites:** Lesson 050.

**Time:** 3–4 hours. **Prerequisites:** Lesson 050.

## What you will be able to do

1. Join task records to tag records through a junction table.
2. Explain how join type affects unmatched rows and why join predicates belong in the correct clause.
3. Aggregate joined rows with `GROUP BY` and filter groups with `HAVING`.
4. Detect duplicate-count errors caused by join cardinality and choose `COUNT(DISTINCT ...)` only when its semantics match the question.

## Retrieve first (no notes)

1. Why should SQL values be parameterized?
2. What does a transaction make atomic?
3. What should the application do if `Commit` fails?
4. How can `RETURNING` distinguish an updated row from no match?

## 1. Problem and model

**What problem this solves.** Retrieve useful report data that spans related tables without moving join and aggregation logic into application loops.

**The contract.** Tasks and tags have a many-to-many relationship through `task_tags(task_id, tag_id)`. Joining tasks to task_tags to tags associates each task with its tags. Aggregation must state its grouping key and account for how joins can multiply rows.

**Worked example.** List task titles with their tag names using `tasks JOIN task_tags JOIN tags`. Then count tags per task with `GROUP BY task_id, title`; use a `LEFT JOIN` if tasks with no tags must remain in the report.

Self-check: for a task with no tag rows, does an inner join or left join preserve that task?

## 2. Mechanism and vocabulary

**Mechanism.** An inner join returns row pairs satisfying the join predicate. A left join preserves each left-side row and supplies NULLs for unmatched right-side values. Moving a right-side filter from `ON` to `WHERE` can remove those NULL-extended rows and make a left join behave like an inner join for that predicate.

**Second point.** `GROUP BY` forms groups before aggregate results are returned; `HAVING` filters groups after aggregation, while `WHERE` filters input rows before grouping. `COUNT(*)` counts rows; `COUNT(column)` excludes NULL values in that column.

Joining one task to several tags creates several joined rows for that task. Counting task IDs after such a join may overcount tasks; use the right grouping key or `COUNT(DISTINCT task_id)` only when the intended quantity is distinct tasks.

**Worked example.** Count tags for each task with `GROUP BY`. Add a status predicate and compare it in `WHERE` versus `HAVING`; explain the difference between filtering source rows and filtering aggregate groups.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
tags
task_tags(task_id, tag_id)
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

Self-check: change exactly one constraint before you run it and write your prediction first.

## 4. Verification, failure, and evidence

**Verification.** Use data containing a task with no tags, a task with multiple tags, and repeated tag names across tasks. Compare inner and left joins; hand-check the grouped counts and distinct-task count.

**Failure mode.** Put a right-table filter in `WHERE` after a `LEFT JOIN` and observe which unmatched tasks disappear. Then move the filter into `ON` and explain the changed result.

**Completion criterion.** You can predict join cardinality, preserve unmatched rows when needed, and produce correct grouped counts without multiplying entities accidentally.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the output row count for a task with three tags under an inner join.
2. Predict whether a task with zero tags appears under a left join.
3. Predict how a predicate in `WHERE` changes NULL-extended rows.
4. State the difference between `WHERE` and `HAVING`.
5. Predict the result of `COUNT(*)` versus `COUNT(tag_id)` on a left join.
6. Predict when `COUNT(DISTINCT task_id)` differs from `COUNT(task_id)`.
7. Identify the grouping columns needed for one row per task.
8. State the expected cardinality before reading the query result.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Query task/tag pairs by joining all three tables.

2. Add a left join so tasks without tags remain visible.

3. Aggregate tag counts per task and add a `HAVING` condition on the count.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Filter a left-joined table in `WHERE` and identify the lost unmatched rows.

2. Move the predicate into `ON` and compare the preserved rows.

3. Create a task with multiple tags and compare raw row count with distinct-task count to expose join multiplication.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 051 — Joins, Aggregation, and Data Processing: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does an inner join return?
2. How does a left join preserve unmatched rows?
3. How can a `WHERE` predicate null-reject a left-joined row?
4. What is the difference between `WHERE` and `HAVING`?
5. How can a many-to-many join multiply counts?
6. When is `COUNT(DISTINCT ...)` appropriate?

## Evidence to record

Create `lesson051/EVIDENCE.md` with exactly this structure:

```text
# Lesson 051 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Using an inner join when unmatched left-side rows must remain.
- Filtering NULL-extended rows in `WHERE` unintentionally.
- Counting joined rows as though each row represented one entity.
- Selecting nonaggregated columns that are not part of the grouping key.
- Confusing row filtering with group filtering.

## Precision checks

1. Predict join cardinality from the relationship data.
2. Explain inner versus left-join behavior.
3. Verify aggregates against a hand-counted small dataset.
4. Place row and group filters correctly.
5. Explain duplicate counts and NULL behavior.

Query correctness depends on cardinality: count the rows created by joins before trusting the aggregate.
## University revision bridge

For the LIG track, connect relational joins to set relationships and aggregation to grouped functions, while accounting for SQL's NULL and duplicate semantics.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 051 — Joins, Aggregation, and Data Processing without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can compute all the core descriptive statistics by hand and verify your code implementation matches, and you can explain — using your own step 4 numbers as evidence, not just citing the general principle — why percentiles are preferred over the mean for skewed, real-world latency-like data. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
