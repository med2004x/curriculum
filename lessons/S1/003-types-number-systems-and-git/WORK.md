# Lesson 003 — Types, Number Systems, and Git Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/floatcheck.c`, `code/limits.c`, `code/exercise.c`.

## Problems

1. `limits.c`: print `INT_MIN`, `INT_MAX`, `UINT_MAX` from `<limits.h>`; then demonstrate unsigned wrap (`0u - 1`) and `uint8_t` wrap. **Do not** demonstrate signed overflow in a normal build; instead, build a *separate* file with `-fsanitize=undefined` and read the diagnostic.
2. `floatcheck.c`: show `0.1 + 0.2 == 0.3` fails, then write `nearly_equal(a, b)` with an epsilon and show it succeeds.
3. In a fresh directory run `git init`, commit `limits.c`, edit it, view `git diff`, commit again, and show `git log --oneline` with at least three meaningful commits.

## Run and check

```powershell
Set-Location lessons/S1/003-types-number-systems-and-git/code
gcc -std=c17 -Wall -Wextra -Wpedantic floatcheck.c -o floatcheck.exe
.\floatcheck.exe
gcc -std=c17 -Wall -Wextra -Wpedantic limits.c -o limits.exe
.\limits.exe
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Typed conversion CLI

Build a C CLI that accepts integers in decimal and hexadecimal, converts between representations, and rejects malformed input without undefined behavior.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
