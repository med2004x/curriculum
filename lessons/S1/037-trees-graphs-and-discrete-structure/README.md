# Lesson 037 — Trees, Graphs, and Discrete Structure

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Represent hierarchical relationships and use ordering to improve search.

**The contract.** In a BST with unique keys, every key in a node's left subtree is smaller than the node key and every key in its right subtree is larger. Inorder traversal visits left subtree, node, then right subtree, producing ascending keys under that invariant. A graph is `G=(V,E)`; edges may be directed/undirected and may carry weights.

**Worked example.** Insert `8, 3, 10, 1, 6` into a BST. The structure is rooted at 8; inorder traversal emits `1, 3, 6, 8, 10`. Trace each insertion using the left-smaller/right-larger rule.

## 2. Mechanism and vocabulary

**Mechanism.** An adjacency list stores each vertex's neighbors and typically uses O(V+E) space. An adjacency matrix uses O(V²) space and provides direct edge lookup; it can suit dense graphs. BFS uses a queue; DFS uses recursion or an explicit stack. Both need a visited set for general graphs.

**Second point.** A topological ordering is a linear ordering of a directed acyclic graph (DAG) in which each edge `u → v` places `u` before `v`. If the graph contains a directed cycle, no such ordering exists. Do not confuse this dependency ordering with BST inorder traversal.

For a graph with `V` vertices and `E` edges, adjacency-list traversal is O(V+E) when each vertex and edge is processed a constant number of times. A plain BST operation is O(h), where `h` is tree height; a degenerate tree has `h=O(n)`.

**Worked example.** Represent the same graph in two ways and implement neighbor iteration.

## 3. Worked case from the reference material

Worked example: Implement a BST and report values in sorted order with an inorder traversal.

## 4. Verification, failure, and evidence

**Verification.** Verify every BST insertion preserves the ordering invariant and that inorder output is sorted. For graph traversals, verify each reachable vertex is visited once. For a DAG, check every edge against the produced topological order; a cyclic input must be reported as having no topological order.

**Failure mode.** Insert already sorted keys into an unbalanced BST and explain why height can grow to `n`. Separately remove the visited set from graph traversal and show why an undirected cycle can cause repeated traversal.

**Completion criterion.** You can state the BST invariant, trace traversal orders, select a graph representation for density, and explain visited-set and DAG-cycle behavior.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
