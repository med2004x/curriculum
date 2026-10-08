# Lesson 046 — TLS, Cryptography, and Trust Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Generate a local, self-signed TLS certificate (e.g.

2. `openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 1 -nodes`) and use it to run a minimal HTTPS server (Go's `net/http` with `ListenAndServeTLS`, pointing at your generated cert/key, is straightforward for this).

3. Connect to it with `curl -k` (the `-k` flag skips certificate verification, appropriate here since you know it's self-signed and you're testing locally) and confirm the connection succeeds and is encrypted.

## Run and check

```powershell
Set-Location lessons/S1/046-tls-crypto-and-trust/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### TLS client

Write a TLS client diagnostic that verifies certificate properties, records the negotiated protocol, and explains which trust decisions are local versus remote.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
