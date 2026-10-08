# Lesson 063 — Virtual Memory, Allocators, and Garbage Collection Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Inspect `/proc/<pid>/maps` for a selected running process and classify several mappings by permissions and backing.

2. Record virtual-size and resident-memory metrics for a process, and explain why they are not interchangeable.

3. Compare C allocation behavior with Go heap/GC observations; state what each tool measures and what it cannot establish.

## Run and check

```powershell
Set-Location lessons/S2/063-virtual-memory-allocators-and-garbage-collection/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Allocator experiment

Build a simple allocator model or instrument allocation patterns, then compare fragmentation and lifetime behavior with garbage-collected allocation.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
