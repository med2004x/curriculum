package main

import "context"

type Repository interface {
	Create(ctx context.Context, title string) error
}

func createWithRetry(ctx context.Context, repo Repository, title string) error {
    // TODO: implement the assigned behavior.
    return nil
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ context.Context
