package main

type Event struct {
	Index int
	Value int
}

func apply(state int, e Event) int {
	// TODO: define a deterministic state transition and reject invalid log order.
	return state + e.Value
}

func main() {}
