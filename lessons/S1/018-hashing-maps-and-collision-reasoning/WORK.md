# Lesson 018 — Problem Set

**Topic:** Hashing, Maps, and Collision Reasoning.
**Your files:** `code/exercise.c`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build a small integer-key hash table using an array of buckets.
2. Implement `insert`, `lookup`, and `delete`.
3. Add a collision scenario where at least two keys map to the same slot.
4. Explain what happens in both chaining and open addressing.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/018-hashing-maps-and-collision-reasoning/code
gcc -std=c17 -Wall -Wextra -Wpedantic exercise.c -o exercise.exe
.\exercise.exe
```

## Stretch problem

### Collision-aware map

Implement a hash map with a deliberately small table so collisions occur frequently. Instrument collision counts and compare two workloads.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
