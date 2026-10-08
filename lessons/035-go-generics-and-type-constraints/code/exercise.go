package main

import "fmt"

func Max[T ~int | ~float64](values []T) (T, bool) {
	// TODO: implement the empty-input contract and generic comparison.
	var zero T
	return zero, false
}

func main() {
	v, ok := Max([]int{3, 7, 2})
	fmt.Println(v, ok)
}
