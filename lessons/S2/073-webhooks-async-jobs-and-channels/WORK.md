# Lesson 073 — Problem Set

**Topic:** Webhooks, Async Jobs, and Channels.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.

2. Extend it with one explicit boundary case: Implement a webhook receiver (a separate small server) that verifies the incoming signature against its own copy of the shared secret before processing the payload, rejecting (with an appropriate 4xx status) any request with a missing or incorrect signature.

3. Rebuild the same behavior with one changed constraint: Test both a correctly signed request (accepted) and a tampered payload with a stale signature (rejected) to confirm signature verification is actually functioning, not just present in code but unused.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S2/073-webhooks-async-jobs-and-channels/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Async webhook pipeline

Build a webhook receiver that validates input, persists an event, and dispatches background work. Handle duplicates and worker shutdown.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
