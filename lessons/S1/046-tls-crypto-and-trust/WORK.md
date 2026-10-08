# Lesson 046 — Problem Set

**Topic:** TLS, Cryptography, and Trust.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Generate a local, self-signed TLS certificate (e.g.

2. Extend it with one explicit boundary case: `openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 1 -nodes`) and use it to run a minimal HTTPS server (Go's `net/http` with `ListenAndServeTLS`, pointing at your generated cert/key, is straightforward for this).

3. Rebuild the same behavior with one changed constraint: Connect to it with `curl -k` (the `-k` flag skips certificate verification, appropriate here since you know it's self-signed and you're testing locally) and confirm the connection succeeds and is encrypted.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/046-tls-crypto-and-trust/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### TLS client

Write a TLS client diagnostic that verifies certificate properties, records the negotiated protocol, and explains which trust decisions are local versus remote.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
