package main

import "time"

type Bucket struct {
	tokens  float64
	updated time.Time
}

func allow(b *Bucket, now time.Time, rate, burst float64) bool {
	// TODO: implement refill, cap, and atomic decision semantics.
	return false
}

func main() {}
