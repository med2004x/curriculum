# Lesson 007 — Arrays, Strings, and Representation

## Threads

**Primary:** C

**Supporting:** arrays; strings; indexing; contiguous memory

**Math companion:** sets and sequences

**Time:** 4 hours. **Prerequisites:** Lesson 006.

## What you will be able to do

1. Compute an array's element count from its declared type.
2. Explain the NUL byte and the difference between array size and string length.
3. Recognize an out-of-bounds index as invalid before trusting its result.
4. Use sanitizers as evidence for a boundary error.

## Retrieve first (no notes)

1. What does pass-by-value mean in C?
2. What makes recursive termination possible?
3. What is the verified result of `fact(5)`?
4. What does an invariant tell you about a loop?

## 1. Array size and element count

A declared array contains a fixed number of elements. For `int a[5]`, the verified `sizeof a` is 20, the element count expression is 5, and adjacent elements differ by 4 bytes in the measured layout.

Worked example:

```c
#include <stdio.h>

int main(void) {
    int a[5] = {10, 20, 30, 40, 50};
    printf("%zu\n", sizeof a);
    printf("%zu\n", sizeof a / sizeof a[0]);
    return 0;
}
```

Self-check: divide the array size by one element size.

## 2. Bounds are part of the representation

The valid positions in an array of length 5 are 0 through 4. `a[5]` is outside the object.

Worked example:

```c
#include <stdio.h>

int main(void) {
    int a[5] = {10, 20, 30, 40, 50};
    printf("%d\n", a[4]);
    return 0;
}
```

The last valid index is 4.

Self-check: write every valid index before you write an access.

## 3. C strings use a NUL terminator

For `char s[] = "hello"`, the verified `sizeof s` is 6, `strlen(s)` is 5, and `s[5]` is zero. The extra byte marks the end of the string.

Worked example:

```c
#include <stdio.h>
#include <string.h>

int main(void) {
    char s[] = "hello";
    printf("%zu %zu %d\n", sizeof s, strlen(s), s[5] == '\0');
    return 0;
}
```

Self-check: why are the measured values 6 and 5?

## 4. Not every character array is a string

`char t[5] = {'h','e','l','l','o'};` has size 5 and no terminator. Calling `strlen(t)` is undefined because the function searches beyond the intended array for a zero byte.

Worked example:

```c
#include <stdio.h>

int main(void) {
    char t[5] = {'h', 'e', 'l', 'l', 'o'};
    printf("%zu\n", sizeof t);
    return 0;
}
```

Self-check: which declaration supplies a NUL byte automatically?

## 5. `sizeof` can expose the representation

`sizeof` measures the object named by its operand. In the verified comparison, `sizeof s` is 6 for `char s[] = "hello"`, while `sizeof lit` is 8 for `char *lit = "abc"` on the verified x86-64 environment. Pointer semantics are deferred to Lesson 013.

Self-check: which expression measures an array object and which measures the declared pointer object?

## 6. Sanitizers expose boundary violations

The verified out-of-bounds access `a[5]` compiled with no warning under `-Wall -Wextra`. With AddressSanitizer enabled, it reported `stack-buffer-overflow ... READ of size 4`.

Worked example:

```c
#include <stdio.h>

int main(void) {
    int a[5] = {10, 20, 30, 40, 50};
    printf("%d\n", a[5]);
    return 0;
}
```

Self-check: why is successful compilation weaker evidence than a sanitizer check?

## Predict, then observe

1. Predict `sizeof(int a[5])`.
2. Predict the element count expression.
3. Predict the final valid index.
4. Predict `sizeof s` and `strlen(s)` for `"hello"`.
5. Predict whether `t` is a C string.
6. Predict `sizeof lit` in the verified environment.
7. Predict the sanitizer report category for `a[5]`.

Run each prediction before checking the result.

## Build from blank

1. `array_stats.c`: print array size and element count, then sum five integers.
2. `string_report.c`: report `sizeof`, `strlen`, and the final-byte value for a C string.
3. `safe_index.c`: reject indexes outside the valid range before accessing the array.

## Failure drill

1. Change `a[4]` to `a[5]` and run only under AddressSanitizer.
2. Remove the NUL by using the five-character initializer; explain why `strlen` is then invalid.
3. Use `strlen` when the program needs capacity; explain the semantic mismatch.
4. Use `sizeof lit` when the program needs string length; explain the mismatch.

## Transfer

The constraint changes: the data can contain an embedded NUL byte. Explain why a NUL-terminated string no longer represents the full stored byte count, then redesign the data around an explicit length.

## Retrieval questions

1. What is the verified size of `int a[5]`?
2. How many elements does `sizeof a / sizeof a[0]` report?
3. What is `strlen("hello")`?
4. Why is `char t[5] = {'h','e','l','l','o'}` not a string?
5. What are the verified ASCII values of `A`, `a`, and `0`?
6. What sanitizer report was observed for `a[5]`?

## Evidence to record

Create `lesson007/EVIDENCE.md` with exactly:

```text
Lesson 007 evidence
===================

Array:
sizeof a = 20
element count = 5
adjacent elements differ by 4 bytes

String:
sizeof s = 6
strlen(s) = 5
s[5] = 0

No terminator:
sizeof t = 5
strlen(t) is undefined

Characters:
'A' = 65
'a' = 97
'0' = 48
'a' - 'A' = 32
'7' - '0' = 7

Bounds:
a[5] compiles clean under -Wall -Wextra
ASan: stack-buffer-overflow ... READ of size 4
```

## Common mistakes

- Starting indexes at 1.
- Confusing capacity with string length.
- Forgetting the NUL byte.
- Treating compilation as proof of valid bounds.
- Using the wrong `sizeof` operand.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can reason about array bounds, distinguish object size from string length, and demonstrate an invalid access with a sanitizer rather than guessing about the result. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
