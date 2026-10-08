package main

import "fmt"

type Item struct {
	Name  string
	Value int
}

func total(items []Item) int {
	// TODO: compute the aggregate without mutating the input.
	return 0
}

func main() {
	items := []Item{{Name: "a", Value: 2}, {Name: "b", Value: 3}}
	fmt.Println(total(items))
}
