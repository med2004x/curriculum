# Lesson 086 — Distributed Failures and System Models Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself.

2. For each point, state whether the client, from a plain timeout alone, can distinguish it from any of the other points.

3. Using `tc netem` (networking Lesson 7's tool) or a similar mechanism, simulate a network partition between a client and server (block all traffic between them for a period, e.g. via a firewall rule or `tc` loss configuration set to 100%), and observe what your client-side code actually experiences — a timeout, a connection refused, or a hang, depending on yo.

## Run and check

```powershell
Set-Location lessons/S2/086-distributed-failures-and-system-models/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Failure model

Design a distributed protocol under a stated failure model. Enumerate what can be lost, delayed, duplicated, or reordered and which guarantees are still achievable.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
