# Lesson 002 — Values, Bytes, and Your First C Program

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

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

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
