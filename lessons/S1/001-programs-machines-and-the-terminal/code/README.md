# Lesson 001 code

Start with the complete examples `hello.c`, `greet.c`, `report.c`, and `arith.c`. Compile one file at a time; each has its own `main` function. Change `exercise.c` for the assignment.

From this directory, with GCC installed:

```sh
gcc -std=c17 -Wall -Wextra -Wpedantic hello.c -o hello
./hello
```

Expected output: `Hello, world!`. Repeat for each example, changing the source and executable names. From the repository root, check your assignment with:

```sh
python tools/check_lesson_001.py
```

The check is expected to fail on the untouched starter and pass after it prints exactly `Hello, DART!`. If the compiler or C headers are missing, follow [`SETUP.md`](../../../SETUP.md).
