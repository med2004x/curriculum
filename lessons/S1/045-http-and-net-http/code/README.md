# Code map — Lesson 045 — HTTP and net/http

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `demo_01.go` | DEMO | Complete reference example. Read and run it. Do not treat it as your submission. |
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |
| `exercise_test.go` | TEST | Automated check/scaffold. Use it to verify your own implementation. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### HTTP service

Build a small `net/http` service with two resources, clear status codes, request logging, and tests that exercise malformed and valid requests.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Deliberately send a `Content-Length` header from your server that doesn't match the actual body's real byte length (e.g. claim 100 bytes but send only 50).

2. Test with `curl -v` and observe the client either hang waiting for the remaining, never-arriving bytes (if it's a keep-alive connection, since the client believes more of the response is still coming), or produce a truncated/corrupted body.

3. Explain, using Lesson 1's framing discussion directly, why this specific bug is a framing-boundary violation — the client is doing exactly what correct framing logic should do (trust the length prefix and read that many bytes), and the bug is entirely in your server providing an inaccurate length, not in the client's read logic.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one constraint in the main case for # 045 — HTTP and net/http: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new

## First run

```sh
go run exercise.go
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
