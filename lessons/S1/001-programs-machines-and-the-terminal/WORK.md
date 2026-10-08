# Lesson 001 - Work card

**Goal:** compile a C program, predict its output, and read its exit status.
**Time:** about 45-60 minutes.
**Start:** read [`README.md`](README.md), then do this card in order.

The lecture's Section 6 suggests creating a separate `lesson001/` folder. For this repository, use the existing `code/` folder and the files named below.

## Before coding

Write your predictions in `EVIDENCE.md` before running anything:

1. How many UTF-8 bytes represent four characters: `a`, e-acute, the euro sign, and the grinning-face emoji?
2. What should `(17 + 5) * 3` print?
3. What does exit status `0` mean?

## Exercises

### Exercise 1 - Build the first program

Complete `code/task_01.c`: make it print exactly `hello`, then a newline, and exit with status 0. Predict the output first; compile it, run it, and record the result.

### Exercise 2 - Match a string contract

Complete `code/hello.c` so it prints exactly `Hello, world!` followed by a newline.

### Exercise 3 - Use a variable in formatted output

Complete `code/greet.c`: store a name in a variable and print `Hello, <name>!` using a format specifier.

### Exercise 4 - Report two values

Complete `code/report.c`: store a name and integer score, and print them on separate labeled lines. Then complete `code/arith.c` to compute and print `(17 + 5) * 3` as an integer.

### Exercise 5 - Meet an exact output contract

Complete `code/exercise.c` so it prints exactly `Hello, DART!` followed by a newline. From the repository root, run `python tools/check_lesson_001.py`. It checks all six programs and should fail until every task is complete. Then temporarily change the final program's `return 0;` to `return 7;`, run it, and inspect the exit status. Restore `return 0;` and make the full check pass.

## Commands

In PowerShell, from the repository root:

```powershell
Set-Location ./lessons/S1/001-programs-machines-and-the-terminal/code
gcc -std=c17 -Wall -Wextra -Wpedantic task_01.c -o task_01.exe
./task_01.exe
$LASTEXITCODE
```

Build each task file separately; each has its own `main` function. For example, run `gcc -std=c17 -Wall -Wextra -Wpedantic task_01.c -o task_01.exe`, then `./task_01.exe`. Repeat for the file you are working on, using its name for the executable. Return to the repository root with `Set-Location ../../..`.

This Windows environment needs a complete C toolchain first. A compiler executable without the C standard headers is not enough. Follow [`SETUP.md`](../../../SETUP.md).

## Record and check

Fill in [`EVIDENCE.md`](EVIDENCE.md) as you go. Try the retrieval answers from memory, then use the lecture and compiler output to check your reasoning.

## Done when

- The check passes after your edit.
- You can explain why the starter failed and why your changed program passes.
- Your exit-status experiment prints `7`, and the restored program passes again.
