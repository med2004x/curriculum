// STUDENT PROBLEM SET — Lesson 037.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Implement BST insertion and inorder traversal; check its invariant after each insertion.
// 2. Represent a small graph as both an adjacency list and an adjacency matrix; implement neighbor iteration.
// 3. Implement BFS with a queue and a visited set; compare its visitation order with DFS on the same graph.
// 4. Given a DAG, produce a topological order and verify every directed edge respects it. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 037: Trees, Graphs, and Discrete Structure
// Your task: implement breadth-first traversal with a queue and visited set.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "fmt"

func BFS(graph map[int][]int, start int) []int {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {
	graph := map[int][]int{1: {2, 3}, 2: {4}, 3: {4}}
	fmt.Println(BFS(graph, 1))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
