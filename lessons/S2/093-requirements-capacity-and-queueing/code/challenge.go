package main

import "fmt"

func requiredServers(arrivalRate, serviceRate, targetUtilization float64) int {
	// CHALLENGE TODO: choose the capacity model, state its assumptions, and round up correctly.
	return 0
}

func main() {
	fmt.Println(requiredServers(800, 120, 0.7))
}
