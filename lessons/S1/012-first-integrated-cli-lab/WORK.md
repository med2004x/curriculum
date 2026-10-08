# Lesson 012 — Problem Set

**Topic:** Integrated Lab 1: Small C CLI.
**Your files:** `code/log.h`, `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Write the record type, enum, and prototypes in `log.h`.
2. Implement and test `parse_record` before opening any file.
3. Implement `record_matches` and verify case sensitivity.
4. Write the CLI argument check, file loop, fixed-capacity array, and output path.
5. Compile the CLI and tests with the exact commands above; save their results.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/012-first-integrated-cli-lab/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Integrated C utility

Build a small CSV-like command-line data tool using the earlier C concepts. It must parse input, maintain records, compute a summary, and report malformed rows.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
