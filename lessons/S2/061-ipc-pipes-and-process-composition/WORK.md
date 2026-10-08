# Lesson 061 — IPC, Pipes, and Process Composition Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end.

2. Ensure the child closes the write end when done, and the parent's read loop correctly detects EOF (a `read()` returning 0) and exits its loop rather than blocking forever.

3. Modify step 1 so the parent (consumer) deliberately reads slowly (add a small `sleep()` between reads) while the child (producer) writes a large volume of data quickly with no delay.

## Run and check

```powershell
Set-Location lessons/S2/061-ipc-pipes-and-process-composition/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### IPC pipeline

Compose two processes with pipes so one transforms another’s output. Handle broken-pipe and child-exit cases without hanging.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
