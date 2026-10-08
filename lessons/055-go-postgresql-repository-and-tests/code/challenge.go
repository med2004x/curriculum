package main

import "context"

type Repository interface {
	Create(ctx context.Context, title string) error
}

func createWithRetry(ctx context.Context, repo Repository, title string) error {
	// CHALLENGE TODO: define which database failures are retryable and keep the operation safe.
	return nil
}

func main() {}
