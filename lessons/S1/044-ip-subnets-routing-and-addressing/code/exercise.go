// STUDENT PROBLEM SET — Lesson 044.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Given a small network diagram, calculate subnets and determine which gateway handles several destinations.
// 2. Extend it with one explicit boundary case: Inspect a real routing table on Linux.
// 3. Rebuild the same behavior with one changed constraint: Write a second case with one changed constraint and verify it. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 044: IP Addressing, Subnets, and Routing
// Your task: parse the CIDR network and host, then check membership.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"fmt"
	"net"
)

func sameSubnet(network, host string) (bool, error) {
    // TODO: implement the assigned behavior.
    return false, nil
}

func main() {
	ok, err := sameSubnet("192.168.1.0/24", "192.168.1.42")
	fmt.Println(ok, err)
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Println
var _ = net.ParseCIDR
