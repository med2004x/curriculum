package main

import (
	"fmt"
	"strings"
)

func normalize(input string) string {
	// TODO: implement the input transformation from the lesson.
	return strings.TrimSpace(input)
}

func main() {
	fmt.Println(normalize("  process  "))
}
