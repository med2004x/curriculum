package main

import "fmt"

func containsString(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func containsInt(values []int, wanted int) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsString([]string{"go", "rust"}, "go"))
	fmt.Println(containsInt([]int{1, 2, 3}, 4))
}
