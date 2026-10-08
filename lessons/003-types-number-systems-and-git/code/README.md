# Code map — Lesson 003 — Types, Number Systems, and Git

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Typed conversion CLI

Build a C CLI that accepts integers in decimal and hexadecimal, converts between representations, and rejects malformed input without undefined behavior.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

Introduce one controlled defect in your implementation. Reproduce it from a clean run, classify the failure, identify the violated assumption, repair it, and add a regression check.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change one important constraint: larger input, malformed data, partial failure, concurrency, persistence, or a different representation. State which invariant must remain true and what implementation detail must change.

Then answer: **What would make your current solution wrong even though all visible examples still pass?**

## 10. Academic drill

1. **Closed-book definition:** State the key formal definition from this lesson in one or two sentences.
2. **Worked derivation:** Draw the state transition or representation for the concrete example used in the lesson.
3. **Exam variation:** Change one parameter and recompute the result without running code.
4. **Engineering bridge:** Identify the exact observation in the terminal that would confirm your written result.

This lesson is also an academic integration point. Do the written drill before the engineering implementation so the two reinfor

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
