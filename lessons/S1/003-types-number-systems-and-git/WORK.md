# Lesson 003 — Problem Set

**Topic:** Types, Number Systems, and Git.
**Your files:** `code/floatcheck.c`, `code/limits.c`, `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. `limits.c`: print `INT_MIN`, `INT_MAX`, `UINT_MAX` from `<limits.h>`; then demonstrate unsigned wrap (`0u - 1`) and `uint8_t` wrap. **Do not** demonstrate signed overflow in a normal build; instead, build a *separate* file with `-fsanitize=undefined` and read the diagnostic.
2. `floatcheck.c`: show `0.1 + 0.2 == 0.3` fails, then write `nearly_equal(a, b)` with an epsilon and show it succeeds.
3. In a fresh directory run `git init`, commit `limits.c`, edit it, view `git diff`, commit again, and show `git log --oneline` with at least three meaningful commits.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

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

## Stretch problem

### Typed conversion CLI

Build a C CLI that accepts integers in decimal and hexadecimal, converts between representations, and rejects malformed input without undefined behavior.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
