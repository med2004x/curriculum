# Lesson 005 — Loops, Tracing, and Complexity

## Threads

**Primary:** C

**Supporting:** iteration; invariants; tracing state

**Math companion:** asymptotics

**Time:** 4 hours. **Prerequisites:** Lesson 004.

## What you will be able to do

1. Trace a `for`, `while`, or `do ... while` loop by recording state changes.
2. Write a counter or accumulator whose meaning can be stated as an invariant.
3. Recognize linear, quadratic, triangular, and logarithmic iteration patterns.
4. Justify the time complexity of a simple loop nest.

## Retrieve first (no notes)

1. In C, what values are false in a condition?
2. What is the difference between `&&` and `&`?
3. What does `!5` evaluate to?
4. State De Morgan's law for `!(A && B)`.

## 1. Loop state

A useful trace records the variables that change and the condition controlling repetition.

Worked example:

```c
#include <stdio.h>

int main(void) {
    int sum = 0;
    for (int i = 1; i <= 5; i++) {
        sum += i;
        printf("%d %d\n", i, sum);
    }
    return 0;
}
```

The verified trace is `(1,1) (2,3) (3,6) (4,10) (5,15)`.

Self-check: what is `sum` after the fifth iteration?

## 2. Loop entry and boundaries

`for` and `while` test before entering the body. `do ... while` tests after the body.

Worked example:

```c
#include <stdio.h>

int main(void) {
    int n = 0;
    while (n > 0) puts("while body");
    do { puts("do body"); } while (n > 0);
    return 0;
}
```

With `n = 0`, the `while` body runs zero times and the `do` body runs once. A second boundary example, `i < 10` runs 10 times, while `i <= 10` runs 11 times.

Self-check: which loop guarantees one body execution?

## 3. Accumulators and control flow

A counter records how often an event occurs. An accumulator combines successive values. `continue` skips the remainder of one iteration; `break` leaves the loop.

Worked example:

```c
#include <stdio.h>

int main(void) {
    for (int i = 0; i <= 5; i++) {
        if (i == 2) continue;
        if (i == 5) break;
        printf("%d ", i);
    }
    putchar('\n');
    return 0;
}
```

The verified output is `0 1 3 4`.

Self-check: which value is skipped, and which value causes termination?

## 4. Invariants

An invariant is a statement that remains true at a chosen program point. For the running sum loop, after iteration `i`, `sum == i*(i+1)/2`.

Worked example: at `i = 4`, the measured accumulator is 10, and `4*(5)/2` is also 10. The same relation holds at `i = 5`, where the trace gives 15.

Self-check: compute the formula for `i = 5` without running code.

## 5. Complexity

Count loop-body executions as a function of `n`. The verified counts are linear: 10/100/1000/10000; nested `n*n`: 100/10000/1000000/100000000; triangular: 45/4950/499500/49995000; doubling: 4/7/10/14 for `n = 10/100/1000/10000`.

Worked example: the triangular count is `n(n-1)/2`, so for `n = 1000` it is 499500. The doubling loop takes 10 verified steps for `n = 1000`.

Self-check: why does multiplying the loop variable produce logarithmic growth?

## 6. Exact counting beats intuition

Two loops can both look short but perform very different amounts of work. The sum `1..100` is 5050, also equal to `100(101)/2`.

Worked example: a linear loop over `n=10000` performs 10000 body executions, while the nested pattern performs 100000000. The source code differs by one additional loop, but the work differs by four orders of magnitude.

Self-check: identify the extra nesting responsible for the quadratic count.

## Predict, then observe

1. Predict the fourth accumulator in the sum trace.
2. Predict the iteration count for `i < 10`.
3. Predict the iteration count for `i <= 10`.
4. Predict the `while` and `do ... while` behavior for `n=0`.
5. Predict the output of the `continue` and `break` program.
6. Predict the triangular count for `n=100`.
7. Predict the doubling count for `n=1000`.

Run each prediction before checking the result.

## Build from blank

1. `sum_to_n.c`: compute `1 + ... + n` with one loop and state its invariant.
2. `gcd.c`: implement Euclid's algorithm for 48 and 18 and record each `(a,b,r)` row.
3. `loop_counter.c`: compare linear, nested, and doubling body counts for command-line `n`.

## Failure drill

1. Change `i < 10` to `i <= 10`; identify the extra iteration.
2. Remove the loop update; identify the missing progress condition, then restore it.
3. Initialize `sum` to 1; find the invariant that fails.
4. Replace `continue` with `break`; explain the changed control flow.

## Transfer

Calculate `1..n` without a loop. Derive the formula from the invariant and verify `n=100`, where the verified sum is 5050.

## Retrieval questions

1. What is the difference between pre-test and post-test loops?
2. How many iterations does `i < 10` perform from zero?
3. What does `continue` do?
4. What is the running-sum invariant?
5. What is the triangular count for `n=100`?
6. Why does `i*=2` lead to logarithmic growth?

## Evidence to record

Create `lesson005/EVIDENCE.md` with exactly:

```text
Lesson 005 evidence
===================

Trace:
(1,1) (2,3) (3,6) (4,10) (5,15)

Bounds:
i < 10  -> 10
i <= 10 -> 11

Control flow:
0 1 3 4

GCD:
(48,18,12)
(18,12,6)
(12,6,0)
gcd = 6

Invariant:
sum == i*(i+1)/2

n = 1000:
linear = 1000
nested = 1000000
triangular = 499500
doubling = 10
```

## Common mistakes

- Confusing the condition with the body.
- Changing an accumulator without stating what it represents.
- Treating `<` and `<=` as interchangeable.
- Calling every loop `O(n)` without counting nested work.
- Using `break` when the algorithm needs to skip only one case.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can trace a loop without running it, state its invariant, justify its complexity, and complete the blank-page exercises. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
