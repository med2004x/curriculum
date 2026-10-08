# Lesson 064 — Systems Lab: xv6 and Kernel Boundaries Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Select the repository's xv6-RISC-V version and one lab focused on syscalls, traps, memory, or scheduling.

2. Trace an existing implementation in source before changing it; write down the user API, dispatch path, handler, and invariant.

3. Implement the smallest specified change, then run the lab's provided tests and inspect the relevant failure path.

## Run and check

```powershell
Set-Location lessons/S2/064-xv6-systems-lab/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### MIT xv6 lab assignment

Complete the MIT 6.1810 system-call lab in its official xv6-labs checkout. The local `exercise.c` is only a source-tracing prompt.

## Course lab

The small local C starter is a prompt, not an xv6 implementation. Use the real [xv6 source](https://github.com/mit-pdos/xv6-riscv) and the [matching MIT lab handout](https://pdos.csail.mit.edu/6.S081/2026/labs/syscall.html). Clone the course lab repository (`git://g.csail.mit.edu/xv6-labs-2026`), check out the `syscall` branch as directed, and run `make grade`. Record the release/commit, selected lab, commands, results, and the kernel invariant in `EVIDENCE.md`.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
