# Lesson 007 — Problem Set

**Topic:** Arrays, Strings, and Representation.
**Your files:** `code/array_stats.c`, `code/safe_index.c`, `code/string_report.c`, `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. `array_stats.c`: print array size and element count, then sum five integers.
2. `string_report.c`: report `sizeof`, `strlen`, and the final-byte value for a C string.
3. `safe_index.c`: reject indexes outside the valid range before accessing the array.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/007-arrays-strings-and-representation/code
gcc -std=c17 -Wall -Wextra -Wpedantic array_stats.c -o array_stats.exe
.\array_stats.exe
gcc -std=c17 -Wall -Wextra -Wpedantic safe_index.c -o safe_index.exe
.\safe_index.exe
gcc -std=c17 -Wall -Wextra -Wpedantic string_report.c -o string_report.exe
.\string_report.exe
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### String analyzer

Write a bounded C string analyzer that reports length, digit/letter counts, and reversed output without writing past the buffer.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
