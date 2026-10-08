package main

type Decision struct {
	Authenticated bool
	Authorized    bool
	Allowed       bool
}

func evaluate(authenticated, authorized, trustedInput bool) Decision {
	// TODO: make the trust boundaries explicit.
	return Decision{authenticated, authorized, authenticated && authorized && trustedInput}
}

func main() {}
