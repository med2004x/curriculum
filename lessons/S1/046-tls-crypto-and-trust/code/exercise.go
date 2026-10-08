// STUDENT PROBLEM SET — Lesson 046.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Generate a local, self-signed TLS certificate (e.g.
// 2. Extend it with one explicit boundary case: `openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 1 -nodes`) and use it to run a minimal HTTPS server (Go's `net/http` with `ListenAndServeTLS`, pointing at your generated cert/key, is straightforward for this).
// 3. Rebuild the same behavior with one changed constraint: Connect to it with `curl -k` (the `-k` flag skips certificate verification, appropriate here since you know it's self-signed and you're testing locally) and confirm the connection succeeds and is encrypted. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 046: TLS, Cryptography, and Trust
// Your task: use the cryptographic primitive appropriate for the exercise and explain what it does not provide.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func fingerprint(data []byte) string {
    // TODO: implement the assigned behavior.
    return ""
}

func main() {
	fmt.Println(fingerprint([]byte("certificate")))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = sha256.Sum256
var _ = hex.EncodeToString
var _ = fmt.Println
