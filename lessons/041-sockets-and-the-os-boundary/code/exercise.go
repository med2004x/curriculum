package main
import ("fmt"; "net")
func handle(conn net.Conn) error {
 defer conn.Close()
 // TODO: read one request and write one correct response.
 return nil
}
func main(){ln,err:=net.Listen("tcp",":9000");if err!=nil{panic(err)};defer ln.Close();fmt.Println("listening on :9000");conn,err:=ln.Accept();if err!=nil{panic(err)};_ = handle(conn)}
