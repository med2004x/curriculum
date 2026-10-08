# Lesson 052 — CTEs, Window Functions, and Query Composition

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build readable analytical SQL and calculate per-row context such as rank, running totals, and previous-row values without collapsing rows into groups.

**The contract.** A CTE names a subquery for use by the following statement. A window function evaluates over related rows while retaining one output row per input row. `PARTITION BY` defines independent groups; `ORDER BY` defines order within each group; a frame determines which ordered rows contribute to frame-sensitive calculations.

**Worked example.** For task events, use `ROW_NUMBER() OVER (PARTITION BY task_id ORDER BY changed_at DESC, id DESC)` to number events newest-first within each task, then use a CTE to retain row number 1.

## 2. Mechanism and vocabulary

**Mechanism.** `ROW_NUMBER` assigns distinct sequential values; `RANK` leaves gaps after ties; `DENSE_RANK` does not. Add a unique tie-breaker when exactly one deterministic row must be selected.

**Second point.** An ordered window aggregate can compute a running total. Specify a frame such as `ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW` when row-by-row accumulation is intended; default frames can include peer rows depending on ordering and database semantics.

`LAG` and `LEAD` access preceding/following rows within an ordered partition. A window's ordering does not guarantee final result order; use the outer query's `ORDER BY` for presentation.

**Worked example.** Use a CTE to rank events and select each task's latest event. `GROUP BY task_id, MAX(changed_at)` finds the maximum time but does not by itself select other columns from the corresponding event row.

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

## 4. Verification, failure, and evidence

**Verification.** Check a fixture with several tasks, tied timestamps, and a task with no events. Verify whether the contract calls for one latest row or every row tied for latest.

**Failure mode.** Omit the tie-breaker and use `ROW_NUMBER` on equal timestamps. Explain why the selected row may vary; use `RANK` if all ties should be returned.

**Completion criterion.** You can compose query stages with a CTE, explain why window functions preserve rows, and state the partition, order, frame, and tie policy.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
