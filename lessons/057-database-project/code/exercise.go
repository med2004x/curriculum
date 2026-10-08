package main

import "fmt"

func saveOrder(id int) error {
	// TODO: define the repository boundary and durable transaction semantics.
	if id <= 0 {
		return fmt.Errorf("invalid order id")
	}
	return nil
}

func main() {
	fmt.Println(saveOrder(1))
}
