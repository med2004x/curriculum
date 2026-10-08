# Code map — Lesson 047 — Packet Analysis and Observability

This directory is the coding workspace for this lesson. It is deliberately not a solution bank.

| File | Role | Use |
|---|---|---|
| `exercise.go` | EXERCISE | Student starter. Incomplete by design. Implement it from WORK.md. |

## File rules

`DEMO` means complete teaching code. `PATTERN` means incomplete reference code. `EXERCISE` and `CHALLENGE` are your work. `TEST` checks your implementation. `SUPPORT` and `SETUP` files exist only to make the exercise runnable.

There are no hidden solution files in this lesson's code directory.

## Challenge from WORK.md

### Packet reasoning lab

Capture or inspect a small HTTP/TCP exchange and annotate the packet sequence, connection state, timing, and observable failure point.

### Acceptance criteria

- The main path works on normal input.
- At least two boundary or invalid cases are tested.
- One failure mode is deliberately reproduced and then fixed.
- The implementation can be explained from a blank page in terms of state and invariants.
- Verification is reproducible from commands recorded in your evidence file.

## 8. Debugging and failure lab

1. Attempt to diagnose a "slow request" using *only* application-level logging (e.g. a log line printed when your Lesson 4 server starts and finishes handling a request, with no packet capture) for the same lossy-network scenario from step 3.

2. Confirm that your application-level logs show the request as slow, but provide no way to distinguish "the network was lossy and retransmission ate the time" from "my server code itself was slow to process the request" — both would produce an identical symptom (a large gap between request-received and response-sent) if the retransmission delay is happening *b.

3. Explain why this specific ambiguity is exactly what motivates packet-level debugging as a distinct skill from application-log analysis: the two failure modes are indistinguishable from inside the application, but trivially distinguishable from a packet capture that shows the network layer independently of what your code was doing.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

Rule: never fix a failure by changing multiple unrelated things at once. The objective is diagnosis, not ritual.

## 9. Transfer and design reasoning

Change exactly one const

## First run

```sh
go run exercise.go
```

For multi-file or database work, follow the exact commands in the lesson and record them in `EVIDENCE.md`.
