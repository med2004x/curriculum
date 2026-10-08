# Code map — Lesson 012 — Integrated Lab 1: Small C CLI

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `challenge.c` | CHALLENGE | Independent harder starter. Do not copy a demo into it. |
| `exercise.c` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `log.c` | SUPPORT | Supporting file for the lesson code. |
| `log.h` | SUPPORT | Header/support file for a multi-file exercise. |
| `main.c` | SUPPORT | Supporting file for the lesson code. |
| `test_log.c` | TEST | Automated check/scaffold. Use it to verify your own implementation. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Integrated C utility

Build a small CSV-like command-line data tool using the earlier C concepts. It must parse input, maintain records, compute a summary, and report malformed rows.

### Acceptance criteria

- Start from a clean workspace or the latest lesson project without copying a completed solution.
- Implement the required behavior end-to-end.
- Include automated tests where the technology allows it, plus one deliberate failure test.
- Record design assumptions and the main invariant(s) in Markdown.
- Demonstrate the project from a reproducible command sequence.
- Finish with a short post-mortem: what failed, why, and what you changed.

## First run

```sh
gcc -Wall -Wextra -std=c11 main.c log.c -o lesson012 && ./lesson012
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
