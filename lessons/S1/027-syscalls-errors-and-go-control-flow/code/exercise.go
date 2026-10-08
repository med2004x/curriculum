package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parseCount(input string) (int, error) {
	// TODO: parse, validate, and wrap failures with useful context.
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid count")
	}
	return n, nil
}

func main() {
	n, err := parseCount("42")
	fmt.Println(n, err)
}
