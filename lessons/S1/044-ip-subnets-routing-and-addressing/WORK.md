# Lesson 044 — IP Addressing, Subnets, and Routing Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Given a small network diagram, calculate subnets and determine which gateway handles several destinations.

2. Inspect a real routing table on Linux.

3. Write a second case with one changed constraint and verify it.

## Run and check

```powershell
Set-Location lessons/S1/044-ip-subnets-routing-and-addressing/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Subnet calculator

Build a subnet calculator that accepts an IPv4 address and prefix, then reports network, broadcast, host range, and capacity with invalid-input checks.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
