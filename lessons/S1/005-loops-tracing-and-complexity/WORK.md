# Lesson 005 — Loops, Tracing, and Complexity Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/gcd.c`, `code/loop_counter.c`, `code/sum_to_n.c`, `code/exercise.c`.

## Problems

1. `sum_to_n.c`: compute `1 + ... + n` with one loop and state its invariant.
2. `gcd.c`: implement Euclid's algorithm for 48 and 18 and record each `(a,b,r)` row.
3. `loop_counter.c`: compare linear, nested, and doubling body counts for command-line `n`.

## Run and check

```powershell
Set-Location lessons/S1/005-loops-tracing-and-complexity/code
gcc -std=c17 -Wall -Wextra -Wpedantic gcd.c -o gcd.exe
.\gcd.exe
gcc -std=c17 -Wall -Wextra -Wpedantic loop_counter.c -o loop_counter.exe
.\loop_counter.exe
gcc -std=c17 -Wall -Wextra -Wpedantic sum_to_n.c -o sum_to_n.exe
.\sum_to_n.exe
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Loop profiler

Build a command-line tool that compares three loop patterns on the same n and prints exact body counts. Explain the measured growth and the theoretical complexity.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
