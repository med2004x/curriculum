package main

import ("reflect"; "testing")
func TestRunWorkers(t *testing.T) {
 for _, workers := range []int{-1, 1, 2, 9} {
  got := runWorkers([]int{3, -2, 0, 5}, workers)
  want := []int{9, 4, 0, 25}
  if !reflect.DeepEqual(got, want) { t.Errorf("workers=%d: got %v, want %v", workers, got, want) }
 }
}
func TestRunWorkersEmpty(t *testing.T) { if got:=runWorkers(nil, 4); len(got)!=0 { t.Fatalf("got %v",got) } }
