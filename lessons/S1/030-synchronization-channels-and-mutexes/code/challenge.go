package main

import (
	"fmt"
	"sync"
)

func runWorkers(jobs []int, workers int) []int {
	// CHALLENGE TODO: use channels and synchronization to process jobs safely.
	var wg sync.WaitGroup
	results := make([]int, len(jobs))
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
		}()
	}
	wg.Wait()
	return results
}

func main() {
	fmt.Println(runWorkers([]int{1, 2, 3, 4}, 2))
}
