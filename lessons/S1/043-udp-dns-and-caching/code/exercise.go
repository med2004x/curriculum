package main

import "net"

func cachedLookup(cache map[string]string, name string) (string, bool) {
	// TODO: define cache hit/miss behavior and expiration policy.
	value, ok := cache[name]
	return value, ok
}

func main() {
	_, _ = net.ResolveUDPAddr("udp", "127.0.0.1:5353")
}
