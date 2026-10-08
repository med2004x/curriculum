# Lesson 045 — HTTP and net/http

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build a minimal HTTP server and client from raw sockets (not a framework) to understand exactly what HTTP adds on top of the TCP byte stream from Lessons 1-2, and correctly implement connection reuse (keep-alive) and status-code semantics.

**The contract.** An HTTP/1.1 message begins with a start line, followed by CRLF-delimited headers and a blank line, then an optional body. Message-body length is determined by HTTP framing rules such as `Content-Length` or `Transfer-Encoding: chunked`; never assume the body always has a content-length header. Header names are case-insensitive, and content length counts octets, not Unicode characters.

**Worked example.** Use `net/http` to serve `/health` with a small body and return `404 Not Found` for an unknown path. Then inspect a request/response with a client and identify the start line, headers, blank-line boundary, status, and body framing.

## 2. Mechanism and vocabulary

**Mechanism.** In Go, an `http.Handler` receives a parsed request and writes a response through `http.ResponseWriter`; the server owns parsing, connection management, and protocol framing. Handler logic should validate input and set status before writing body bytes because headers are committed when output begins.

**Second point.** HTTP/1.1 connections are persistent by default unless the protocol or peer closes them. A response can use `Content-Length`, chunked transfer coding, or connection-close framing where applicable; `Connection: close` requests closure rather than keep-alive behavior.

Do not hand-roll a general HTTP parser for production. The raw-TCP exercise is only to understand framing; use `net/http` for real server behavior and its request parsing, body limits, and connection lifecycle.

**Worked example.** Test your server with a real client (`curl -v` shows the raw request/response headers) and confirm the response is correctly interpreted — `curl` should display your status code and body without error, and `curl -v`'s verbose output should show the exact bytes exchanged, letting you confirm your hand-built response matches the format real clients expect.

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

## 4. Verification, failure, and evidence

**Verification.** Send a valid request and an unknown-path request; record status, headers, and body. Send multiple requests through a client that reuses connections and distinguish HTTP persistence from the fact that TCP itself has no request boundaries.

**Failure mode.** Deliberately mismatch a response `Content-Length` and body byte count in an isolated local experiment. Explain why a client can wait for missing bytes or treat following bytes as part of the body; restore correct framing afterward.

**Completion criterion.** You can explain HTTP message framing, implement route-specific outcomes using `net/http`, distinguish persistent HTTP connections from TCP message boundaries, and explain the effect of an incorrect body length.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
