/* WORKED EXAMPLE: A runnable race demonstration; implement the synchronized task in exercise.go. */
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex
    total := 0
    const workers = 4
    const increments = 1000
    wg.Add(workers)
    for i := 0; i < workers; i++ {
        go func() {
            defer wg.Done()
            for j := 0; j < increments; j++ {
                mu.Lock()
                total++
                mu.Unlock()
            }
        }()
    }
    wg.Wait()
    fmt.Println(total)
}
