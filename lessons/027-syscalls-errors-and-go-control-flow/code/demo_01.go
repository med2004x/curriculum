package main

import (
	"fmt"
	"strconv"
)

func parseCount(text string) (int, error) {
	count, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("parse count %q: %w", text, err)
	}
	return count, nil
}
