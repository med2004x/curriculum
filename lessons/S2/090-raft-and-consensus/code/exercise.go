package main

type Role int

const (
	Follower Role = iota
	Candidate
	Leader
)

type Node struct {
	Role  Role
	Term  int
	Votes int
}

func becomeCandidate(n *Node) {
	// TODO: increment term, become candidate, and vote for self.
	n.Role = Candidate
	n.Term++
	n.Votes = 1
}

func main() {}
