// STUDENT TASK: implement the election-start transition described in WORK.md.
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
	// TODO: begin a new term and cast this node's first vote for itself.
}

func main() {}
