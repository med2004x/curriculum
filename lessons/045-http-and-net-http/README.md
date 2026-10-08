# Lesson 045 — HTTP and net/http

## Threads
**Primary:** Networking/Go

**Supporting:** HTTP semantics; request/response; handlers

**Math companion:** functions and state machines

**Time:** 3–4 hours. **Prerequisites:** Lesson 044.

**Time:** 3–4 hours. **Prerequisites:** Lesson 044.

## What you will be able to do

1. Explain the HTTP request/response structure and the role of method, target, headers, status, and body.
2. Distinguish body framing with `Content-Length` from HTTP/1.1 chunked transfer coding.
3. Explain persistent HTTP/1.1 connections and when a connection closes.
4. Implement a handler with Go's `net/http` and choose status codes from the request outcome.

## Retrieve first (no notes)

1. What does a CIDR prefix specify?
2. How does longest-prefix match choose a route?
3. What is the difference between an IP next hop and a link-layer destination?
4. What does ARP resolve on an IPv4 Ethernet link?

## 1. Problem and model

**What problem this solves.** Build a minimal HTTP server and client from raw sockets (not a framework) to understand exactly what HTTP adds on top of the TCP byte stream from Lessons 1-2, and correctly implement connection reuse (keep-alive) and status-code semantics.

**The contract.** An HTTP/1.1 message begins with a start line, followed by CRLF-delimited headers and a blank line, then an optional body. Message-body length is determined by HTTP framing rules such as `Content-Length` or `Transfer-Encoding: chunked`; never assume the body always has a content-length header. Header names are case-insensitive, and content length counts octets, not Unicode characters.

**Worked example.** Use `net/http` to serve `/health` with a small body and return `404 Not Found` for an unknown path. Then inspect a request/response with a client and identify the start line, headers, blank-line boundary, status, and body framing.

Self-check: compute a response body's `Content-Length` in bytes if setting it explicitly; UTF-8 character count can differ from byte count.

## 2. Mechanism and vocabulary

**Mechanism.** In Go, an `http.Handler` receives a parsed request and writes a response through `http.ResponseWriter`; the server owns parsing, connection management, and protocol framing. Handler logic should validate input and set status before writing body bytes because headers are committed when output begins.

**Second point.** HTTP/1.1 connections are persistent by default unless the protocol or peer closes them. A response can use `Content-Length`, chunked transfer coding, or connection-close framing where applicable; `Connection: close` requests closure rather than keep-alive behavior.

Do not hand-roll a general HTTP parser for production. The raw-TCP exercise is only to understand framing; use `net/http` for real server behavior and its request parsing, body limits, and connection lifecycle.

**Worked example.** Test your server with a real client (`curl -v` shows the raw request/response headers) and confirm the response is correctly interpreted — `curl` should display your status code and body without error, and `curl -v`'s verbose output should show the exact bytes exchanged, letting you confirm your hand-built response matches the format real clients expect.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

The assigned material contains a concrete implementation or command sequence. Start from this excerpt, then run it in your own workspace:

```go
package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, "ok\n")
}
```

Source: supporting material is integrated into this lesson; see `../../SOURCES.md`.

This handler delegates parsing and framing to `net/http`; check how an unknown route differs from an internal server error and why the response body must not be written before choosing the status.

Self-check: what status should be returned for a syntactically valid request to an unknown path?

## 4. Verification, failure, and evidence

**Verification.** Send a valid request and an unknown-path request; record status, headers, and body. Send multiple requests through a client that reuses connections and distinguish HTTP persistence from the fact that TCP itself has no request boundaries.

**Failure mode.** Deliberately mismatch a response `Content-Length` and body byte count in an isolated local experiment. Explain why a client can wait for missing bytes or treat following bytes as part of the body; restore correct framing afterward.

**Completion criterion.** You can explain HTTP message framing, implement route-specific outcomes using `net/http`, distinguish persistent HTTP connections from TCP message boundaries, and explain the effect of an incorrect body length.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Identify the request line, headers, header terminator, and body in an HTTP/1.1 request.
2. Predict the status for `/health` and an unknown path.
3. Predict how `Content-Length` differs from chunked transfer coding.
4. Predict whether HTTP/1.1 normally permits connection reuse.
5. Predict what occurs if a declared body length exceeds the bytes sent.
6. Predict why a UTF-8 body can have a byte length different from its character count.
7. State which parsing and connection behavior `net/http` manages.
8. Explain why one TCP read is not one HTTP request.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Using raw TCP sockets (Lesson 1, not an HTTP library), implement a minimal HTTP server that parses an incoming request line and headers (up to the blank-line terminator), and responds with a hardcoded `200 OK` response and a small body, including a correct `Content-Length` header matching the actual body byte length.

2. Extend it with one explicit boundary case: Test your server with a real client (`curl -v` shows the raw request/response headers) and confirm the response is correctly interpreted — `curl` should display your status code and body without error, and `curl -v`'s verbose output should show the exact bytes exchanged, letting you confirm your hand-built response matches the format real clients expect.

3. Rebuild the same behavior with one changed constraint: Extend your server to support keep-alive: after responding, don't close the connection — loop back and parse another request on the same TCP connection, unless the client's request included `Connection: close`.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Deliberately send a `Content-Length` header from your server that doesn't match the actual body's real byte length (e.g. claim 100 bytes but send only 50).

2. Test with `curl -v` and observe the client either hang waiting for the remaining, never-arriving bytes (if it's a keep-alive connection, since the client believes more of the response is still coming), or produce a truncated/corrupted body.

3. Explain, using Lesson 1's framing discussion directly, why this specific bug is a framing-boundary violation — the client is doing exactly what correct framing logic should do (trust the length prefix and read that many bytes), and the bug is entirely in your server providing an inaccurate length, not in the client's read logic.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 045 — HTTP and net/http: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What components make up an HTTP/1.1 request message?
2. What does the blank line after headers mark?
3. How can HTTP/1.1 frame a message body?
4. What does `net/http` provide beyond a handler function?
5. How does a persistent HTTP connection differ from a TCP message boundary?
6. Why is content length measured in bytes?

## Evidence to record

Create `lesson045/EVIDENCE.md` with exactly this structure:

```text
# Lesson 045 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Assuming every HTTP message has a `Content-Length` header.
- Counting Unicode characters rather than body octets.
- Treating HTTP keep-alive as message framing.
- Writing a status/body before deciding the response status.
- Using a hand-rolled parser for general or untrusted production traffic.

## Precision checks

1. Trace request and response framing through their header/body boundaries.
2. State which body-length rule applies to each example.
3. Explain status selection and unknown-route handling in the Go handler.
4. Explain connection reuse without claiming that it creates message boundaries.
5. Explain why general HTTP parsing belongs to `net/http`.

HTTP adds request semantics and message framing above TCP; Go's standard library should own protocol parsing for application servers.
## University revision bridge

For the LIG track, connect HTTP message grammar and status semantics to the transport byte stream, then implement application behavior through Go's server abstractions.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 045 — HTTP and net/http without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. Your hand-rolled HTTP server correctly handles request parsing, accurate `Content-Length` responses, keep-alive connection reuse (verified at the TCP level, not just assumed), and distinct status codes for different paths, and you've directly demonstrated — via the failure drill — what breaks when the length-prefix framing contract is violated, connecting it. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
