# Lesson 011 — Testing, Verification, and Reading Code Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. In a fresh directory, write the enum and `classify` function from the contract without copying the code block.
2. Write a test function that compares actual and expected values and counts mismatches.
3. Add all six cases from the table, including both sides of the threshold and both invalid inputs.
4. Make the executable return zero only if the mismatch count is zero.
5. Compile with the command above, run it, and record the process status.

Before compiling, trace scores 49, 50, 100, and 101 by hand. Your trace should identify the branch each input takes, not only its final label.

## Run and check

```powershell
Set-Location lessons/S1/011-testing-verification-and-code-reading/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Verification harness

Write a minimal test harness for a small C library. Include normal cases, boundary cases, and one test that exposes a seeded defect.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
