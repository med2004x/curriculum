# Lesson 021 — Assembly and the Compiler View Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`, `code/exercise.s`.

## Problems

1. Write the two-argument `add` function and annotate its ABI-defined inputs and output.

2. Add a `sum_to_n` loop and trace its control flow and accumulator for `n = 3`.

3. Compare output at `-O0` and `-O2`; explain a transformation without claiming its textual form is guaranteed.

## Run and check

```powershell
Set-Location lessons/S1/021-assembly-and-the-compiler-view/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Compiler-view experiment

Compile the same C function at two optimization levels, inspect assembly, and explain three source-to-machine transformations visible in the output.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
