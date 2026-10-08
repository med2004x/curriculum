package main
import ("context"; "sync"; "testing")
func TestMemoryRepositoryCreateGet(t *testing.T) {
 r:=&MemoryRepository{}; task,err:=r.Create(context.Background(),"  ship it "); if err!=nil {t.Fatal(err)}
 if task.ID!=1 || task.Title!="ship it" {t.Fatalf("task=%+v",task)}
 got,err:=r.Get(context.Background(),task.ID); if err!=nil || got!=task {t.Fatalf("got=%+v err=%v",got,err)}
}
func TestMemoryRepositoryValidationAndCancellation(t *testing.T) {
 r:=&MemoryRepository{}; if _,err:=r.Create(context.Background()," "); err==nil {t.Fatal("empty title accepted")}
 ctx,cancel:=context.WithCancel(context.Background()); cancel()
 if _,err:=r.Get(ctx,1); err==nil {t.Fatal("canceled context accepted")}
}
func TestMemoryRepositoryConcurrentIDs(t *testing.T) {
 r:=&MemoryRepository{}; var wg sync.WaitGroup; ids:=make(chan int64,64)
 for i:=0;i<64;i++ {wg.Add(1); go func(){defer wg.Done(); task,err:=r.Create(context.Background(),"x"); if err!=nil {t.Error(err); return}; ids<-task.ID}()}
 wg.Wait(); close(ids); seen:=map[int64]bool{}; for id:=range ids {if seen[id] {t.Fatalf("duplicate id %d",id)}; seen[id]=true}
 if len(seen)!=64 {t.Fatalf("got %d IDs",len(seen))}
}
