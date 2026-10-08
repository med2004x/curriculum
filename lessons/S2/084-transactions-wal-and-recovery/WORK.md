# Lesson 084 — Transactions, WAL, and Recovery Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Draw the order of a WAL record, durable log flush, commit record, and later data-page write.

2. Explain how checkpointing bounds recovery work without replacing WAL.

3. Separately trace the prescribed two-session isolation anomaly and record the exact PostgreSQL result/error.

## Run and check

```powershell
Set-Location lessons/S2/084-transactions-wal-and-recovery/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### WAL + recovery simulation

Model a sequence of page changes and WAL records, simulate a crash at two points, and determine the recovered state under each point.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
