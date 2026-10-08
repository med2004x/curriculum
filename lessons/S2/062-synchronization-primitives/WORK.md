# Lesson 062 — Synchronization Primitives in the Kernel and in Go Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Implement a protected counter and ensure every access follows the mutex protocol.

2. Rewrite the counter as channel-owned state and identify the send/receive synchronization relation.

3. Add a `sync.Cond` predicate wait and explain lock release/reacquisition plus the required loop.

## Run and check

```powershell
Set-Location lessons/S2/062-synchronization-primitives/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Synchronization primitive

Implement or model a synchronization primitive and use it to protect a shared state. Compare the primitive’s guarantees with a Go channel solution.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
