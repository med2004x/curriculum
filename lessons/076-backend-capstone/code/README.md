# Code map — Lesson 076 — Integrated Project 3: Production-Grade Go Service

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `README_CAPSTONE.md` | SUPPORT | Supporting file for the lesson code. |
| `demo_ref_02.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `demo_ref_03.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `go.mod` | SETUP | Build or test support file. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Production backend capstone

Build a production-shaped Go service combining routing, validation, persistence, authentication, idempotency, async work, observability, integration tests, and deployment notes.

### Acceptance criteria

- Start from a clean workspace or the latest lesson project without copying a completed solution.
- Implement the required behavior end-to-end.
- Include automated tests where the technology allows it, plus one deliberate failure test.
- Record design assumptions and the main invariant(s) in Markdown.
- Demonstrate the project from a reproducible command sequence.
- Finish with a short post-mortem: what failed, why, and what you changed.

## First run

```sh
go run exercise.go
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
