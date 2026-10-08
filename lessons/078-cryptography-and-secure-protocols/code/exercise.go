package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

func verifyMAC(key, message, supplied []byte) bool {
	// TODO: compare MACs with the constant-time primitive intended for this purpose.
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expected := mac.Sum(nil)
	return hmac.Equal(expected, supplied)
}

func main() {}
