package main

import "fmt"

func BFS(graph map[int][]int, start int) []int {
	// TODO: implement breadth-first traversal with a queue and visited set.
	return nil
}

func main() {
	graph := map[int][]int{1: {2, 3}, 2: {4}, 3: {4}}
	fmt.Println(BFS(graph, 1))
}
