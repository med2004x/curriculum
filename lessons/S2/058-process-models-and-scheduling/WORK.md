# Lesson 058 — Problem Set

**Topic:** Processes, Threads, and Scheduling Revisited.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Use `strace -c` (from Linux tools Lesson 5) on a simple program (e.g. one that reads a file and prints its contents) and identify which specific syscalls correspond to which C library calls you wrote (`fopen`→likely `openat`, `fread`→`read`, `printf`→eventually `write`).

2. Extend it with one explicit boundary case: Note any syscalls you didn't expect (e.g.

3. Rebuild the same behavior with one changed constraint: `mmap` calls from the C library's own startup or buffering).

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/058-process-models-and-scheduling/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Scheduler simulator

Implement a simulator for two scheduling policies over the same workload and compute completion, turnaround, and waiting metrics.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
