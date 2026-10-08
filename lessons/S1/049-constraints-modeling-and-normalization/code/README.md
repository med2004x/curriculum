# Code map — Lesson 049 — Constraints, Modeling, and Normalization

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `challenge.sql` | CHALLENGE | Independent harder starter. Do not copy a demo into it. |
| `demo_ref_02.sql` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `demo_ref_04.sql` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `exercise.sql` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `pattern_ref_01.sql` | PATTERN | Incomplete reference pattern from the original material. Use it to study structure, then implement your own version. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Constraint-driven model

Model a small business domain using keys, foreign keys, uniqueness, checks, and normalized relationships. Show a malformed state each constraint prevents.

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
