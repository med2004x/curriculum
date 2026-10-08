package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "normal output")
	fmt.Fprintln(os.Stderr, "diagnostic output")
}
