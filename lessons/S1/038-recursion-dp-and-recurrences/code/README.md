# Code map — Lesson 038 — Recursion, Dynamic Programming, and Recurrences

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `exercise_test.go` | TEST | Automated check/scaffold. Use it to verify your own implementation. |
| `pattern_ref_01.go` | PATTERN | Incomplete reference pattern from the original material. Use it to study structure, then implement your own version. |
| `pattern_ref_02.go` | PATTERN | Incomplete reference pattern from the original material. Use it to study structure, then implement your own version. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Dynamic-programming solver

Solve a problem with both a recursive baseline and a memoized or tabulated version. Compare call counts and explain the recurrence.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## First run

```sh
go run exercise.go
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
