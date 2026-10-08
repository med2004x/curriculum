# Lesson 020 — Problem Set

**Topic:** Foundations Project: In-Memory Data Tool.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define `LogRecord`, the line grammar, field bounds, and malformed-line policy in a short design sketch.

2. Implement a parser that accepts one valid record and rejects one malformed record without modifying the output record on failure.

3. Add file loading, an in-memory array with an explicit capacity policy, a search command, and clear stdout/stderr behavior.

4. Split the program into at least two `.c` files and a header; add automated tests for parser and search behavior.

Save the design sketch before coding. Keep memory allocation bounded and justified by the file-size/record-count policy.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/020-foundations-project-data-tool/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Data tool project

Extend the earlier in-memory data tool with filtering, sorting, aggregation, persistence format, and a repeatable regression test suite.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
