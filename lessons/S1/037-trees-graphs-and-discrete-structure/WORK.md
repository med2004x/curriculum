# Lesson 037 — Problem Set

**Topic:** Trees, Graphs, and Discrete Structure.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Implement BST insertion and inorder traversal; check its invariant after each insertion.

2. Represent a small graph as both an adjacency list and an adjacency matrix; implement neighbor iteration.

3. Implement BFS with a queue and a visited set; compare its visitation order with DFS on the same graph.

4. Given a DAG, produce a topological order and verify every directed edge respects it.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/037-trees-graphs-and-discrete-structure/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Graph service

Represent a directed graph and implement reachability plus cycle detection. Provide a counterexample for a naive implementation.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
