# Lesson 018 — Hashing, Maps, and Collision Reasoning Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.c`.

## Problems

1. Build a small integer-key hash table using an array of buckets.
2. Implement `insert`, `lookup`, and `delete`.
3. Add a collision scenario where at least two keys map to the same slot.
4. Explain what happens in both chaining and open addressing.

## Run and check

```powershell
Set-Location lessons/S1/018-hashing-maps-and-collision-reasoning/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Optional challenge

### Collision-aware map

Implement a hash map with a deliberately small table so collisions occur frequently. Instrument collision counts and compare two workloads.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
