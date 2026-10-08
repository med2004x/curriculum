# Lesson 006 — Functions, the Call Stack, and Debugging

## Threads

**Primary:** C

**Supporting:** functions; call frames; debugger workflow

**Math companion:** functions and relations

**Time:** 4 hours. **Prerequisites:** Lesson 005.

## What you will be able to do

1. Define a function with an explicit input, output, and contract.
2. Explain pass-by-value using a concrete trace.
3. Trace nested and recursive calls as stack frames.
4. Use a debugger to find the first incorrect state.

## Retrieve first (no notes)

1. State the invariant for the Lesson 005 running-sum loop.
2. How many iterations does `i <= 10` perform from zero?
3. What does `continue` do?
4. What is the verified gcd of 48 and 18?

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

Self-check: what result should `square(12)` return?

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

Self-check: which values did the function actually exchange?

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

Self-check: write the call chain for `fact(5)`.

## 4. Stack observation

On the verified gcc 13 x86-64 machine, three recursion levels showed successive local addresses decreasing by `0x30`, which is 48 bytes. Exact addresses vary because runtime layout is not fixed.

Worked observation: the useful fact is the difference between frames, not a memorized address.

Self-check: separate the stable observation from the machine-specific detail.

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

Self-check: identify the exact expression that guarantees progress.

## 6. Debugging is controlled observation

A useful debugging cycle is reproduce, isolate, inspect, hypothesize, change one thing, rerun, and record. A debugger lets you stop at a selected function, inspect current values, and move line by line.

Worked procedure: put a breakpoint at the first function whose contract might be wrong, inspect its parameters, step one line, and compare actual state to expected state.

Self-check: why is changing three lines before rerunning bad evidence?

## Predict, then observe

1. Predict `fact(5)`.
2. Predict `fact(12)`.
3. Predict the values of `x` and `y` after `swap_bad` with 1 and 2.
4. Predict whether recursive frames have independent local state.
5. Predict whether a recursion step that increases `n` approaches `n == 0`.
6. Predict the measured frame-to-frame difference on the verified machine.
7. Predict what should be observed before making the first code change.

Run each prediction before checking the result.

## Build from blank

1. `factorial.c`: implement recursive factorial and verify 5 and 12.
2. `swap.c`: demonstrate the unchanged caller values and explain pass-by-value.
3. `trace_calls.c`: create three nested functions and inspect the call stack in a debugger.

## Failure drill

1. Remove the recursion base case and label the source `/* DELIBERATELY BROKEN */`; the verified infinite-recursion experiment reached `Segmentation fault`, exit 139.
2. Make the recursive argument increase; observe loss of progress toward the base case.
3. Remove the temporary in `swap_bad`; trace which value is lost.
4. Set the debugger breakpoint after the bug; move it to the earliest observable wrong state.

## Transfer

The constraint changes: compute factorial iteratively while preserving the verified results `120` for 5 and `479001600` for 12. State what replaces the recursive call chain.

## Retrieval questions

1. What does pass-by-value mean in C?
2. What two properties make recursion terminate?
3. What does a stack frame represent?
4. What frame-to-frame address difference was measured?
5. What are `fact(5)` and `fact(12)`?
6. What should debugging collect before a code change?

## Evidence to record

Create `lesson006/EVIDENCE.md` with exactly:

```text
Lesson 006 evidence
===================

fact(5) = 120
fact(12) = 479001600

swap_bad(1, 2):
caller remains 1 2

Stack:
3-level recursion showed 0x30 = 48 byte decrease per frame on gcc 13 x86-64
exact addresses vary

Infinite recursion:
Segmentation fault
exit status = 139; default stack limit = 8192 KB
```

## Common mistakes

- Treating parameters as aliases for caller scalars.
- Writing recursion without a reachable base case.
- Debugging by changing several things at once.
- Treating one observed stack address as portable.
- Stopping after seeing a plausible result instead of checking the contract.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can state a function contract, explain pass-by-value from a trace, walk a recursive call chain, and use the debugger to identify the first wrong state. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
