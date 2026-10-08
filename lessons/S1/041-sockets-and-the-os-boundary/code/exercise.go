// STUDENT PROBLEM SET — Lesson 041.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Write a minimal TCP server (bind, listen, accept) and client (connect) in Go or C that exchanges a single short message.
// 2. Extend it with one explicit boundary case: Confirm it works for a small message.
// 3. Rebuild the same behavior with one changed constraint: Demonstrate the "no message boundaries" fact directly: modify the client to send two separate messages via two separate `write()`/`Write()` calls in quick succession (e.g. "hello" then "world", with no delay between them), and have the server do a single `read()` immediately after accepting the connection, printing exactly what it received in that one read. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 041: Sockets and the OS Boundary
// Your task: read one request and write one correct response.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main
import ("fmt"; "net")
func handle(conn net.Conn) error {
    // TODO: implement the assigned behavior.
    return nil
}
func main(){ln,err:=net.Listen("tcp",":9000");if err!=nil{panic(err)};defer ln.Close();fmt.Println("listening on :9000");conn,err:=ln.Accept();if err!=nil{panic(err)};_ = handle(conn)}
