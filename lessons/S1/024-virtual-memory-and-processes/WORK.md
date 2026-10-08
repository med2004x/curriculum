# Lesson 024 — Virtual Memory and Address Spaces Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Create a Linux program that maps or allocates a buffer and records its initial memory metrics.

2. Touch one byte per page in a selected range, then separately touch the whole range.

3. Record `/proc/self/status` `VmSize` and `VmRSS` before and after each phase; explain what these process metrics indicate and their limitations.

## Run and check

```powershell
Set-Location lessons/S1/024-virtual-memory-and-processes/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Address-space map

Write a small process-memory probe that records stack, heap, static, and mapped regions where observable. Explain virtual addresses versus physical memory.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
