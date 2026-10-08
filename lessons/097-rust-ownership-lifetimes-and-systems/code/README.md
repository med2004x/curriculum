# Code map — Lesson 097 — Rust Ownership, Lifetimes, and Systems Thinking

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.rs` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Rust systems exercise

Implement a small resource-owning Rust type and deliberately trigger an ownership/borrow error. Fix it while explaining the lifetime relationship.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## First run

```sh
rustc exercise.rs -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
