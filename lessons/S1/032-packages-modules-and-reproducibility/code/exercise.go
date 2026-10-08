package main

import "fmt"

func makeLabel(name string, n int) string {
	// TODO: move reusable logic into a package boundary and make the build reproducible.
	return fmt.Sprintf("%s:%d", name, n)
}

func main() {
	fmt.Println(makeLabel("lesson", 32))
}
