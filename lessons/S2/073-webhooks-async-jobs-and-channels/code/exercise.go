package main

import "context"

func worker(ctx context.Context, jobs <-chan string, done chan<- string) {
	// TODO: consume jobs, handle cancellation, and report completion exactly once.
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			done <- job
		}
	}
}

func main() {}
