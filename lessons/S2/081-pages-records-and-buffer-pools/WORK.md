# Lesson 081 — Pages, Records, and Buffer Pools Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Define a page header and slot record format.
2. Implement insertion and retrieval for variable-length records with bounds checks.
3. Test empty, exact-fit, overflow, deletion, and compaction cases.
4. Add a small buffer-pool model with page ID, pin count, dirty state, and an eviction decision.

Do not rely on host struct layout as a portable on-disk format; define byte order and field encoding explicitly.

## Run and check

```powershell
Set-Location lessons/S2/081-pages-records-and-buffer-pools/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Page layout model

Model database pages and records in memory. Simulate inserts, deletes, and page splits while tracking free space and record locations.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
