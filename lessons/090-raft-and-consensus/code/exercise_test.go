package main
import "testing"
func TestBecomeCandidate(t *testing.T) {
 n:=Node{Role:Follower,Term:4,Votes:0}; becomeCandidate(&n)
 if n.Role!=Candidate || n.Term!=5 || n.Votes!=1 {t.Fatalf("state after election start: %+v",n)}
}
func TestBecomeCandidateFromExistingTerm(t *testing.T) {
 n:=Node{Role:Candidate,Term:8,Votes:2}; becomeCandidate(&n)
 if n.Role!=Candidate || n.Term!=9 || n.Votes!=1 {t.Fatalf("new election must advance term and reset vote count: %+v",n)}
}
