package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// TODO: enforce the method/path contract and return structured output.
	fmt.Fprintln(w, "not implemented")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler)
	fmt.Println(http.ListenAndServe(":8080", mux))
}
