# Lesson 015 — Undefined Behavior, Sanitizers, and Defensive Debugging

## Threads
**Primary:** C/Linux

**Supporting:** undefined behavior; sanitizer tooling; defensive checks

**Math companion:** reasoning about program states and invariants

**Time:** 3–4 hours. **Prerequisites:** Lesson 014.

## What you will be able to do

1. Explain what “undefined behavior” means in the C standard and why a compiler is allowed to optimize aggressively around it.
2. Identify the main categories of UB in C: signed overflow, null dereference, out-of-bounds access, use-after-free, and uninitialized reads.
3. Explain why `-fsanitize=undefined` and `-fsanitize=address` are not optional when debugging low-level C code.
4. State one defensive habit that turns a silent bug into a visible, diagnoseable failure.

## Retrieve first (no notes)

1. What is the key property of a valid pointer in C?
2. What is the difference between an invalid pointer and a null pointer?
3. Why can a program that compiles cleanly still be wrong?
4. What does `malloc` return if the allocator fails?
5. What does a sanitizer add to a program at runtime?
6. Why is a compiler allowed to assume that UB never happens?

## 1. Problem and model

C is a language with a precise contract, but the contract is only guaranteed for well-defined programs. Once a program steps outside those rules, the standard gives no requirement for the result. That is the heart of undefined behavior: the program is no longer in a defined state, and the compiler is free to assume that the undefined case never occurs.

This matters because compilers use this assumption to optimize. A common example is signed overflow:

```c
#include <limits.h>
#include <stdio.h>

int main(void) {
    int x = INT_MAX;
    int y = x + 1; // UB: signed overflow
    printf("%d\n", y);
    return 0;
}
```

The expression `x + 1` may be folded away or transformed because the compiler may assume the earlier arithmetic was valid. The program is not “wrong in a portable way”; it is out of the language’s contract.

## 2. Mechanism and vocabulary

The important terms are:

- undefined behavior (UB): behavior that is not specified by the C standard; anything can happen
- implementation-defined behavior: the standard allows multiple valid outcomes, and the implementation must document which one it chooses
- sanitizer: compiler- or runtime-inserted checks that detect dangerous programming patterns
- optimization assumption: compilers often assume UB cannot occur and generate code based on that assumption

The classic UB examples are:

- signed integer overflow
- out-of-bounds array access
- null pointer dereference
- use-after-free or dangling pointer access
- reading an uninitialized local variable

A sanitizer does not “fix” the program; it turns silent corruption into a crash, diagnostic, or clear violation message before the program continues into nonsense.

## 3. Worked case: signed overflow and sanitizer output

Write a minimal program that adds one to `INT_MAX` and prints the value. Then compile it once without sanitizers and once with a sanitizer:

```bash
gcc -O2 -Wall -Wextra overflow.c -o overflow
gcc -O2 -fsanitize=undefined -g overflow.c -o overflow_ubsan
```

The unsanitized version may appear to work, or may produce a value that looks plausible. The sanitized version should emit an error pointing to the exact source line and the offending operation. This is the key lesson: UB is not just “bad code”; it is a state that makes optimizer assumptions invalid.

## 4. Verification, failure, and evidence

The evidence you record should include:

- the source line that triggers UB
- the compilation flags used
- the observed output or crash
- the sanitizer diagnosis

For example, a good note is:

- `gcc -O2 -Wall -Wextra` produced a plausible but wrong-looking result
- `gcc -O2 -fsanitize=undefined -g` reported signed integer overflow at line 6
- `-fsanitize=address` is useful for invalid memory access, not just arithmetic overflow

A defensive C habit is: always assume an invalid state is possible until proven otherwise. Check all pointer values, array bounds, and integer arithmetic before using an untrusted value.

## Predict, then observe

1. Predict what happens when `INT_MAX + 1` is computed in a signed `int`.
2. Predict whether a plain compiler warning will catch every UB case.
3. Predict the effect of compiling with `-O0` compared with `-O2` on the same UB example.
4. Predict which bug is more likely to be silent: integer overflow or a null pointer dereference.
5. Predict how AddressSanitizer and UndefinedBehaviorSanitizer differ in what they detect.
6. Predict which one of the two sanitizers is more likely to report a precise source line for an out-of-bounds access.

## Build from blank

1. Write a program that triggers signed overflow.
2. Compile it without a sanitizer and record the output.
3. Compile it with `-fsanitize=undefined` and record the diagnostic.
4. Repeat the same exercise with a null pointer or invalid index case if you want a second example.

Do not copy a finished answer before you have produced your own minimal example.

## Failure drill

1. Replace a valid pointer with `NULL` and write a dereference.
2. Access one element past the end of an array.
3. Read an uninitialized local variable.
4. For each case, record whether the failure is silent, deterministic, or diagnosed by a sanitizer.

The goal is not to memorize a list of errors. The goal is to see that UB turns an ordinary program into one whose behavior depends on compiler optimization and runtime layout.

## Transfer

Change exactly one constraint in the main UB case: use a different type, different signed boundary, or a different memory bug. For example, replace `INT_MAX` with `INT_MIN`, or use a one-past-the-end access in a small array. Explain why the bug still violates the underlying contract, and what the sanitizer would report.

## Retrieval questions

1. What does undefined behavior mean in C?
2. Why can a compiler optimize aggressively around UB?
3. What are two classes of UB that are especially common in C?
4. What does `-fsanitize=undefined` catch?
5. What does `-fsanitize=address` catch?
6. Why does a sanitizer improve the debugging process?

## Evidence to record

Create `lesson015/EVIDENCE.md` with exactly this structure:

```text
# Lesson 015 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating UB as “random” instead of as a contract violation.
- Assuming warnings are equivalent to correctness.
- Confusing a crash with a fix.
- Ignoring the compiler optimization level when debugging low-level C code.

## Precision checks

1. State the contract the program violated in one sentence.
2. Explain the role of the optimizer in one sentence.
3. Name one sanitizer that helps diagnose out-of-bounds memory access.
4. Name one sanitizer that helps diagnose signed overflow and invalid arithmetic.
5. Explain why a well-defined program is not the same as a program that happens to run once.

## University revision bridge

In the university-academic setting, you should be able to state the legal meaning of UB in the standard and connect it to the practical debugging tools used in C. The engineering task is to observe the real failure mode and then fix the contract violation; the academic task is to explain why the language left that case undefined.

## Assigned material

- `03-c-memory-and-systems/lesson-04-undefined-behavior/README.md`
- `05-linux-and-tools/lesson-05-debuggers/README.md`

## Mastery gate

You pass when you can explain UB in your own words, name at least two UB categories, and describe how sanitizers turn silent bad states into visible errors. You should be able to reproduce one small UB example and explain why the optimizer is allowed to transform it unpredictably.
