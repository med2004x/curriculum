package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("listening on :9000")
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	_ = handle(conn)
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
var _ net.Conn
