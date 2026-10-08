# Lesson 084 — Problem Set

**Topic:** Transactions, WAL, and Recovery.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Draw the order of a WAL record, durable log flush, commit record, and later data-page write.

2. Explain how checkpointing bounds recovery work without replacing WAL.

3. Separately trace the prescribed two-session isolation anomaly and record the exact PostgreSQL result/error.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/084-transactions-wal-and-recovery/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### WAL + recovery simulation

Model a sequence of page changes and WAL records, simulate a crash at two points, and determine the recovered state under each point.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
