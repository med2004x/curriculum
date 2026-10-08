# Code map — Lesson 053 — Indexes, B-Trees, and EXPLAIN

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `challenge.sql` | CHALLENGE | Independent harder starter. Do not copy a demo into it. |
| `demo_ref_02.sql` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `demo_ref_03.sql` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `demo_ref_04.sql` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `exercise.sql` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Index investigation

Create a realistic dataset, compare a query with and without an index, inspect `EXPLAIN`, and explain whether the planner decision matches your expectation.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## First run

```sh
psql "$DATABASE_URL" -f exercise.sql
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
