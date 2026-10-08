// STUDENT PROBLEM SET — Lesson 090.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: MIT's 6.5840 (referenced throughout this track) provides a rigorous, well-tested lab sequence for implementing Raft incrementally, with a real test harness that specifically injects network partitions, delays, and crashes — using it directly is strongly preferred over reconstructing an equivalent test harness from scratch, since correctly testing a consensus.
// 2. Extend it with one explicit boundary case: Implement Raft leader election (the course's first Raft lab): nodes start as followers, time out and become candidates, request votes, and a majority-vote winner becomes leader for that term.
// 3. Rebuild the same behavior with one changed constraint: Test it under the course's provided test harness, including scenarios where the current leader is forcibly disconnected and a new election must correctly occur. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 090: Raft and Consensus
// Your task: increment term, become candidate, and vote for self.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
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
    // TODO: implement the assigned behavior.
}

func main() {}
