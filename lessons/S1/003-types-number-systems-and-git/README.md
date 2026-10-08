# Lesson 003 — Types, Number Systems, and Git

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

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

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
