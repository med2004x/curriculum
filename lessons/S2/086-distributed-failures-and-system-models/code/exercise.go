package main

type Outcome int

const (
	Success Outcome = iota
	RetryableFailure
	PermanentFailure
)

func classify(err error) Outcome {
	// TODO: classify failures using the system model from the lesson.
	if err == nil {
		return Success
	}
	return RetryableFailure
}

func main() {}
