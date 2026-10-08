# Code map — Lesson 100 — Graduation: Design, Build, Measure, Explain

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `CAPSTONE-README.md` | SUPPORT | Supporting file for the lesson code. |
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `exercise_test.go` | TEST | Automated check/scaffold. Use it to verify your own implementation. |
| `go.mod` | SETUP | Build or test support file. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Graduation capstone

Design, build, measure, secure, test, document, and defend a substantial backend or systems project. Treat the result as a technical review, not a demo.

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
