package main

import "encoding/binary"

func encodeFrame(payload []byte) []byte {
	// TODO: prefix the payload with a fixed-width length and return a new buffer.
	out := make([]byte, 4+len(payload))
	binary.BigEndian.PutUint32(out, uint32(len(payload)))
	copy(out[4:], payload)
	return out
}

func decodeFrame(buf []byte) ([]byte, []byte, bool) {
	// TODO: distinguish incomplete, complete, and malformed frames.
	return nil, buf, false
}

func main() {}
