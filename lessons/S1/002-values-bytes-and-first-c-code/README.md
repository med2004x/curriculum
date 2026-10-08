# Lesson 002 — Values, Bytes, and Your First C Program

## Threads

**Primary:** C

**Supporting:** Bits/bytes; source code; shell execution

**Math companion:** number representations

**Time:** 3–4 hours. **Prerequisites:** Lesson 001.

## What you will be able to do

1. Convert between binary, hexadecimal, and decimal by hand.
2. Say how many values *n* bits can hold, and what range a signed or unsigned integer of a given width covers.
3. Use `printf` format specifiers correctly and predict what a wrong one prints.
4. Explain integer division and remainder, including with negative operands.

## Retrieve first (no notes, 5 minutes)

From Lesson 001: the four stages from source text to exit status; why `./` is needed; what exit status `0` means; why a 4-character file can be 10 bytes. Write the answers, *then* check Lesson 001. Skipping this step is the single most common reason the spacing effect fails.

## 1. Place value in any base

A number in base *b* is a sum of digits times powers of *b*. Decimal is base 10, binary base 2, hex base 16.

```text
1011 (binary) = 1·8 + 0·4 + 1·2 + 1·1 = 11
0x2f (hex)    = 2·16 + 15·1          = 47
```

**Decimal → binary by repeated division by 2.** Read remainders bottom to top.

```text
 37 / 2 = 18 r 1   <- least significant bit
 18 / 2 =  9 r 0
  9 / 2 =  4 r 1
  4 / 2 =  2 r 0
  2 / 2 =  1 r 0
  1 / 2 =  0 r 1   <- most significant bit
  => 37 = 100101 (binary) = 0x25
```

**Binary ↔ hex by grouping in fours from the right** (a hex digit is exactly 4 bits):

```text
1010 1111 0011  =  a f 3  =  0xaf3
```

Check yourself: `0xaf3 = 10·256 + 15·16 + 3 = 2803`.

## 2. How many values fit in n bits?

*n* bits give 2ⁿ distinct patterns.

| Bits | Patterns | Unsigned range | Signed (two's complement) range |
|---:|---:|---|---|
| 8 | 256 | 0 to 255 | −128 to 127 |
| 16 | 65 536 | 0 to 65 535 | −32 768 to 32 767 |
| 32 | 4 294 967 296 | 0 to 4 294 967 295 | −2 147 483 648 to 2 147 483 647 |

The signed range has one more negative value than positive because zero takes one of the non-negative slots.

## 3. Types have sizes, and the sizes are facts you can print

C does not fix the width of `int`; the standard sets minimums and the platform decides. Never assume, print:

```c
printf("int=%zu long=%zu ptr=%zu\n", sizeof(int), sizeof(long), sizeof(void*));
```

Measured on this course's reference platform (x86-64 Linux, gcc 13):

| Type | Size (bytes) | | Type | Size (bytes) |
|---|---:|---|---|---:|
| `char` | 1 | | `long long` | 8 |
| `short` | 2 | | `float` | 4 |
| `int` | 4 | | `double` | 8 |
| `long` | 8 | | `void *` | 8 |

`sizeof(long)` is 8 here but is 4 on 64-bit Windows. Portable code uses fixed-width types from `<stdint.h>` (`int32_t`, `uint8_t`) when width matters. `%zu` is the correct format for the `size_t` that `sizeof` produces.

## 4. Format specifiers are a contract

`printf` cannot check that your argument matches your specifier; *you* promise. Break the promise and the behaviour is undefined.

| Specifier | Expects | Example |
|---|---|---|
| `%d` | `int` | `printf("%d", 42)` |
| `%u` | `unsigned int` | `printf("%u", 42u)` |
| `%ld` / `%lld` | `long` / `long long` | |
| `%f` | `double` | `%.2f` gives two decimals |
| `%c` | a character as `int` | `printf("%c", 65)` prints `A` |
| `%x` | `unsigned int`, hex | `printf("%x", 255)` prints `ff` |
| `%zu` | `size_t` | for `sizeof` |
| `%p` | `void *` | cast to `(void*)` |
| `%%` | a literal `%` | |

Compile with `-Wall`: gcc checks format strings against arguments at compile time and warns on a mismatch. That warning is the reason to never build without it.

## 5. Arithmetic you must be able to do exactly

**Integer division truncates toward zero.** The remainder `%` takes the sign of the *dividend* (C99 and later).

```text
 7 / 2  =  3        7 % 2  =  1
-7 / 2  = -3       -7 % 2  = -1
 7 / -2 = -3        7 % -2 =  1
```

The identity `(a / b) * b + (a % b) == a` always holds. Verify it on the rows above.

**Integer versus floating division** depends on operand *types*, not on whether you wanted a fraction:

```c
7 / 2      // 3     both int
7 / 2.0    // 3.5   one operand is double, the other is converted
```

**Characters are small integers.** `'A'` is 65, `'a'` is 97, `'0'` is 48. So `'a' - 'A'` is 32, and `'7' - '0'` is the number 7. This is how you turn a digit character into a value.

## 6. Predict, then observe

Write predictions first.

1. `printf("%d\n", 7 / 2);` and `printf("%d\n", -7 / 2);`
2. `printf("%d\n", -7 % 2);`
3. `printf("%.2f\n", 7 / 2.0);`
4. `printf("%c %d\n", 'A' + 2, 'A' + 2);`
5. `printf("%x\n", 255);`
6. What is `sizeof(char)`? Is it ever anything other than 1?

Verified outputs to compare against: `3`, `-3`, `-1`, `3.50`, `C 67`, `ff`, and `1` (the standard *defines* `sizeof(char)` as 1).

## 7. Build from blank

Write `values.c` that prints, for the integer `n = 200`:

1. `n` in decimal, hex (`%x`), and as a character (`%c`);
2. the size in bytes of `int`, `long`, `char`, and `double` using `%zu`;
3. the quotient and remainder of `n` divided by 7 (and one negative case: `-n`);
4. `n / 3` computed as an integer and as a `double` with three decimals.

Then write `convert.c` that hard-codes one decimal number and prints its binary form, computed with a loop and `%` and `/` (do not use a library). Compare against your hand conversion from section 1.

## 8. Failure drill

Do each on purpose and record what you see.

1. Pass a `double` to `%d`. Note the compile warning text, then run it and note the output. The output is garbage and can differ across builds. *Write down why "it printed something" is not evidence that it works.*
2. Use `%d` for a `long` holding `5000000000`. Compare to `%ld`.
3. Divide an `int` by a divisor read at runtime (`atoi(argv[1])`, then run with `0`), and a `double` by `0.0`. Verified on gcc 13: the integer version dies with `Floating point exception` and exit status 136 (128 + SIGFPE; the message name is historical and misleading, since it is an *integer* trap). The `double` version prints `inf`. **Trap inside the trap:** if you write `int z = 0; 1 / z` where the compiler can see the zero, behaviour is undefined and one test build printed `0` and exited normally. Undefined means "anything", including "looks fine". That is the lesson.
4. Print `sizeof` of an array and of a pointer to it (you will meet this again in Lesson 007).

## 9. Transfer

Change the constraint: input is now a **negative** number in 8-bit two's complement, given as a pattern such as `1111 0110`. Work out its signed value by hand (invert the bits, add 1, negate: `1111 0110` → `0000 1001` → `0000 1010` → **−10**), then explain what your `convert.c` would need to change to print it.

## 10. Retrieval questions

1. What is 0xa7 in binary and decimal?
2. How many distinct values does a 12-bit field hold?
3. What are the ranges of `uint8_t` and `int8_t`?
4. Why is `sizeof(long)` not safe to assume?
5. What does `-7 / 2` print, and what does `-7 % 2` print? State the identity that connects them.
6. What does `'7' - '0'` evaluate to, and why is that useful?

## 11. Evidence to record

`lesson002/EVIDENCE.md`: your outputs for section 6 next to your *predictions* (mark each right or wrong); the compiler warning from failure drill 1 copied verbatim; the three retrieval answers you got wrong first time, with the correction.

## 12. Common mistakes

- Mixing up "how many bits" with "the largest value" (8 bits hold 256 values but the largest unsigned one is 255).
- Assuming `%` means "modulo" for negatives. In C it is the *remainder* and carries the dividend's sign.
- Writing `1 / 2` and expecting `0.5`.
- Trusting output from a mismatched `printf`. It compiled and printed; that proves nothing.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You may advance when, without notes, you can convert between binary/hex/decimal for 8-bit values, state the ranges of common integer widths, predict `printf` output for integer and floating division including negatives, and explain why a mismatched format specifier is not "working" just because it printed.

**Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
