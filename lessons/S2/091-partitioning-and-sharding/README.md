# Lesson 091 — Partitioning and Sharding

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Distribute data and workload across partitions while balancing ownership, query locality, and reconfiguration cost. Consistent hashing can reduce movement for a ring-based assignment, but range/hash partitioning, directory maps, and database-native partitioning may be better for other workloads.

**The contract.** `hash(key) % M` can remap many keys when `M` changes. Consistent hashing places nodes and keys on a ring so adding/removing a node typically moves keys in the affected arc, subject to hash distribution and virtual-node configuration. Movement is not zero and imbalance remains possible.

**Worked example.** Start with the smallest concrete case described by the source material: Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`.

## 2. Mechanism and vocabulary

**Mechanism.** Range partitioning can support range scans but risks hotspots when key distributions are skewed. Hash partitioning tends to spread keys but can make range queries scatter. Consistent hashing trades straightforward routing for ring metadata, virtual nodes, and explicit data movement/rebalancing.

**Second point.** A partition key affects not only ownership balance but also transaction locality, query fan-out, and hot-key behavior. Rebalancing must move data safely and coordinate ownership metadata; changing the lookup function alone does not migrate records.

**Worked example.** Populate it with, say, 10,000 test keys across 4 nodes, then simulate adding a 5th node (recompute `% 5` for every key) and measure exactly what fraction of the 10,000 keys ended up mapped to a *different* node than before — report the actual percentage, and compare it to what you'd naively expect (should be a very large fraction, close to `(numNodes-1)/numN.

## 3. Worked case from the reference material

Worked example: Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`.

## 4. Verification, failure, and evidence

**Verification.** Produce a table comparing: naive modulo (step 1), plain consistent hashing (steps 2-3), and consistent hashing with virtual nodes (step 4) — for each, report the percentage of keys that moved when adding the 5th node, and the evenness of the resulting distribution across nodes.

**Failure mode.** Implement consistent hashing with a deliberately *small* number of virtual nodes per physical node (e.g. just 2, instead of 150) and rerun the distribution-evenness measurement from step 4.

**Completion criterion.** You can compare movement and balance for stated strategies, explain virtual-node trade-offs, and identify query/locality risks of the chosen partition key.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
