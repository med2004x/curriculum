package main

import "fmt"

func imageDigest(contents []byte) string {
	// TODO: replace this toy identity with the digest and provenance check required by the lesson.
	return fmt.Sprintf("bytes=%d", len(contents))
}

func main() {}
