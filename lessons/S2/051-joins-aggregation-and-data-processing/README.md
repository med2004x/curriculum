# Lesson 051 — Joins, Aggregation, and Data Processing

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Retrieve useful report data that spans related tables without moving join and aggregation logic into application loops.

**The contract.** Tasks and tags have a many-to-many relationship through `task_tags(task_id, tag_id)`. Joining tasks to task_tags to tags associates each task with its tags. Aggregation must state its grouping key and account for how joins can multiply rows.

**Worked example.** List task titles with their tag names using `tasks JOIN task_tags JOIN tags`. Then count tags per task with `GROUP BY task_id, title`; use a `LEFT JOIN` if tasks with no tags must remain in the report.

## 2. Mechanism and vocabulary

**Mechanism.** An inner join returns row pairs satisfying the join predicate. A left join preserves each left-side row and supplies NULLs for unmatched right-side values. Moving a right-side filter from `ON` to `WHERE` can remove those NULL-extended rows and make a left join behave like an inner join for that predicate.

**Second point.** `GROUP BY` forms groups before aggregate results are returned; `HAVING` filters groups after aggregation, while `WHERE` filters input rows before grouping. `COUNT(*)` counts rows; `COUNT(column)` excludes NULL values in that column.

Joining one task to several tags creates several joined rows for that task. Counting task IDs after such a join may overcount tasks; use the right grouping key or `COUNT(DISTINCT task_id)` only when the intended quantity is distinct tasks.

**Worked example.** Count tags for each task with `GROUP BY`. Add a status predicate and compare it in `WHERE` versus `HAVING`; explain the difference between filtering source rows and filtering aggregate groups.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```text
tags
task_tags(task_id, tag_id)
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

Worked example check: identify the input, the operation, and the observation that would falsify your mental model.

## 4. Verification, failure, and evidence

**Verification.** Use data containing a task with no tags, a task with multiple tags, and repeated tag names across tasks. Compare inner and left joins; hand-check the grouped counts and distinct-task count.

**Failure mode.** Put a right-table filter in `WHERE` after a `LEFT JOIN` and observe which unmatched tasks disappear. Then move the filter into `ON` and explain the changed result.

**Completion criterion.** You can predict join cardinality, preserve unmatched rows when needed, and produce correct grouped counts without multiplying entities accidentally.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
