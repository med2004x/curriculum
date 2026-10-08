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
	// TODO: decode/validate input, call the application layer, and encode JSON.
	_ = json.NewEncoder(w).Encode([]Product{})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", productsHandler)
	_ = http.ListenAndServe(":8080", mux)
}
