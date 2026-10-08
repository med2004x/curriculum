package main

import (
	"context"
	"errors"
	"fmt"
)

type Task struct {
	ID    int64
	Title string
	Done  bool
}

type Repository interface {
	Get(ctx context.Context, id int64) (Task, error)
	Create(ctx context.Context, title string) (Task, error)
}

type MemoryRepository struct {
	next  int64
	tasks map[int64]Task
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{next: 1, tasks: make(map[int64]Task)}
}

func (r *MemoryRepository) Get(ctx context.Context, id int64) (Task, error) {
	if err := ctx.Err(); err != nil {
		return Task{}, err
	}
	task, ok := r.tasks[id]
	if !ok {
		return Task{}, errors.New("task not found")
	}
	return task, nil
}

func (r *MemoryRepository) Create(ctx context.Context, title string) (Task, error) {
	if err := ctx.Err(); err != nil {
		return Task{}, err
	}
	task := Task{ID: r.next, Title: title}
	r.tasks[task.ID] = task
	r.next++
	return task, nil
}

func main() {
	repo := NewMemoryRepository()
	task, err := repo.Create(context.Background(), "read SQL plan")
	if err != nil {
		panic(err)
	}
	loaded, err := repo.Get(context.Background(), task.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", loaded)
}
