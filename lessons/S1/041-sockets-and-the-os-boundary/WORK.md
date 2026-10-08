# Lesson 041 — Sockets and the OS Boundary Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Write a minimal TCP server (bind, listen, accept) and client (connect) in Go or C that exchanges a single short message.

2. Confirm it works for a small message.

3. Demonstrate the "no message boundaries" fact directly: modify the client to send two separate messages via two separate `write()`/`Write()` calls in quick succession (e.g. "hello" then "world", with no delay between them), and have the server do a single `read()` immediately after accepting the connection, printing exactly what it received in that one read.

## Run and check

```powershell
Set-Location lessons/S1/041-sockets-and-the-os-boundary/code
go test -race exercise.go exercise_test.go
```

## Optional challenge

### Socket protocol

Implement a small client/server protocol over TCP with explicit framing and error responses. Test partial reads and client disconnects.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
