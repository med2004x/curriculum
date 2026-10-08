# Lesson 015 — Undefined Behavior, Sanitizers, and Defensive Debugging Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Write a program that triggers signed overflow.
2. Compile it without a sanitizer and record the output.
3. Compile it with `-fsanitize=undefined` and record the diagnostic.
4. Repeat the same exercise with a null pointer or invalid index case if you want a second example.

Do not copy a finished answer before you have produced your own minimal example.

## Run and check

```powershell
Set-Location lessons/S1/015-undefined-behavior-and-sanitizers/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Sanitizer lab

Take a small unsafe C program, reproduce at least two memory defects under sanitizers, then fix them while keeping the observable behavior unchanged.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
