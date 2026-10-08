package main

import (
	"fmt"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	// TODO: implement the POST contract and status semantics.
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "not implemented")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", createProduct)
	_ = http.ListenAndServe(":8080", mux)
}
