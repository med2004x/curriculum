# Lesson 021 — Problem Set

**Topic:** Assembly and the Compiler View.
**Your files:** `code/exercise.c`, `code/exercise.s`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Write the two-argument `add` function and annotate its ABI-defined inputs and output.

2. Add a `sum_to_n` loop and trace its control flow and accumulator for `n = 3`.

3. Compare output at `-O0` and `-O2`; explain a transformation without claiming its textual form is guaranteed.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/021-assembly-and-the-compiler-view/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Compiler-view experiment

Compile the same C function at two optimization levels, inspect assembly, and explain three source-to-machine transformations visible in the output.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
