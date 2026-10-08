# Lesson 006 — Functions, the Call Stack, and Debugging

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Function contracts

A function is easier to reason about when its contract is known before its body. Inputs become parameters, and the return value carries the result.

Worked example:

```c
#include <stdio.h>

int square(int x) { return x * x; }

int main(void) {
    printf("%d\n", square(5));
    return 0;
}
```

The arithmetic gives 25.

## 2. Pass-by-value

C passes argument values into parameters. Reassigning the parameter does not reassign the caller's scalar variable.

Worked example:

```c
#include <stdio.h>

void swap_bad(int a, int b) {
    int tmp = a;
    a = b;
    b = tmp;
}

int main(void) {
    int x = 1;
    int y = 2;
    swap_bad(x, y);
    printf("%d %d\n", x, y);
    return 0;
}
```

The verified caller values remain `1 2`.

## 3. Calls and the stack

Each active function invocation has its own call-frame state. Recursive calls create another frame before the previous call returns.

Worked example:

```c
#include <stdio.h>

int fact(int n) {
    if (n <= 1) return 1;
    return n * fact(n - 1);
}

int main(void) {
    printf("%d\n", fact(5));
    return 0;
}
```

The verified values are `fact(5)=120` and `fact(12)=479001600`. Trace the calls down to the base case, then trace returns upward.

## 4. Stack observation

On the verified gcc 13 x86-64 machine, three recursion levels showed successive local addresses decreasing by `0x30`, which is 48 bytes. Exact addresses vary because runtime layout is not fixed.

Worked observation: the useful fact is the difference between frames, not a memorized address.

## 5. Recursion needs progress

A recursive function needs a reachable base case and a step that moves toward it.

Worked example:

```c
#include <stdio.h>

int countdown(int n) {
    if (n == 0) return 0;
    return countdown(n - 1);
}

int main(void) {
    printf("%d\n", countdown(5));
    return 0;
}
```

The recursive argument decreases toward zero.

## 6. Debugging is controlled observation

A useful debugging cycle is reproduce, isolate, inspect, hypothesize, change one thing, rerun, and record. A debugger lets you stop at a selected function, inspect current values, and move line by line.

Worked procedure: put a breakpoint at the first function whose contract might be wrong, inspect its parameters, step one line, and compare actual state to expected state.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
