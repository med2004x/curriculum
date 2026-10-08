# Code map — Lesson 026 — File Descriptors and Go I/O

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `demo_01.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### File copier with boundaries

Build a Go file-copy utility with explicit buffering, error handling, and close-order checks. Test empty, large, and unreadable inputs.

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
