package main
import ("fmt"; "sync")
func incrementMany(total *int, workers, increments int) {
 // TODO: protect the shared update and wait for all workers.
 var wg sync.WaitGroup
 wg.Add(workers)
 for i:=0;i<workers;i++ { go func(){ defer wg.Done(); for j:=0;j<increments;j++ { *total++ } }() }
 wg.Wait()
}
func main(){ total:=0; incrementMany(&total,4,1000); fmt.Println(total) }
