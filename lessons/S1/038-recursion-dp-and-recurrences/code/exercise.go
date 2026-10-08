package main

import "fmt"

func Fib(n int, memo map[int]int) int {
	// TODO: use the recurrence and memoization to avoid repeated work.
	if n < 2 {
		return n
	}
	return 0
}

func main() {
	fmt.Println(Fib(10, map[int]int{}))
}
