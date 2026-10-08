# Lesson 004 — Conditions, Boolean Logic, and Proof Thinking

## Threads

**Primary:** C

**Supporting:** if/switch; predicates; truth tables

**Math companion:** logic and proofs

**Time:** 4 hours. **Prerequisites:** Lesson 003.

## What you will be able to do

1. Write and evaluate compound conditions and build a truth table for them.
2. State and apply De Morgan's laws, and simplify a condition correctly.
3. Use short-circuit evaluation deliberately and predict which side effects occur.
4. Write a small argument (a proof by cases or by counterexample) that a condition is correct.

## Retrieve first (no notes)

Why is signed overflow different from unsigned overflow? Why is `-1 < 1u` false? Why is `0.1 + 0.2 == 0.3` false? Check Lesson 003 afterwards.

## 1. What C treats as true

C has no true boolean in the core language. **Zero is false; any non-zero value is true.** Comparison and logical operators produce exactly `0` or `1`. `<stdbool.h>` gives you the names `bool`, `true`, `false`.

```text
a == b   1 if equal, else 0         !a       1 if a is 0, else 0
a != b   1 if different             a && b   1 if both non-zero
a <  b   etc.                       a || b   1 if either non-zero
```

Verified: `!5` is `0`, `!!5` is `1`, `5 && 7` is `1`, but `5 & 7` (bitwise) is `5`. Confusing `&&` with `&` and `||` with `|` is a top-five beginner bug.

## 2. The assignment-in-condition bug

```c
int x = 5;
if (x = 0) { ... }     // assigns 0, then tests 0 (false)
```

Verified: this assigns `0` to `x` and the branch is **not** taken. It compiles. gcc warns `suggest parentheses around assignment used as truth value` under `-Wall`. You wanted `x == 0`. The habit that prevents this: read your warnings, always.

## 3. Chained comparisons do not mean what mathematics means

```c
1 < 2 < 3      // 1
3 > 2 > 1      // 0 !
```

Both verified. `3 > 2 > 1` parses as `(3 > 2) > 1`, that is `1 > 1`, which is `0`. To test that `x` is between bounds, write `lo <= x && x <= hi`. gcc warns about this under `-Wall`.

## 4. Short-circuit evaluation

`&&` and `||` evaluate left to right and **stop as soon as the answer is known**. The right side may never run.

Instrumenting two functions that each count their calls (verified):

| Expression | Result | Right side called? |
|---|---:|---|
| `f() && t()` where `f` returns 0 | 0 | **No**, calls = 1 |
| `t() || f()` where `t` returns 1 | 1 | **No**, calls = 1 |
| `f() || t()` | 1 | Yes, calls = 2 |

This is a feature you use on purpose:

```c
if (p != NULL && p->count > 0) { ... }     // never dereferences a null p
if (n != 0 && total / n > limit) { ... }   // never divides by zero
```

The order matters. Swapping the operands turns a guard into a crash. It also means that side effects on the right (`i++`, a function call) may silently not happen.

## 5. Boolean algebra and De Morgan

Two laws you must know cold:

```text
!(A && B)  ==  !A || !B
!(A || B)  ==  !A && !B
```

Prove one by truth table, the only proof method you need for finite propositions:

| A | B | A && B | !(A && B) | !A | !B | !A \|\| !B |
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| 0 | 0 | 0 | **1** | 1 | 1 | **1** |
| 0 | 1 | 0 | **1** | 1 | 0 | **1** |
| 1 | 0 | 0 | **1** | 0 | 1 | **1** |
| 1 | 1 | 1 | **0** | 0 | 0 | **0** |

The bold columns match on all four rows, so the two expressions are equivalent. A table with *n* variables has 2ⁿ rows, which is why this stops scaling around 5–6 variables and why you learn algebraic laws too.

**Applying it.** "Reject if the age is under 18 or over 65" is `age < 18 || age > 65`. The accepting condition is its negation: `!(age < 18 || age > 65)` = `age >= 18 && age <= 65`. Note that `<` flips to `>=`, not `>`. Getting that boundary wrong is the most common off-by-one in conditions.

## 6. `if`, `else if`, and `switch`

An `if / else if / else` chain tests conditions **in order and takes the first that is true**. Order is part of the logic:

```c
if (score >= 50) grade = 'P';
else if (score >= 90) grade = 'A';     // unreachable: any score >= 90 already took the first branch
```

`switch` compares one integer expression against constants. It **falls through** to the next label unless you `break`:

```c
switch (2) {
    case 1: printf("one\n");
    case 2: printf("two\n");
    case 3: printf("three\n"); break;
    case 4: printf("four\n");
}
```

Verified output: `two` then `three`. Execution enters at `case 2` and runs on into `case 3` until the `break`. gcc flags it with `-Wimplicit-fallthrough`. Fallthrough is occasionally intended; make it visible with a comment, and never leave it by accident.

## 7. Proof thinking

A test shows a condition works on the inputs you tried. A **proof** shows it works on all inputs of a class. For a finite, small domain, exhaustion (enumerate every case) is a proof. For an infinite domain, use one of three moves:

- **By cases.** Split the domain into exhaustive, non-overlapping cases and show the claim in each.
- **By counterexample.** To *disprove* "for all x, P(x)", one value where P fails is enough.
- **By invariant** (Lesson 005): a fact true before and after each step.

**Example, by cases.** Claim: for any integer `n`, the value of `n % 2 == 0 || n % 2 == 1` is true. This is *false in C for negatives*: `-3 % 2` is `-1`, so both sides are 0. Counterexample found. The right test for "odd" that holds for all integers is `n % 2 != 0`. This is a real bug in real code, found by proof-thinking, not by luck.

## 8. Predict, then observe

Write predictions first; verified results are at the end of the list.

1. `printf("%d", !!7);`
2. `printf("%d", 5 && 7);` versus `printf("%d", 5 & 7);`
3. `int x = 5; if (x = 0) printf("A"); else printf("B");`: output, and the final value of `x`?
4. With `f()` returning 0 and counting calls, how many calls happen in `f() && f()`?
5. `3 > 2 > 1`
6. The `switch (2)` block above.
7. Does `-3 % 2 == 1` hold? What about `-3 % 2 != 0`?

Verified: `1`; `1` and `5`; `B`, `x` is `0`; `1` call; `0`; `two`/`three`; no (it is `-1`) and yes.

## 9. Build from blank

1. `truth.c`: print the full truth table for `(A && B) || !C`, all 8 rows, using nested loops or a counter over 0..7.
2. `demorgan.c`: for all four `(A, B)` pairs, assert `!(A && B) == (!A || !B)` and the dual law; print `ok` only if every row passes.
3. `classify.c`: given an integer read from `argv[1]` (use `atoi`), print whether it is negative, zero, or positive, and whether it is even or odd, correctly for negative numbers.
4. `leap.c`: a year is a leap year if divisible by 4, except centuries, unless divisible by 400. Write the condition, then test 1900, 2000, 2023, 2024. State your test cases *before* running.

## 10. Failure drill

1. Introduce `if (x = 0)` and record the compiler warning verbatim.
2. Reorder `p != NULL && p->count > 0` to put the dereference first and run it with a null pointer. Observe the crash.
3. In `classify.c`, use `n % 2 == 1` for "odd" and run it on `-3`. Record what breaks and fix it.
4. Remove a `break` from a `switch` and record the changed output.

## 11. Transfer

Change the domain: the condition now applies to **strings**, where "empty" means the first byte is `'\0'`. Write `is_blank(const char *s)` returning 1 when the string is empty or contains only spaces. State your test cases, including `NULL`, `""`, `"   "`, and `" a "`, and say in one sentence how short-circuiting protects you.

## 12. Retrieval questions

1. What are the only two values comparison operators produce?
2. State both De Morgan laws and negate `a > 0 && b <= 3`.
3. Why is `if (n != 0 && x / n > 2)` safe but `if (x / n > 2 && n != 0)` not?
4. Give a counterexample to "for every int n, n % 2 is 0 or 1".
5. What does `switch` do when you forget `break`?
6. Why does an `else if (score >= 90)` after `if (score >= 50)` never run?

## 13. Evidence to record

`lesson004/EVIDENCE.md`: your truth table output; the `leap.c` test table with predicted versus actual; the failure-drill diagnostics verbatim; a 5-sentence proof-by-cases that your `classify.c` odd/even test is correct for negatives.

## 14. Common mistakes

- `=` for `==`; `&` for `&&`; `3 > 2 > 1`.
- Negating a condition and forgetting that `<` flips to `>=`.
- Assuming `n % 2 == 1` means odd.
- Unreachable branches caused by ordering.
- Relying on the right side of `&&` or `||` to have run.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You may advance when, without notes, you can build a truth table, apply and prove De Morgan's laws, predict short-circuit side effects, write a boundary-correct condition, and give a counterexample-based disproof for a flawed condition.

**Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
