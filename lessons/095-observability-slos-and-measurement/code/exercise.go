package main

import "time"

type Window struct {
	Good  int
	Total int
	Start time.Time
}

func record(w *Window, ok bool) {
	// TODO: keep an SLI window and make the measurement semantics explicit.
	w.Total++
	if ok {
		w.Good++
	}
}

func main() {}
