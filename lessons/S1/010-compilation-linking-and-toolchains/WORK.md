# Lesson 010 — Compilation, Linking, and Toolchains Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/macro.c`, `code/main.c`, `code/util.c`, `code/util.h`, `code/exercise.c`.

## Problems

1. `main.c`, `util.c`, `util.h`: define `add`, declare it in the header, and build the executable in explicit stages.
2. `macro.c`: verify `BAD_SQ(1+2) = 5` and `SQ(1+2) = 9`.
3. `symbols.md`: record `nm` results and explain every relevant `U` and `T` symbol.

## Run and check

```powershell
Set-Location lessons/S1/010-compilation-linking-and-toolchains/code
gcc -std=c17 -Wall -Wextra -Wpedantic main.c util.c -o add.exe
.\add.exe
gcc -std=c17 -Wall -Wextra -Wpedantic macro.c -o macro.exe
.\macro.exe
```

## Optional challenge

### Build pipeline

Create a tiny multi-file C project with a header, two implementation units, a test executable, and a Makefile. Break compilation and linking separately and diagnose each.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
