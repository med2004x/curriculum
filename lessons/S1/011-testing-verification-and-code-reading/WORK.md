# Lesson 011 — Problem Set

**Topic:** Testing, Verification, and Reading Code.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. In a fresh directory, write the enum and `classify` function from the contract without copying the code block.
2. Write a test function that compares actual and expected values and counts mismatches.
3. Add all six cases from the table, including both sides of the threshold and both invalid inputs.
4. Make the executable return zero only if the mismatch count is zero.
5. Compile with the command above, run it, and record the process status.

Before compiling, trace scores 49, 50, 100, and 101 by hand. Your trace should identify the branch each input takes, not only its final label.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/011-testing-verification-and-code-reading/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Verification harness

Write a minimal test harness for a small C library. Include normal cases, boundary cases, and one test that exposes a seeded defect.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
