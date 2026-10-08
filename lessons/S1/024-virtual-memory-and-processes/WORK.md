# Lesson 024 — Problem Set

**Topic:** Virtual Memory and Address Spaces.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Create a Linux program that maps or allocates a buffer and records its initial memory metrics.

2. Touch one byte per page in a selected range, then separately touch the whole range.

3. Record `/proc/self/status` `VmSize` and `VmRSS` before and after each phase; explain what these process metrics indicate and their limitations.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/024-virtual-memory-and-processes/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Address-space map

Write a small process-memory probe that records stack, heap, static, and mapped regions where observable. Explain virtual addresses versus physical memory.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
