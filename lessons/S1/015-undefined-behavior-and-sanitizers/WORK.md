# Lesson 015 — Problem Set

**Topic:** Undefined Behavior, Sanitizers, and Defensive Debugging.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Write a program that triggers signed overflow.
2. Compile it without a sanitizer and record the output.
3. Compile it with `-fsanitize=undefined` and record the diagnostic.
4. Repeat the same exercise with a null pointer or invalid index case if you want a second example.

Do not copy a finished answer before you have produced your own minimal example.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/015-undefined-behavior-and-sanitizers/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Sanitizer lab

Take a small unsafe C program, reproduce at least two memory defects under sanitizers, then fix them while keeping the observable behavior unchanged.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
