# Lesson 007 — Arrays, Strings, and Representation Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/array_stats.c`, `code/safe_index.c`, `code/string_report.c`, `code/exercise.c`.

## Problems

1. `array_stats.c`: print array size and element count, then sum five integers.
2. `string_report.c`: report `sizeof`, `strlen`, and the final-byte value for a C string.
3. `safe_index.c`: reject indexes outside the valid range before accessing the array.

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

## Optional challenge

### String analyzer

Write a bounded C string analyzer that reports length, digit/letter counts, and reversed output without writing past the buffer.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
