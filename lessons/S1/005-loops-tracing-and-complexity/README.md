# Lesson 005 — Loops, Tracing, and Complexity

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

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

## 4. Invariants

An invariant is a statement that remains true at a chosen program point. For the running sum loop, after iteration `i`, `sum == i*(i+1)/2`.

Worked example: at `i = 4`, the measured accumulator is 10, and `4*(5)/2` is also 10. The same relation holds at `i = 5`, where the trace gives 15.

## 5. Complexity

Count loop-body executions as a function of `n`. The verified counts are linear: 10/100/1000/10000; nested `n*n`: 100/10000/1000000/100000000; triangular: 45/4950/499500/49995000; doubling: 4/7/10/14 for `n = 10/100/1000/10000`.

Worked example: the triangular count is `n(n-1)/2`, so for `n = 1000` it is 499500. The doubling loop takes 10 verified steps for `n = 1000`.

## 6. Exact counting beats intuition

Two loops can both look short but perform very different amounts of work. The sum `1..100` is 5050, also equal to `100(101)/2`.

Worked example: a linear loop over `n=10000` performs 10000 body executions, while the nested pattern performs 100000000. The source code differs by one additional loop, but the work differs by four orders of magnitude.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
