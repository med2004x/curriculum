package main

import (
	"fmt"
	"sync"
)

func incrementMany(total *int, workers, increments int) {
    // TODO: implement the assigned behavior.
}

func main() {
	total := 0
	incrementMany(&total, 4, 1000)
	fmt.Println(total)
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
var _ sync.WaitGroup
