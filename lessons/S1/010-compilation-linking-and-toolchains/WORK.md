# Lesson 010 — Problem Set

**Topic:** Compilation, Linking, and Toolchains.
**Your files:** `code/macro.c`, `code/main.c`, `code/util.c`, `code/util.h`, `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. `main.c`, `util.c`, `util.h`: define `add`, declare it in the header, and build the executable in explicit stages.
2. `macro.c`: verify `BAD_SQ(1+2) = 5` and `SQ(1+2) = 9`.
3. `symbols.md`: record `nm` results and explain every relevant `U` and `T` symbol.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/010-compilation-linking-and-toolchains/code
gcc -std=c17 -Wall -Wextra -Wpedantic main.c util.c -o add.exe
.\add.exe
gcc -std=c17 -Wall -Wextra -Wpedantic macro.c -o macro.exe
.\macro.exe
```

## Stretch problem

### Build pipeline

Create a tiny multi-file C project with a header, two implementation units, a test executable, and a Makefile. Break compilation and linking separately and diagnose each.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
