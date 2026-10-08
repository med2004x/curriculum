# Lesson 052 — CTEs, Window Functions, and Query Composition

## Threads
**Primary:** Databases/Math

**Supporting:** SQL composition; analytic queries

**Math companion:** sequences and functions

**Time:** 3–4 hours. **Prerequisites:** Lesson 051.

**Time:** 3–4 hours. **Prerequisites:** Lesson 051.

## What you will be able to do

1. Use a common table expression (CTE) to name and compose a query stage.
2. Distinguish grouped aggregation, which reduces rows, from window functions, which preserve rows.
3. Use `PARTITION BY`, `ORDER BY`, and a frame to define rank or running aggregate behavior.
4. Explain tie and NULL behavior in analytical queries.

## Retrieve first (no notes)

1. How does a left join preserve unmatched rows?
2. What is the difference between `WHERE` and `HAVING`?
3. How can a many-to-many join multiply counts?
4. Why should expected join cardinality be worked out before aggregation?

## 1. Problem and model

**What problem this solves.** Build readable analytical SQL and calculate per-row context such as rank, running totals, and previous-row values without collapsing rows into groups.

**The contract.** A CTE names a subquery for use by the following statement. A window function evaluates over related rows while retaining one output row per input row. `PARTITION BY` defines independent groups; `ORDER BY` defines order within each group; a frame determines which ordered rows contribute to frame-sensitive calculations.

**Worked example.** For task events, use `ROW_NUMBER() OVER (PARTITION BY task_id ORDER BY changed_at DESC, id DESC)` to number events newest-first within each task, then use a CTE to retain row number 1.

Self-check: unlike `GROUP BY`, does `ROW_NUMBER` keep each event row in the result? Yes; it adds a value per row.

## 2. Mechanism and vocabulary

**Mechanism.** `ROW_NUMBER` assigns distinct sequential values; `RANK` leaves gaps after ties; `DENSE_RANK` does not. Add a unique tie-breaker when exactly one deterministic row must be selected.

**Second point.** An ordered window aggregate can compute a running total. Specify a frame such as `ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW` when row-by-row accumulation is intended; default frames can include peer rows depending on ordering and database semantics.

`LAG` and `LEAD` access preceding/following rows within an ordered partition. A window's ordering does not guarantee final result order; use the outer query's `ORDER BY` for presentation.

**Worked example.** Use a CTE to rank events and select each task's latest event. `GROUP BY task_id, MAX(changed_at)` finds the maximum time but does not by itself select other columns from the corresponding event row.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```sql
WITH ranked_events AS (
	SELECT task_id, status, changed_at,
		   ROW_NUMBER() OVER (
			   PARTITION BY task_id
			   ORDER BY changed_at DESC, id DESC
		   ) AS row_num
	FROM app.task_events
)
SELECT task_id, status, changed_at
FROM ranked_events
WHERE row_num = 1
ORDER BY task_id;
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

The unique `id` tie-breaker makes selection deterministic when timestamps match. State how NULL timestamps should be handled under the application contract.

Self-check: would replacing `ROW_NUMBER` with `RANK` always return exactly one row per task when timestamps tie?

## 4. Verification, failure, and evidence

**Verification.** Check a fixture with several tasks, tied timestamps, and a task with no events. Verify whether the contract calls for one latest row or every row tied for latest.

**Failure mode.** Omit the tie-breaker and use `ROW_NUMBER` on equal timestamps. Explain why the selected row may vary; use `RANK` if all ties should be returned.

**Completion criterion.** You can compose query stages with a CTE, explain why window functions preserve rows, and state the partition, order, frame, and tie policy.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the row numbers for multiple events in two task partitions.
2. Predict which row `ROW_NUMBER` selects when timestamps tie and no tie-breaker exists.
3. Predict how `RANK` and `DENSE_RANK` differ on tied timestamps.
4. Predict whether adding a window expression changes the number of input rows.
5. Predict the effect of a `ROWS` frame on a running sum.
6. Predict whether a CTE's internal ordering guarantees final output order.
7. Predict the result when a task has no event rows.
8. State the tie and NULL policies before writing the query.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Write a query that ranks events newest-first within each task.

2. Put the ranking query in a CTE and select the latest event per task.

3. Add a tie-breaker and compare `ROW_NUMBER` with `RANK` when timestamps tie.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Call your `cosineSimilarity` function with one vector being the zero vector `(0, 0, 0, 0)`.

2. Observe what happens (likely division by zero, producing `NaN` or `Inf` in Go, or a runtime error in Python).

3. Explain why this is not a bug in your formula but a genuine mathematical undefined case — the angle between the zero vector and anything is undefined, since the zero vector has no direction — and state what defensive check a production implementation (like the kind pgvector or an embeddings pipeline would need) should include.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one requirement: return every row tied for latest, or return exactly one deterministic row. Choose and justify the ranking function and ordering.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does a CTE name?
2. How does a window function differ from `GROUP BY`?
3. What do `PARTITION BY` and `ORDER BY` control?
4. How do `ROW_NUMBER`, `RANK`, and `DENSE_RANK` treat ties?
5. Why can a tie-breaker be necessary?
6. What does an explicit window frame control?

## Evidence to record

Create `lesson052/EVIDENCE.md` with exactly this structure:

```text
# Lesson 052 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Assuming a window function reduces rows like `GROUP BY`.
- Using `ROW_NUMBER` without deterministic ordering when one row must be selected.
- Assuming a CTE's internal order determines final output order.
- Treating `MAX(timestamp)` as if it also selects that event's other columns.
- Leaving the running-total frame implicit.

## Precision checks

1. State the purpose of each CTE stage.
2. Identify partition, ordering, and frame for each window expression.
3. Explain the tie policy and determinism of selected rows.
4. Distinguish window output cardinality from grouped output cardinality.
5. State the intended NULL behavior.

CTEs name query stages; window expressions add context to rows without replacing those rows.
## University revision bridge

For the LIG track, connect query composition to relational operations and window partitions to grouped sequences, then make ordering and tie semantics explicit.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 052 — CTEs, Window Functions, and Query Composition without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can compute dot product, norm, and cosine similarity by hand for small vectors, your from-scratch implementation matches your hand calculations, and you can explain in one sentence why cosine similarity is undefined for a zero vector. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
