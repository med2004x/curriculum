# Code map — Lesson 030 — Synchronization, Channels, and Mutexes

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `challenge.go` | CHALLENGE | Independent harder starter. Do not copy a demo into it. |
| `demo_ref_01.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `exercise_test.go` | TEST | Automated check/scaffold. Use it to verify your own implementation. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Concurrent pipeline

Build a three-stage Go pipeline using channels and cancellation. Prove that every goroutine has an exit path and that no stage can block forever.

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
