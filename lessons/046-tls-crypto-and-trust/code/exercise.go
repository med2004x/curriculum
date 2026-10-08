package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func fingerprint(data []byte) string {
	// TODO: use the cryptographic primitive appropriate for the exercise and explain what it does not provide.
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func main() {
	fmt.Println(fingerprint([]byte("certificate")))
}
