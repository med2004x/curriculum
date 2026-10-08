# Lesson 001 — Programs, Machines, and the Terminal Problems

Read [README.md](README.md), then complete these tasks in order. Write your predictions in [EVIDENCE.md](EVIDENCE.md) before running programs.

## Problems

1. In `code/task_01.c`, print `hello` followed by a newline.
2. In `code/hello.c`, print `Hello, world!` followed by a newline.
3. In `code/greet.c`, store a name in a variable and print a greeting with a format specifier.
4. In `code/report.c`, print a labeled name and score on separate lines. In `code/arith.c`, compute and print `(17 + 5) * 3`.
5. In `code/exercise.c`, print `Hello, DART!` followed by a newline. Change its exit status to 7 once, observe the result, then restore 0.

## Run and check

From the repository root, enter the code folder and run these commands:

```powershell
Set-Location lessons/S1/001-programs-machines-and-the-terminal/code
make task_01
.\task_01.exe
make check
```

To build another program, use `make hello`, `make greet`, `make report`, `make arith`, or `make exercise`, then run its `.exe` file. `make` builds all six programs.

If the compiler or Make is missing, follow [SETUP.md](../../../SETUP.md).

## Submit

Record predictions, commands, output, one repaired error, and what you learned in [EVIDENCE.md](EVIDENCE.md).
