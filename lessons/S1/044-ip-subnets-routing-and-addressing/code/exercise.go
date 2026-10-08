package main

import (
	"fmt"
	"net"
)

func sameSubnet(network, host string) (bool, error) {
	// TODO: parse the CIDR network and host, then check membership.
	_, _, err := net.ParseCIDR(network)
	if err != nil {
		return false, err
	}
	_, _, err = net.ParseCIDR(host)
	if err != nil {
		return false, err
	}
	return false, nil
}

func main() {
	ok, err := sameSubnet("192.168.1.0/24", "192.168.1.42")
	fmt.Println(ok, err)
}
