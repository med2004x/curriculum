# Lesson 082 — B+ Trees and Index Structures

## Threads
**Primary:** DB Internals/Algorithms

**Supporting:** tree height; page fanout; index trade-offs

**Math companion:** logarithms

**Time:** 3–4 hours. **Prerequisites:** Lesson 081.

**Time:** 3–4 hours. **Prerequisites:** Lesson 081.

## What you will be able to do

1. State B+ tree occupancy, separator-key, and leaf-order invariants.
2. Trace lookup and insertion through internal nodes to leaves.
3. Explain leaf and internal-node splits, separator propagation, and root growth.
4. Explain why linked leaves support range scans and why index maintenance has write cost.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 081, # Lesson 081 — Pages, Records, and Buffer Pools.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Implement search and insert for a B+ tree — the data structure behind nearly every database index — and understand precisely why it's the right shape for disk-backed data, not just "a balanced tree."

**The contract.** A B+ tree keeps search keys in sorted leaf entries; internal nodes route searches using separator keys. All leaves remain at the same depth, and leaf links support ordered range scans. With fanout `F`, height grows logarithmically with the number of entries, but actual page reads depend on cache and storage state.

**Worked example.** Start with the smallest concrete case described by the source material: Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page).

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** High fanout reduces tree height compared with a binary tree, which can reduce page visits for a lookup. Node size and fanout are implementation/page-format decisions; do not assume every node is exactly one disk I/O or that every lookup reaches storage rather than cache.

**Second point.** A split redistributes entries between siblings and inserts/promotes a separator into the parent. If the root splits, a new root is created and height increases by one. Duplicate-key handling and minimum occupancy must be defined for the implementation.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page).

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** For step 2, produce a small log (or diagram) showing the exact insert sequence and which inserts triggered splits, at which node, and whether the split propagated to the parent — enough detail that someone else could verify your split logic by hand-tracing the same sequence.

**Failure mode.** Deliberately implement `rangeSearch` *without* the leaf-chain optimization — instead, for every key you want in the range, perform a fresh, independent tree search from the root.

**Completion criterion.** You can implement and verify lookup, insertion/splits, and ordered range traversal against a multi-level tree, and explain the space/write/lookup trade-offs without asserting a fixed page-read count.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 082 — B+ Trees and Index Structures case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page).

2. Extend it with one explicit boundary case: Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.

3. Rebuild the same behavior with one changed constraint: Insert keys in increasing order (1, 2, 3, 4, 5, ...) and trace, at each insert, whether a split occurred and at which level.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Deliberately implement `rangeSearch` *without* the leaf-chain optimization — instead, for every key you want in the range, perform a fresh, independent tree search from the root.

2. Measure and compare the total number of node visits for a range query spanning, say, 20 keys using this naive re-search approach versus your step 3 leaf-chain approach.

3. Report the actual difference in node-visit count, and explain why the leaf-chain approach's advantage grows with the size of the range being queried — a range of 2 keys barely matters, but a range of thousands does, which is exactly the shape of query B+ tree indexes are specifically optimized to make cheap.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 082 — B+ Trees and Index Structures: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 082 — B+ Trees and Index Structures?
2. What two concrete practices are required in the assigned material for # 082 — B+ Trees and Index Structures?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson082/EVIDENCE.md` with exactly this structure:

```text
# Lesson 082 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 082 — B+ Trees and Index Structures.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 082 — B+ Trees and Index Structures without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your B+ tree correctly implements search, insert-with-splitting, and leaf-chain range search, verified against a real multi-level tree (not just a single-node toy case), and you've directly measured and compared node-visit counts for both point lookups (versus a theoretical binary tree) and range queries (versus a naive re-search-per-key approach), connectin. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
