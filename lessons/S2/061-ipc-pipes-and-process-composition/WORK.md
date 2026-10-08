# Lesson 061 — Problem Set

**Topic:** IPC, Pipes, and Process Composition.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end.

2. Extend it with one explicit boundary case: Ensure the child closes the write end when done, and the parent's read loop correctly detects EOF (a `read()` returning 0) and exits its loop rather than blocking forever.

3. Rebuild the same behavior with one changed constraint: Modify step 1 so the parent (consumer) deliberately reads slowly (add a small `sleep()` between reads) while the child (producer) writes a large volume of data quickly with no delay.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/061-ipc-pipes-and-process-composition/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### IPC pipeline

Compose two processes with pipes so one transforms another’s output. Handle broken-pipe and child-exit cases without hanging.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
