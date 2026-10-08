# Lesson 082 — B+ Trees and Index Structures

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Implement search and insert for a B+ tree — the data structure behind nearly every database index — and understand precisely why it's the right shape for disk-backed data, not just "a balanced tree."

**The contract.** A B+ tree keeps search keys in sorted leaf entries; internal nodes route searches using separator keys. All leaves remain at the same depth, and leaf links support ordered range scans. With fanout `F`, height grows logarithmically with the number of entries, but actual page reads depend on cache and storage state.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page).

## 2. Mechanism and vocabulary

**Mechanism.** High fanout reduces tree height compared with a binary tree, which can reduce page visits for a lookup. Node size and fanout are implementation/page-format decisions; do not assume every node is exactly one disk I/O or that every lookup reaches storage rather than cache.

**Second point.** A split redistributes entries between siblings and inserts/promotes a separator into the parent. If the root splits, a new root is created and height increases by one. Duplicate-key handling and minimum occupancy must be defined for the implementation.

**Worked example.** Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.

## 3. Worked case from the reference material

Worked example: Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page).

## 4. Verification, failure, and evidence

**Verification.** For step 2, produce a small log (or diagram) showing the exact insert sequence and which inserts triggered splits, at which node, and whether the split propagated to the parent — enough detail that someone else could verify your split logic by hand-tracing the same sequence.

**Failure mode.** Deliberately implement `rangeSearch` *without* the leaf-chain optimization — instead, for every key you want in the range, perform a fresh, independent tree search from the root.

**Completion criterion.** You can implement and verify lookup, insertion/splits, and ordered range traversal against a multi-level tree, and explain the space/write/lookup trade-offs without asserting a fixed page-read count.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
