# Lesson 058 — Processes, Threads, and Scheduling Revisited Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Use `strace -c` (from Linux tools Lesson 5) on a simple program (e.g. one that reads a file and prints its contents) and identify which specific syscalls correspond to which C library calls you wrote (`fopen`→likely `openat`, `fread`→`read`, `printf`→eventually `write`).

2. Note any syscalls you didn't expect (e.g.

3. `mmap` calls from the C library's own startup or buffering).

## Run and check

```powershell
Set-Location lessons/S2/058-process-models-and-scheduling/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Scheduler simulator

Implement a simulator for two scheduling policies over the same workload and compute completion, turnaround, and waiting metrics.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
