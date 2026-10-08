package main

import "context"

type Product struct {
	ID   int
	Name string
}
type Repository interface {
	Get(context.Context, int) (Product, error)
}
type Service struct{ repo Repository }

func (s Service) Get(ctx context.Context, id int) (Product, error) {
	// TODO: keep business rules in the service boundary.
	return s.repo.Get(ctx, id)
}

func main() {}
