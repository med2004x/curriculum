# Lesson 078 — Cryptography and Secure Protocols Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Compute SHA-256 hashes of several inputs (using your language's standard crypto library, never a hand-rolled implementation) and confirm the hash changes completely for even a single-character input change (the "avalanche effect" — a defining property of a good cryptographic hash), demonstrating this isn't a simple checksum where small input changes produce.

2. Implement HMAC-based message signing and verification (extending API engineering Lesson 11's practical webhook use directly), and confirm that altering either the message or the signature (but not both consistently) causes verification to correctly fail — test both cases explicitly, not just the happy path.

3. Generate a signing key pair using an established library, sign a message, verify with the correct public key, then verify failure with an unrelated public key.

## Run and check

```powershell
Set-Location lessons/S2/078-cryptography-and-secure-protocols/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Secure protocol

Implement a small authenticated message exchange using a standard cryptographic primitive from the standard library or a vetted package. Explain the threat it solves and the one it does not.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
