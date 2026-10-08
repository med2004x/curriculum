# Lesson 086 — Problem Set

**Topic:** Distributed Failures and System Models.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself.

2. Extend it with one explicit boundary case: For each point, state whether the client, from a plain timeout alone, can distinguish it from any of the other points.

3. Rebuild the same behavior with one changed constraint: Using `tc netem` (networking Lesson 7's tool) or a similar mechanism, simulate a network partition between a client and server (block all traffic between them for a period, e.g. via a firewall rule or `tc` loss configuration set to 100%), and observe what your client-side code actually experiences — a timeout, a connection refused, or a hang, depending on yo.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/086-distributed-failures-and-system-models/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Failure model

Design a distributed protocol under a stated failure model. Enumerate what can be lost, delayed, duplicated, or reordered and which guarantees are still achievable.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
