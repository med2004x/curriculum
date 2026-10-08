# Lesson 090 — Raft and Consensus Problems

Read [README.md](README.md). First complete the small local exercise; then use MIT's lab for the full Raft implementation.

## Problems

1. In `code/exercise.go`, implement `becomeCandidate`: increment the term, set the role to `Candidate`, and set the self-vote count to 1.
2. Run the supplied tests. They check starting an election from a follower and restarting an election from an existing candidate.
3. In the MIT lab, implement leader election and use its tests for partitions, delays, and leader failure.

## Run and check

From this lesson's `code` folder, run `go test -race exercise.go exercise_test.go`.

## Course lab

Use the [MIT 6.5840 lab](https://pdos.csail.mit.edu/6.5840/) and the source/tests from the same release. In the supplied 2026 checkout, run `make -C src raft1`. Record the release, command, and test results in [EVIDENCE.md](EVIDENCE.md). The local Go exercise teaches one state transition; the MIT lab is the full consensus assignment.

## Submit

Record the test results and explain why a node increments its term and votes for itself when an election starts.
