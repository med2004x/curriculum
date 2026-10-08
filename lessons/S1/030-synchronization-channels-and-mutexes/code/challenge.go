package main

import (
	"fmt"
	"sync"
)

func runWorkers(jobs []int, workers int) []int {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {
	fmt.Println(runWorkers([]int{1, 2, 3, 4}, 2))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
var _ sync.WaitGroup
