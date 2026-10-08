# Lesson 045 — Problem Set

**Topic:** HTTP and net/http.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Using raw TCP sockets (Lesson 1, not an HTTP library), implement a minimal HTTP server that parses an incoming request line and headers (up to the blank-line terminator), and responds with a hardcoded `200 OK` response and a small body, including a correct `Content-Length` header matching the actual body byte length.

2. Extend it with one explicit boundary case: Test your server with a real client (`curl -v` shows the raw request/response headers) and confirm the response is correctly interpreted — `curl` should display your status code and body without error, and `curl -v`'s verbose output should show the exact bytes exchanged, letting you confirm your hand-built response matches the format real clients expect.

3. Rebuild the same behavior with one changed constraint: Extend your server to support keep-alive: after responding, don't close the connection — loop back and parse another request on the same TCP connection, unless the client's request included `Connection: close`.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/045-http-and-net-http/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### HTTP service

Build a small `net/http` service with two resources, clear status codes, request logging, and tests that exercise malformed and valid requests.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
