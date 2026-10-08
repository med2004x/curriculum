# Lesson 041 — Problem Set

**Topic:** Sockets and the OS Boundary.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Write a minimal TCP server (bind, listen, accept) and client (connect) in Go or C that exchanges a single short message.

2. Extend it with one explicit boundary case: Confirm it works for a small message.

3. Rebuild the same behavior with one changed constraint: Demonstrate the "no message boundaries" fact directly: modify the client to send two separate messages via two separate `write()`/`Write()` calls in quick succession (e.g. "hello" then "world", with no delay between them), and have the server do a single `read()` immediately after accepting the connection, printing exactly what it received in that one read.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/041-sockets-and-the-os-boundary/code
go test -race exercise.go exercise_test.go
```

## Stretch problem

### Socket protocol

Implement a small client/server protocol over TCP with explicit framing and error responses. Test partial reads and client disconnects.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
