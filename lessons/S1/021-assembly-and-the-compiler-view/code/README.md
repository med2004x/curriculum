# Code map — Lesson 021 — Assembly and the Compiler View

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Compiler-view experiment

Compile the same C function at two optimization levels, inspect assembly, and explain three source-to-machine transformations visible in the output.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Take the `-O0` assembly from step 1, manually flip the `add` instruction's operand order or change the immediate value, reassemble it by hand-editing (or note the change if you can't reassemble), and predict the wrong output before checking.

2. This confirms you're reading the semantics of the instruction, not pattern-matching the syntax.

3. Take the two-file program from step 4 and change `helper.h`'s function signature (e.g. change a parameter type) without recompiling `helper.c`, only recompiling and relinking against the stale `helper.o`.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint: target ARM64 instead of x86-64, or change the optimization level. State which observations should change (instructions, registers, or layout) and which source-level result should remain the same for defined code.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

Then a

## First run

```sh
gcc -Wall -Wextra -std=c11 exercise.c -o exercise && ./exercise
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
