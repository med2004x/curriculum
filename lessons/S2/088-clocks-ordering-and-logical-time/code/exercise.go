package main

type Clock struct {
	ID    int
	Value int
}

func (c *Clock) Tick() { c.Value++ }
func (c *Clock) Receive(remote int) {
	// TODO: implement Lamport's receive rule.
	if remote >= c.Value {
		c.Value = remote + 1
	} else {
		c.Value++
	}
}

func main() {}
