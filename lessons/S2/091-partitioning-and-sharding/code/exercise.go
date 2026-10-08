// STUDENT PROBLEM SET — Lesson 091.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`.
// 2. Extend it with one explicit boundary case: Populate it with, say, 10,000 test keys across 4 nodes, then simulate adding a 5th node (recompute `% 5` for every key) and measure exactly what fraction of the 10,000 keys ended up mapped to a *different* node than before — report the actual percentage, and compare it to what you'd naively expect (should be a very large fraction, close to `(numNodes-1)/numN.
// 3. Rebuild the same behavior with one changed constraint: Implement consistent hashing: hash both nodes and keys onto a numeric ring (e.g. using a 32-bit or 64-bit hash function), and implement the "find the first node clockwise from the key's position" lookup (a sorted list of node positions plus a binary search, discrete math/algorithms track, is a natural implementation approach). Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 091: Partitioning and Sharding
// Your task: handle invalid shard counts and define the stable hashing contract.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "hash/fnv"

func shardFor(key string, shards int) int {
    // TODO: implement the assigned behavior.
    return 0
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ = fnv.New32a
