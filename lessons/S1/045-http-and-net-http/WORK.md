# Lesson 045 — HTTP and net/http Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Using raw TCP sockets (Lesson 1, not an HTTP library), implement a minimal HTTP server that parses an incoming request line and headers (up to the blank-line terminator), and responds with a hardcoded `200 OK` response and a small body, including a correct `Content-Length` header matching the actual body byte length.

2. Test your server with a real client (`curl -v` shows the raw request/response headers) and confirm the response is correctly interpreted — `curl` should display your status code and body without error, and `curl -v`'s verbose output should show the exact bytes exchanged, letting you confirm your hand-built response matches the format real clients expect.

3. Extend your server to support keep-alive: after responding, don't close the connection — loop back and parse another request on the same TCP connection, unless the client's request included `Connection: close`.

## Run and check

```powershell
Set-Location lessons/S1/045-http-and-net-http/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### HTTP service

Build a small `net/http` service with two resources, clear status codes, request logging, and tests that exercise malformed and valid requests.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
