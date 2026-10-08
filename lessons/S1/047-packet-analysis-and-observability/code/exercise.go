package main

import (
	"encoding/binary"
	"fmt"
)

func ipv4Summary(packet []byte) (string, bool) {
	// TODO: validate length/version/header size before reading fields.
	if len(packet) < 20 || packet[0]>>4 != 4 {
		return "", false
	}
	src := netIPv4(packet[12:16])
	dst := netIPv4(packet[16:20])
	_ = binary.BigEndian.Uint16(packet[2:4])
	return fmt.Sprintf("%s -> %s", src, dst), true
}

func netIPv4(b []byte) string { return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3]) }

func main() {}
