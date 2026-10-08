# Lesson 073 — Webhooks, Async Jobs, and Channels Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.

2. Implement a webhook receiver (a separate small server) that verifies the incoming signature against its own copy of the shared secret before processing the payload, rejecting (with an appropriate 4xx status) any request with a missing or incorrect signature.

3. Test both a correctly signed request (accepted) and a tampered payload with a stale signature (rejected) to confirm signature verification is actually functioning, not just present in code but unused.

## Run and check

```powershell
Set-Location lessons/S2/073-webhooks-async-jobs-and-channels/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Async webhook pipeline

Build a webhook receiver that validates input, persists an event, and dispatches background work. Handle duplicates and worker shutdown.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
