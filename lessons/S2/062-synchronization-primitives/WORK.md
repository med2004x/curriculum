# Lesson 062 — Problem Set

**Topic:** Synchronization Primitives in the Kernel and in Go.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement a protected counter and ensure every access follows the mutex protocol.

2. Rewrite the counter as channel-owned state and identify the send/receive synchronization relation.

3. Add a `sync.Cond` predicate wait and explain lock release/reacquisition plus the required loop.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/062-synchronization-primitives/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Synchronization primitive

Implement or model a synchronization primitive and use it to protect a shared state. Compare the primitive’s guarantees with a Go channel solution.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
