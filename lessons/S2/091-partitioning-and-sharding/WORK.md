# Lesson 091 — Problem Set

**Topic:** Partitioning and Sharding.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`.

2. Extend it with one explicit boundary case: Populate it with, say, 10,000 test keys across 4 nodes, then simulate adding a 5th node (recompute `% 5` for every key) and measure exactly what fraction of the 10,000 keys ended up mapped to a *different* node than before — report the actual percentage, and compare it to what you'd naively expect (should be a very large fraction, close to `(numNodes-1)/numN.

3. Rebuild the same behavior with one changed constraint: Implement consistent hashing: hash both nodes and keys onto a numeric ring (e.g. using a 32-bit or 64-bit hash function), and implement the "find the first node clockwise from the key's position" lookup (a sorted list of node positions plus a binary search, discrete math/algorithms track, is a natural implementation approach).

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/091-partitioning-and-sharding/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Sharded keyspace

Partition a keyspace across nodes, route requests, move one partition, and document what happens to ownership during the move.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
