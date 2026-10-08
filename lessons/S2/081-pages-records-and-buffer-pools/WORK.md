# Lesson 081 — Problem Set

**Topic:** Pages, Records, and Buffer Pools.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Define a page header and slot record format.
2. Implement insertion and retrieval for variable-length records with bounds checks.
3. Test empty, exact-fit, overflow, deletion, and compaction cases.
4. Add a small buffer-pool model with page ID, pin count, dirty state, and an eviction decision.

Do not rely on host struct layout as a portable on-disk format; define byte order and field encoding explicitly.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/081-pages-records-and-buffer-pools/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Page layout model

Model database pages and records in memory. Simulate inserts, deletes, and page splits while tracking free space and record locations.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
