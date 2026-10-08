// STUDENT PROBLEM SET — Lesson 042.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Annotate a TCP handshake trace with both initial sequence numbers and acknowledgment values.
// 2. Annotate a data trace with byte ranges, cumulative acknowledgments, and any retransmission.
// 3. Compare a receiver-window change with a congestion-window change and explain which endpoint or network condition each represents. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 042: TCP State, Reliability, and Streams
// Your task: prefix the payload with a fixed-width length and return a new buffer. distinguish incomplete, complete, and malformed frames.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "encoding/binary"

func encodeFrame(payload []byte) []byte {
    // TODO: implement the assigned behavior.
    return nil
}

func decodeFrame(buf []byte) ([]byte, []byte, bool) {
    // TODO: implement the assigned behavior.
    return nil, nil, false
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ = binary.BigEndian
