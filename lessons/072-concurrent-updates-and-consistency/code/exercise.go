package main

import "errors"

type Account struct {
	Balance int
	Version int
}

func update(a *Account, expectedVersion, delta int) error {
	// TODO: reject stale versions and preserve the account invariant.
	if a.Version != expectedVersion {
		return errors.New("stale version")
	}
	a.Balance += delta
	a.Version++
	return nil
}

func main() {}
