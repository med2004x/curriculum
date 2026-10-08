# Lesson 020 — Foundations Project: In-Memory Data Tool Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Define `LogRecord`, the line grammar, field bounds, and malformed-line policy in a short design sketch.

2. Implement a parser that accepts one valid record and rejects one malformed record without modifying the output record on failure.

3. Add file loading, an in-memory array with an explicit capacity policy, a search command, and clear stdout/stderr behavior.

4. Split the program into at least two `.c` files and a header; add automated tests for parser and search behavior.

Save the design sketch before coding. Keep memory allocation bounded and justified by the file-size/record-count policy.

## Run and check

```powershell
Set-Location lessons/S1/020-foundations-project-data-tool/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Data tool project

Extend the earlier in-memory data tool with filtering, sorting, aggregation, persistence format, and a repeatable regression test suite.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
