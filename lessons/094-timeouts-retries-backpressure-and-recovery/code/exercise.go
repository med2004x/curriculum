package main

import (
	"context"
	"net/http"
	"time"
)

func request(ctx context.Context, client *http.Client, url string) error {
	// TODO: apply bounded retries, per-attempt deadlines, and backoff without retrying unsafe operations.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(req)
	if err == nil {
		return nil
	}
	time.Sleep(10 * time.Millisecond)
	return err
}

func main() {}
