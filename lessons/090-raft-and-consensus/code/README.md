# Code map â€” Lesson 090 â€” Raft and Consensus

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `exercise_test.go` | TEST | Automated check/scaffold. Use it to verify your own implementation. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Consensus simulator

Implement the core state transitions of a simplified consensus algorithm and exercise leader loss, stale terms, and conflicting logs.

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

## Upstream lab: MIT 6.5840 Raft

Use the official course lab harness for consensus implementation and testing. The local state-transition file is a short reading exercise, not a substitute for the replicated state machine lab.

- Upstream source: git://g.csail.mit.edu/6.5840-golabs-2026
- Local audit checkout: `../../../../dart-upstreams/mit-6.5840`
- Lab guide: https://pdos.csail.mit.edu/6.5840/

Use the handout and tests from the same course release. In the supplied 2026 checkout, the Raft target is `make -C src raft1` (the `src/Makefile` runs the race-enabled `raft1` Go tests). Follow the course page and makefile for the selected release. Record the commit, command, and results in `EVIDENCE.md`.
