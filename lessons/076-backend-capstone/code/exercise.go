package main

import "context"

type Service struct{}

func (Service) Create(ctx context.Context, name string) error {
	// TODO: implement the complete service path: validation, persistence, HTTP, observability.
	return nil
}

func main() {}
