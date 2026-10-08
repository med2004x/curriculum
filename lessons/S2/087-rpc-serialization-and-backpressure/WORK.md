# Lesson 087 — RPC, Serialization, and Backpressure Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Build a minimal RPC client/server (Go's built-in `net/rpc` package is a reasonable, quick way to get a working RPC mechanism without hand-rolling wire-format parsing, though you're welcome to build directly on networking Lesson 1's raw sockets/framing if you want the deeper exercise) with at least one method that takes an argument and returns a result.

2. Add a client-side timeout: if the server doesn't respond within a configured duration, the client call returns an error rather than blocking indefinitely.

3. Test it by adding an artificial delay to the server's handler (longer than your configured timeout) and confirming the client correctly times out rather than hanging.

## Run and check

```powershell
Set-Location lessons/S2/087-rpc-serialization-and-backpressure/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### RPC pipeline

Build an RPC-like request/response layer with serialization, bounded queues, deadlines, and explicit overload behavior.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
