package main
import "sync"
var pairGate sync.Mutex
func withLocks(a,b *sync.Mutex, work func()) {
 // TODO: replace the global gate with a stable lock ordering to permit safe concurrency.
 pairGate.Lock()
 a.Lock(); b.Lock(); work(); b.Unlock(); a.Unlock()
 pairGate.Unlock()
}
func deliberatelyBroken(){ var a,b sync.Mutex; withLocks(&a,&b,func(){}) }
func main(){deliberatelyBroken()}
