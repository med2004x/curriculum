# Lesson 087 — Problem Set

**Topic:** RPC, Serialization, and Backpressure.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Build a minimal RPC client/server (Go's built-in `net/rpc` package is a reasonable, quick way to get a working RPC mechanism without hand-rolling wire-format parsing, though you're welcome to build directly on networking Lesson 1's raw sockets/framing if you want the deeper exercise) with at least one method that takes an argument and returns a result.

2. Extend it with one explicit boundary case: Add a client-side timeout: if the server doesn't respond within a configured duration, the client call returns an error rather than blocking indefinitely.

3. Rebuild the same behavior with one changed constraint: Test it by adding an artificial delay to the server's handler (longer than your configured timeout) and confirming the client correctly times out rather than hanging.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/087-rpc-serialization-and-backpressure/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### RPC pipeline

Build an RPC-like request/response layer with serialization, bounded queues, deadlines, and explicit overload behavior.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
