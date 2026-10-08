# Lesson 091 — Partitioning and Sharding Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`.

2. Populate it with, say, 10,000 test keys across 4 nodes, then simulate adding a 5th node (recompute `% 5` for every key) and measure exactly what fraction of the 10,000 keys ended up mapped to a *different* node than before — report the actual percentage, and compare it to what you'd naively expect (should be a very large fraction, close to `(numNodes-1)/numN.

3. Implement consistent hashing: hash both nodes and keys onto a numeric ring (e.g. using a 32-bit or 64-bit hash function), and implement the "find the first node clockwise from the key's position" lookup (a sorted list of node positions plus a binary search, discrete math/algorithms track, is a natural implementation approach).

## Run and check

```powershell
Set-Location lessons/S2/091-partitioning-and-sharding/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Sharded keyspace

Partition a keyspace across nodes, route requests, move one partition, and document what happens to ownership during the move.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
