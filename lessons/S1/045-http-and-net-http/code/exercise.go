// STUDENT PROBLEM SET — Lesson 045.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Using raw TCP sockets (Lesson 1, not an HTTP library), implement a minimal HTTP server that parses an incoming request line and headers (up to the blank-line terminator), and responds with a hardcoded `200 OK` response and a small body, including a correct `Content-Length` header matching the actual body byte length.
// 2. Extend it with one explicit boundary case: Test your server with a real client (`curl -v` shows the raw request/response headers) and confirm the response is correctly interpreted — `curl` should display your status code and body without error, and `curl -v`'s verbose output should show the exact bytes exchanged, letting you confirm your hand-built response matches the format real clients expect.
// 3. Rebuild the same behavior with one changed constraint: Extend your server to support keep-alive: after responding, don't close the connection — loop back and parse another request on the same TCP connection, unless the client's request included `Connection: close`. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 045: HTTP and net/http
// Your task: enforce the method/path contract and return structured output. fmt.Fprintln(w, "implement the required behavior")
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // TODO: implement the assigned behavior.
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler)
	fmt.Println(http.ListenAndServe(":8080", mux))
}

// Keep starter imports available while the student implementation is unfinished.
var _ = fmt.Fprintln
var _ = http.ResponseWriter
