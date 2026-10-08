// STUDENT PROBLEM SET — Lesson 078.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Compute SHA-256 hashes of several inputs (using your language's standard crypto library, never a hand-rolled implementation) and confirm the hash changes completely for even a single-character input change (the "avalanche effect" — a defining property of a good cryptographic hash), demonstrating this isn't a simple checksum where small input changes produce.
// 2. Extend it with one explicit boundary case: Implement HMAC-based message signing and verification (extending API engineering Lesson 11's practical webhook use directly), and confirm that altering either the message or the signature (but not both consistently) causes verification to correctly fail — test both cases explicitly, not just the happy path.
// 3. Generate a signing key pair using an established library, sign a message, verify with the correct public key, then verify failure with an unrelated public key. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 078: Cryptography and Secure Protocols
// Your task: compare MACs with the constant-time primitive intended for this purpose.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

func verifyMAC(key, message, supplied []byte) bool {
    // TODO: implement the assigned behavior.
    return false
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ = hmac.New
var _ = sha256.New
