# Lesson 002 — Values, Bytes, and Your First C Program Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/convert.c`, `code/values.c`, `code/exercise.c`.

## Problems

Write `values.c` that prints, for the integer `n = 200`:

1. `n` in decimal, hex (`%x`), and as a character (`%c`);
2. the size in bytes of `int`, `long`, `char`, and `double` using `%zu`;
3. the quotient and remainder of `n` divided by 7 (and one negative case: `-n`);
4. `n / 3` computed as an integer and as a `double` with three decimals.

Then write `convert.c` that hard-codes one decimal number and prints its binary form, computed with a loop and `%` and `/` (do not use a library). Compare against your hand conversion from section 1.

## Run and check

```powershell
Set-Location lessons/S1/002-values-bytes-and-first-c-code/code
gcc -std=c17 -Wall -Wextra -Wpedantic convert.c -o convert.exe
.\convert.exe
gcc -std=c17 -Wall -Wextra -Wpedantic values.c -o values.exe
.\values.exe
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Byte counter

Write a C utility that reads a file byte-by-byte and reports bytes, newline count, and ASCII/non-ASCII byte counts. Verify it on both ASCII and UTF-8 input.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
