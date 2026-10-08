# Code map — Lesson 068 — Handlers, Services, Repositories, and Boundaries

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `demo_ref_01.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `demo_ref_02.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `demo_ref_03.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `demo_ref_04.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Service boundaries

Split a backend feature into handler, service, and repository layers. Test each layer at its appropriate scope and document one boundary that prevents coupling.

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
