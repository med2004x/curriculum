# Lesson 053 — Problem Set

**Topic:** Indexes, B-Trees, and EXPLAIN.
**Your files:** `code/exercise.sql`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Capture a baseline plan for the project/task query and record the relevant table/index context.

2. Add a composite B-tree index matching the filter and ordering pattern; capture the new plan.

3. Compare selective and nonselective variants and explain the planner's cost trade-off.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

Run each SQL task against a disposable PostgreSQL database using `psql -v ON_ERROR_STOP=1 -f <task-file>`. Check the returned rows and verify invalid writes are rejected.

## Stretch problem

### Index investigation

Create a realistic dataset, compare a query with and without an index, inspect `EXPLAIN`, and explain whether the planner decision matches your expectation.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
