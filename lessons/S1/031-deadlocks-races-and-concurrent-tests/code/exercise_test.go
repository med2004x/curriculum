package main
import ("sync"; "testing"; "time")
func TestWithLocksOppositeOrderCompletes(t *testing.T) {
 var a,b sync.Mutex; const workers=40; done:=make(chan struct{},workers)
 for i:=0;i<workers;i++ { go func(reverse bool){ if reverse { withLocks(&b,&a,func(){}) } else { withLocks(&a,&b,func(){}) }; done<-struct{}{} }(i%2==0) }
 for i:=0;i<workers;i++ { select {case <-done: case <-time.After(2*time.Second): t.Fatal("lock acquisition deadlocked") } }
}
