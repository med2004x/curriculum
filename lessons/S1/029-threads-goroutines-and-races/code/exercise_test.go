package main

import "testing"

func TestIncrementMany(t *testing.T) {
 total := 0
 incrementMany(&total, 8, 1500)
 if total != 12000 { t.Fatalf("total = %d, want 12000", total) }
}
func TestIncrementManyIgnoresInvalidWork(t *testing.T) {
 total := 7
 incrementMany(&total, 0, 10)
 incrementMany(&total, 2, 0)
 incrementMany(nil, 2, 10)
 if total != 7 { t.Fatalf("total = %d, want unchanged 7", total) }
}
