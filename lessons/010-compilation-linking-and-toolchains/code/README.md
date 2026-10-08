# Code map — Lesson 010 — Compilation, Linking, and Toolchains

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `challenge.c` | CHALLENGE | Independent harder starter. Do not copy a demo into it. |
| `demo_01.c` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `exercise_main.c` | SUPPORT | Supporting file for the multi-file student exercise. |
| `exercise_math.c` | SUPPORT | Supporting file for the multi-file student exercise. |
| `exercise_math.h` | SUPPORT | Header/support file for a multi-file exercise. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Build pipeline

Create a tiny multi-file C project with a header, two implementation units, a test executable, and a Makefile. Break compilation and linking separately and diagnose each.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise_main.c exercise_math.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
