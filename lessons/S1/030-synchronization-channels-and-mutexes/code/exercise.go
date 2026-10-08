package main
import ("fmt"; "sync")
func runWorkers(jobs []int, workers int) []int {
 // TODO: use a bounded worker pool and preserve result order.
 var wg sync.WaitGroup
 results:=make([]int,len(jobs)); wg.Add(workers)
 for w:=0;w<workers;w++ { go func(){defer wg.Done()}() }
 wg.Wait(); return results
}
func main(){fmt.Println(runWorkers([]int{1,2,3,4},2))}
