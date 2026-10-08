package main

import (
	"fmt"
	"net"
)

func nextHop(routes []net.IPNet, destination net.IP) int {
    // TODO: implement the assigned behavior.
    return 0
}

func main() {
	_, network, _ := net.ParseCIDR("10.0.0.0/8")
	fmt.Println(nextHop([]net.IPNet{*network}, net.ParseIP("10.20.30.40")))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
var _ = net.IPNet
