# Lesson 004 — Conditions, Boolean Logic, and Proof Thinking Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/classify.c`, `code/demorgan.c`, `code/leap.c`, `code/truth.c`, `code/exercise.c`.

## Problems

1. `truth.c`: print the full truth table for `(A && B) || !C`, all 8 rows, using nested loops or a counter over 0..7.
2. `demorgan.c`: for all four `(A, B)` pairs, assert `!(A && B) == (!A || !B)` and the dual law; print `ok` only if every row passes.
3. `classify.c`: given an integer read from `argv[1]` (use `atoi`), print whether it is negative, zero, or positive, and whether it is even or odd, correctly for negative numbers.
4. `leap.c`: a year is a leap year if divisible by 4, except centuries, unless divisible by 400. Write the condition, then test 1900, 2000, 2023, 2024. State your test cases *before* running.

## Run and check

```powershell
python tools/check_lesson_004.py
```

## Optional challenge

### Predicate checker

Implement a small C program that evaluates a set of business rules as Boolean predicates. Add a truth-table test and explain which conditions are mutually exclusive.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
