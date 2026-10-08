package main
import ("io"; "net"; "testing")
func TestHandleRoundTrip(t *testing.T) {
 server, client := net.Pipe(); done:=make(chan error,1)
 go func(){ done<-handle(server) }()
 if _,err:=io.WriteString(client,"ping\n"); err!=nil { t.Fatal(err) }
 got,err:=io.ReadAll(client); if err!=nil { t.Fatal(err) }
 if string(got)!="received: ping\n" { t.Fatalf("response %q",got) }
 if err:=<-done; err!=nil { t.Fatal(err) }
}
