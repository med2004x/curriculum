# Lesson 012 — Integrated Lab 1: Small C CLI Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/log.h`, `code/exercise.c`.

## Problems

1. Write the record type, enum, and prototypes in `log.h`.
2. Implement and test `parse_record` before opening any file.
3. Implement `record_matches` and verify case sensitivity.
4. Write the CLI argument check, file loop, fixed-capacity array, and output path.
5. Compile the CLI and tests with the exact commands above; save their results.

## Run and check

```powershell
Set-Location lessons/S1/012-first-integrated-cli-lab/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Integrated C utility

Build a small CSV-like command-line data tool using the earlier C concepts. It must parse input, maintain records, compute a summary, and report malformed rows.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
