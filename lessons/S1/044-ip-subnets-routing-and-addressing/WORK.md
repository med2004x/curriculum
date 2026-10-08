# Lesson 044 — Problem Set

**Topic:** IP Addressing, Subnets, and Routing.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Given a small network diagram, calculate subnets and determine which gateway handles several destinations.

2. Extend it with one explicit boundary case: Inspect a real routing table on Linux.

3. Rebuild the same behavior with one changed constraint: Write a second case with one changed constraint and verify it.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/044-ip-subnets-routing-and-addressing/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Subnet calculator

Build a subnet calculator that accepts an IPv4 address and prefix, then reports network, broadcast, host range, and capacity with invalid-input checks.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
