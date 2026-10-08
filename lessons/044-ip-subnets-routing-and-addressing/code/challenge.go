package main

import (
	"fmt"
	"net"
)

func nextHop(routes []net.IPNet, destination net.IP) int {
	// CHALLENGE TODO: choose the most specific matching route. Return -1 when none matches.
	_ = routes
	_ = destination
	return -1
}

func main() {
	_, network, _ := net.ParseCIDR("10.0.0.0/8")
	fmt.Println(nextHop([]net.IPNet{*network}, net.ParseIP("10.20.30.40")))
}
