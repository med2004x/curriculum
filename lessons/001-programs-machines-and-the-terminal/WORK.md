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

### Exercise 1 - Run the hello program

Compile `code/hello.c`, run it, and record its output and exit status. Expected output: `Hello, world!`.

### Exercise 2 - Run the greeting program

Compile `code/greet.c` and run it. Expected output: `Hello, Ada!`.

### Exercise 3 - Run the report and arithmetic programs

Compile and run `code/report.c` and `code/arith.c`. Expected outputs are two lines (`Name: Ada`, then `Score: 42`) and `66` respectively.

### Exercise 4 - Change the student program

Edit `code/exercise.c` so it prints exactly `Hello, DART!` followed by a newline. Predict the output before running it.

### Exercise 5 - Check and inspect an exit status

From the repository root, run `python tools/check_lesson_001.py`. It should fail before Exercise 4 and pass after it. Then temporarily change the program's `return 0;` to `return 7;`, run it, and inspect the exit status. Restore `return 0;` and make the check pass again.

## Commands

In PowerShell, from the repository root:

```powershell
Set-Location ./lessons/001-programs-machines-and-the-terminal/code
gcc -std=c17 -Wall -Wextra -Wpedantic hello.c -o hello.exe
./hello.exe
$LASTEXITCODE
```

Repeat with `greet.c`, `report.c`, and `arith.c`, changing the executable name to match. Return to the repository root with `Set-Location ../../..`.

This Windows environment needs a complete C toolchain first. A compiler executable without the C standard headers is not enough. Follow [`SETUP.md`](../../../SETUP.md).

## Record and check

Fill in [`EVIDENCE.md`](EVIDENCE.md) as you go. Try the retrieval answers before checking [`ANSWERS.md`](ANSWERS.md).

## Done when

- The check passes after your edit.
- You can explain why the starter failed and why your changed program passes.
- Your exit-status experiment prints `7`, and the restored program passes again.
