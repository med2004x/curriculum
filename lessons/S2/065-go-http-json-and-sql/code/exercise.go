// STUDENT PROBLEM SET — Lesson 065.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Define typed create-request and response structs and a bounded JSON decoder.
// 2. Validate the request and call a context-aware repository method using parameters.
// 3. Map success, malformed input, invalid fields, not-found, and internal error to distinct outcomes. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 065: Go HTTP, JSON, and SQL
// Your task: decode/validate input, call the application layer, and encode JSON.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: implement the assigned behavior.
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", productsHandler)
	_ = http.ListenAndServe(":8080", mux)
}

// Keep starter imports available while the student implementation is unfinished.
var _ = json.NewEncoder
var _ = http.ResponseWriter
