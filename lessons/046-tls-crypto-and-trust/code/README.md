# Code map — Lesson 046 — TLS, Cryptography, and Trust

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### TLS client

Write a TLS client diagnostic that verifies certificate properties, records the negotiated protocol, and explains which trust decisions are local versus remote.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Take your step 1 self-signed server and, using a packet capture tool (Lesson 7 preview) or a network proxy tool, attempt to read the actual HTTP request/response content passing over your `curl -k` connection at the raw TCP byte level (not through the TLS-aware tool itself, but genuinely at the wire).

2. Confirm the payload is unreadable ciphertext, not plaintext HTTP — directly demonstrating the confidentiality guarantee is real and independent of whether authentication was properly verified in this specific test (since you used `-k` to skip that check).

3. Explain why this is the concrete, hands-on version of the abstract "separable guarantees" point from Learn: your step 1/2 test showed authentication can fail while encryption still works, and this drill shows encryption genuinely works (not just claimed) even in a scenario where you deliberately bypassed authentication.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint in the main case for # 046 — TLS

## First run

```sh
go run exercise.go
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
