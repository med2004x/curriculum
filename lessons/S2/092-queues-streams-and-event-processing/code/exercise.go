package main

type Message struct {
	Topic string
	Value string
}

func publish(topics map[string][]Message, msg Message) {
	// TODO: append without violating ordering or ownership assumptions.
	topics[msg.Topic] = append(topics[msg.Topic], msg)
}

func main() {}
