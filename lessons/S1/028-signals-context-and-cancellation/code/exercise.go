package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) error {
	// TODO: stop promptly when ctx is cancelled.
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	err := work(context.Background())
	fmt.Println(err)
}
