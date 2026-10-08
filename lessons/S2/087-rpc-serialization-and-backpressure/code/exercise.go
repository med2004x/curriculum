package main

type Request struct {
	ID      int
	Payload []byte
}

func send(queue chan<- Request, req Request) bool {
	// TODO: make admission/backpressure behavior explicit and testable.
	select {
	case queue <- req:
		return true
	default:
		return false
	}
}

func main() {}
