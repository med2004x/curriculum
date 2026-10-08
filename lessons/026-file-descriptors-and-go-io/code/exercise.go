package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) error {
	// TODO: open both files, copy bytes, and close descriptors correctly.
	return os.ErrInvalid
}

func main() {
	if err := copyFile("input.txt", "output.txt"); err != nil && err != os.ErrInvalid {
		fmt.Println(err)
	}
	_ = io.EOF
}
