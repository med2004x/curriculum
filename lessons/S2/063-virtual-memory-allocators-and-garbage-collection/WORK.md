# Lesson 063 — Problem Set

**Topic:** Virtual Memory, Allocators, and Garbage Collection.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Inspect `/proc/<pid>/maps` for a selected running process and classify several mappings by permissions and backing.

2. Record virtual-size and resident-memory metrics for a process, and explain why they are not interchangeable.

3. Compare C allocation behavior with Go heap/GC observations; state what each tool measures and what it cannot establish.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/063-virtual-memory-allocators-and-garbage-collection/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Allocator experiment

Build a simple allocator model or instrument allocation patterns, then compare fragmentation and lifetime behavior with garbage-collected allocation.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
