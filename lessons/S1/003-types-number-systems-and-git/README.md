# Lesson 003 — Types, Number Systems, and Git

## Threads

**Primary:** C

**Supporting:** Integer/float representation; Git snapshots

**Math companion:** arithmetic and algebra foundations

**Time:** 4 hours. **Prerequisites:** Lesson 002.

## What you will be able to do

1. Explain two's complement and predict what happens on overflow for signed and unsigned types.
2. Explain why floating-point arithmetic is approximate and when `==` on floats is wrong.
3. Predict implicit type conversions, including the signed/unsigned comparison trap.
4. Use Git to snapshot, inspect, and undo your own work.

## Retrieve first (no notes)

Convert `0xa7` to binary and decimal. State the range of `int8_t` and `uint8_t`. Say what `-7 / 2` and `-7 % 2` print. Then check Lesson 002.

## 1. Unsigned integers wrap

An unsigned type with *n* bits does arithmetic **modulo 2ⁿ**. This is defined behaviour.

```c
unsigned u = 0;
u--;                 // 4294967295  (2^32 - 1)
uint8_t b = 250;
b += 10;             // 4           ((250+10) mod 256)
```

Both results are verified on gcc 13. Wrapping is why an unsigned loop counter that counts down and tests `>= 0` never ends: it is always true.

## 2. Signed integers: two's complement

To store a negative number in *n* bits, the top bit carries weight −2ⁿ⁻¹ instead of +2ⁿ⁻¹.

```text
8-bit pattern    unsigned    signed (two's complement)
0000 0000            0            0
0111 1111          127          127
1000 0000          128         -128
1111 1111          255           -1
```

**To negate: invert every bit, add 1.** `0000 0101` (5) → `1111 1010` → `1111 1011` (−5). Check: `0000 0101 + 1111 1011 = 1 0000 0000`; the ninth bit falls off and the result is 0, as it must be.

Why hardware uses this: addition works identically for signed and unsigned bit patterns, so the CPU needs one adder.

## 3. Signed overflow is undefined, not "wrapping"

This is the most important distinction in the lesson.

| | Behaviour |
|---|---|
| `unsigned` overflow | **Defined**: wraps modulo 2ⁿ |
| `signed` overflow | **Undefined behaviour**: the program has no meaning past this point |

```c
int x = INT_MAX;
printf("%d\n", x + 1);
```

On gcc 13 this printed `-2147483648`, and with `-fsanitize=undefined` it reports:

```text
runtime error: signed integer overflow: 2147483647 + 1 cannot be represented in type 'int'
```

Printing `-2147483648` is *what one machine did*, not what C promises. An optimising compiler is allowed to assume signed overflow never happens and may delete a check like `if (x + 1 < x)`. Never rely on wrap for signed values. To detect overflow, test *before* the operation (`if (x > INT_MAX - 1)`), or use unsigned arithmetic deliberately.

## 4. The signed/unsigned comparison trap

When an `int` and an `unsigned int` meet, the `int` is converted to unsigned.

```c
int x = -1; unsigned y = 1;
printf("%d\n", x < y);     // prints 0
```

Verified: this prints `0`. `-1` becomes `4294967295`, which is not less than 1. gcc warns with `-Wsign-compare` (part of `-Wextra`). This is a classic real-world bug; it is why you build with `-Wextra`.

**Rule of thumb for conversions:** operands are promoted to a common type, and mixing signed with unsigned of the same rank favours unsigned. `char` and `short` are promoted to `int` before arithmetic.

## 5. Floating point is approximate

A `double` has 64 bits: 1 sign, 11 exponent, 52 fraction. It stores a *binary* fraction, and most decimal fractions have no exact binary form, just as 1/3 has no exact decimal form.

```c
printf("%d\n", (0.1 + 0.2) == 0.3);        // 0
printf("%.17f\n", 0.1 + 0.2);              // 0.30000000000000004
```

Both verified. Consequences:

- **Never compare floats with `==`** after arithmetic. Compare `fabs(a - b) < epsilon`, choosing epsilon for your problem.
- **Precision is finite.** A 32-bit `float` has a 24-bit significand, so above 2²⁴ = 16 777 216 it cannot represent every integer: `16777216.0f + 1.0f` is still `16777216.0`. (Verified.)
- **Money must not be a float.** Store cents as an integer.
- `double` has ~15–16 significant decimal digits; `float` ~7.

## 6. Git: a database of snapshots

Git stores **snapshots** of your files, each identified by a hash and pointing to its parent. A *commit* is a named snapshot. You control what goes into it in two steps:

```text
working tree  --git add-->  staging area (index)  --git commit-->  repository history
```

Set up once:

```text
git config --global user.name  "Your Name"
git config --global user.email "you@example.com"
```

The daily loop:

```text
git init                    create a repository in this directory
git status                  what changed? what is staged?
git add file.c              stage a file (or: git add -p to choose hunks)
git commit -m "message"     snapshot what is staged
git log --oneline           history, one line per commit
git diff                    unstaged changes;   git diff --staged  = staged changes
git restore file.c          discard unstaged edits to a file (destructive)
```

**Commit messages** should state *what changed and why* in the imperative: `Add binary conversion loop`, not `stuff` or `fixed it`.

**`.gitignore`** lists files never to track. Compiled binaries and editor folders belong there:

```text
*.o
a.out
build/
```

Track source, not build output; anyone can regenerate the binary from the source.

## 7. Predict, then observe

1. `uint8_t b = 255; b++;`: value of `b`?
2. `printf("%u\n", 0u - 1u);`
3. `-1 < 1u`: true or false, and why?
4. Is `0.5 + 0.25 == 0.75` reliably true? (Yes: 0.5, 0.25, and 0.75 are exact binary fractions. Verify, then explain why this differs from `0.1 + 0.2`.)
5. What does `int x = 7; double d = x / 2;` store in `d`? (3.0, not 3.5. The division happens in `int` first.)

## 8. Build from blank

1. `limits.c`: print `INT_MIN`, `INT_MAX`, `UINT_MAX` from `<limits.h>`; then demonstrate unsigned wrap (`0u - 1`) and `uint8_t` wrap. **Do not** demonstrate signed overflow in a normal build; instead, build a *separate* file with `-fsanitize=undefined` and read the diagnostic.
2. `floatcheck.c`: show `0.1 + 0.2 == 0.3` fails, then write `nearly_equal(a, b)` with an epsilon and show it succeeds.
3. In a fresh directory run `git init`, commit `limits.c`, edit it, view `git diff`, commit again, and show `git log --oneline` with at least three meaningful commits.

## 9. Failure drill

1. Write a countdown `for (unsigned i = 5; i >= 0; i--)` and observe it never terminates (interrupt with Ctrl-C, or add a guard counter). Verified output: `5 4 3 2 1 0 4294967295 4294967294 ...`. gcc warns `comparison of unsigned expression in '>= 0' is always true [-Wtype-limits]`. Explain in one sentence, then fix it.
2. Commit a compiled binary by accident, then add it to `.gitignore`. Notice that ignoring does *not* untrack an already-committed file, and look up `git rm --cached`.
3. Make an edit you regret, and recover the previous version with `git restore`. Record the exact command.

## 10. Transfer

Change the constraint: you must add two **money** amounts entered as strings like `"19.99"` and `"0.02"` and print `20.01` exactly. Explain why `double` is the wrong tool, and design (on paper) how to parse the input into integer cents.

## 11. Retrieval questions

1. How does two's complement negate a number? Negate `0001 0100`.
2. Which overflow is defined and which is undefined? What does a compiler do with an undefined case?
3. Why does `-1 < 1u` evaluate to false?
4. Why is `0.1 + 0.2 != 0.3`?
5. What is the difference between `git add` and `git commit`?
6. Why is `a.out` not committed?

## 12. Evidence to record

`lesson003/EVIDENCE.md`: your `git log --oneline` output; the UBSan diagnostic copied verbatim; the epsilon you chose for `nearly_equal` and one sentence justifying it.

## 13. Common mistakes

- Treating signed wrap as guaranteed because your machine did it.
- Ignoring `-Wsign-compare` warnings.
- Comparing floats with `==`, or using floats for currency.
- `git add .` on a directory full of binaries.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You may advance when, without notes, you can explain why signed and unsigned overflow differ, predict a mixed signed/unsigned comparison, explain why float `==` is unreliable, and use Git to commit, inspect, and recover your work.

**Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
