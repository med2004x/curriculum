package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) })
	// TODO: add build metadata, graceful shutdown, and a release-safe health check.
	_ = http.ListenAndServe(":8080", mux)
}
