// STUDENT PROBLEM SET — Lesson 068.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Define one create-task use case and a narrow repository interface.
// 2. Implement service validation and a fake repository for isolated tests.
// 3. Add an HTTP handler that maps typed errors to documented responses; add a separate database-backed repository test. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 068: Handlers, Services, Repositories, and Boundaries
// Your task: keep business rules in the service boundary.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
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
    // TODO: implement the assigned behavior.
}

func main() {}

// Keep starter imports available while the student implementation is unfinished.
var _ context.Context
