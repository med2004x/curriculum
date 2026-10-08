# Lesson 004 — Conditions, Boolean Logic, and Proof Thinking

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

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

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
