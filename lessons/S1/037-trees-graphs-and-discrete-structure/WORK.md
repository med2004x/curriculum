# Lesson 037 — Trees, Graphs, and Discrete Structure Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement BST insertion and inorder traversal; check its invariant after each insertion.

2. Represent a small graph as both an adjacency list and an adjacency matrix; implement neighbor iteration.

3. Implement BFS with a queue and a visited set; compare its visitation order with DFS on the same graph.

4. Given a DAG, produce a topological order and verify every directed edge respects it.

## Run and check

```powershell
Set-Location lessons/S1/037-trees-graphs-and-discrete-structure/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Graph service

Represent a directed graph and implement reachability plus cycle detection. Provide a counterexample for a naive implementation.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
