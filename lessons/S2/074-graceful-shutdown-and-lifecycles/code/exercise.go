package main

import (
	"context"
	"net/http"
)

func runServer(ctx context.Context, srv *http.Server) error {
	// TODO: coordinate listener, shutdown, deadlines, and returned errors.
	return srv.ListenAndServe()
}

func main() {}
