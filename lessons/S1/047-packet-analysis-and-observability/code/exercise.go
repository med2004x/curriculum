// STUDENT PROBLEM SET — Lesson 047.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets b.
// 2. Extend it with one explicit boundary case: In the same capture, identify the first HTTP request packet (containing the actual `GET`/`POST` line) and the first response packet (containing `HTTP/1.1 ...`), and compute the time delta between them — this is your directly measured "server think time" for that specific request.
// 3. Rebuild the same behavior with one changed constraint: Deliberately introduce artificial packet loss to observe retransmission directly: using a tool like `tc` (Linux traffic control, e.g. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 047: Packet Analysis and Observability
// Your task: validate length/version/header size before reading fields.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"encoding/binary"
	"fmt"
)

func ipv4Summary(packet []byte) (string, bool) {
    // TODO: implement the assigned behavior.
    return "", false
}

func netIPv4(b []byte) string { return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3]) }

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ = binary.BigEndian
var _ = fmt.Sprintf
