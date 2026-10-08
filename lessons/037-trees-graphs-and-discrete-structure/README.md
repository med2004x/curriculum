# Lesson 037 — Trees, Graphs, and Discrete Structure

## Threads
**Primary:** Algorithms/Math

**Supporting:** graph representations; traversal

**Math companion:** graph theory

**Time:** 3–4 hours. **Prerequisites:** Lesson 036.

**Time:** 3–4 hours. **Prerequisites:** Lesson 036.

## What you will be able to do

1. State the binary-search-tree invariant and connect inorder traversal to sorted output.
2. Compare adjacency-list and adjacency-matrix graph representations.
3. Trace breadth-first and depth-first traversal while marking visited vertices.
4. Explain when a plain BST degenerates and how graph cycles affect traversal or topological ordering.

## Retrieve first (no notes)

1. What precondition does binary search require?
2. What is the postcondition of a correct sorting algorithm?
3. Why must a binary-search interval shrink on each iteration?
4. What does a linear scan cost for `n` elements?

## 1. Problem and model

**What problem this solves.** Represent hierarchical relationships and use ordering to improve search.

**The contract.** In a BST with unique keys, every key in a node's left subtree is smaller than the node key and every key in its right subtree is larger. Inorder traversal visits left subtree, node, then right subtree, producing ascending keys under that invariant. A graph is `G=(V,E)`; edges may be directed/undirected and may carry weights.

**Worked example.** Insert `8, 3, 10, 1, 6` into a BST. The structure is rooted at 8; inorder traversal emits `1, 3, 6, 8, 10`. Trace each insertion using the left-smaller/right-larger rule.

Self-check: if values may be duplicated, define whether duplicates go left, go right, or are stored as a count; the invariant and tests must match that policy.

## 2. Mechanism and vocabulary

**Mechanism.** An adjacency list stores each vertex's neighbors and typically uses O(V+E) space. An adjacency matrix uses O(V²) space and provides direct edge lookup; it can suit dense graphs. BFS uses a queue; DFS uses recursion or an explicit stack. Both need a visited set for general graphs.

**Second point.** A topological ordering is a linear ordering of a directed acyclic graph (DAG) in which each edge `u → v` places `u` before `v`. If the graph contains a directed cycle, no such ordering exists. Do not confuse this dependency ordering with BST inorder traversal.

For a graph with `V` vertices and `E` edges, adjacency-list traversal is O(V+E) when each vertex and edge is processed a constant number of times. A plain BST operation is O(h), where `h` is tree height; a degenerate tree has `h=O(n)`.

**Worked example.** Represent the same graph in two ways and implement neighbor iteration.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement a BST and report values in sorted order with an inorder traversal.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Verify every BST insertion preserves the ordering invariant and that inorder output is sorted. For graph traversals, verify each reachable vertex is visited once. For a DAG, check every edge against the produced topological order; a cyclic input must be reported as having no topological order.

**Failure mode.** Insert already sorted keys into an unbalanced BST and explain why height can grow to `n`. Separately remove the visited set from graph traversal and show why an undirected cycle can cause repeated traversal.

**Completion criterion.** You can state the BST invariant, trace traversal orders, select a graph representation for density, and explain visited-set and DAG-cycle behavior.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the inorder traversal of the example BST.
2. Predict the BST height after sorted keys are inserted without balancing.
3. Predict how adjacency-list space changes as edges are added.
4. Predict which representation is suitable for a dense graph and why.
5. Predict how BFS and DFS differ in frontier structure.
6. Predict what happens when a visited set is omitted on a cyclic graph.
7. Predict whether a directed cycle can have a topological order.
8. State the invariant or graph property each prediction uses.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement BST insertion and inorder traversal; check its invariant after each insertion.

2. Represent a small graph as both an adjacency list and an adjacency matrix; implement neighbor iteration.

3. Implement BFS with a queue and a visited set; compare its visitation order with DFS on the same graph.

4. Given a DAG, produce a topological order and verify every directed edge respects it.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Insert already sorted input and explain why a plain BST can degrade toward linear behavior.

2. Add one edge to your DAG from step 3 that creates a cycle (e.g., make the "last" package in your dependency chain depend back on the "first" one).

3. Attempt topological sort again by hand using the same removal process, and observe that at some point, no remaining vertex has zero incoming (or outgoing) edges — every remaining vertex has at least one dependency among the remaining set.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 037 — Trees, Graphs, and Discrete Structure: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. State the BST invariant for a node and its subtrees.
2. Why does inorder traversal produce sorted keys in a valid BST?
3. What are the space costs of adjacency list and matrix representations?
4. What data structure does BFS use for its frontier?
5. Why does graph traversal need a visited set?
6. What condition must hold for a topological order to exist?

## Evidence to record

Create `lesson037/EVIDENCE.md` with exactly this structure:

```text
# Lesson 037 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Confusing BST inorder traversal with topological ordering.
- Assuming a plain BST remains logarithmic after sorted insertion.
- Traversing a cyclic graph without a visited set.
- Choosing an adjacency matrix for a sparse graph without considering O(V²) storage.
- Reporting a topological order without checking each directed edge.

## Precision checks

1. State the BST invariant and height-dependent cost.
2. Trace inorder, preorder, and postorder on one small tree.
3. Compare adjacency-list and adjacency-matrix space requirements.
4. Explain BFS's queue and visited-set roles.
5. Check a topological ordering edge by edge and identify cycle failure.

Trees impose hierarchy; graphs model general relationships. Their invariants and traversal rules are related but not interchangeable.
## University revision bridge

For the LIG track, connect BST ordering to recursive structures and graph traversal to adjacency, reachability, degree, and DAG reasoning.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can verify a BST invariant and traversal, choose and implement a graph representation, trace BFS/DFS with a visited set, and validate or reject a topological ordering based on directed cycles.
