# Lesson 042 — TCP State, Reliability, and Streams Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Annotate a TCP handshake trace with both initial sequence numbers and acknowledgment values.

2. Annotate a data trace with byte ranges, cumulative acknowledgments, and any retransmission.

3. Compare a receiver-window change with a congestion-window change and explain which endpoint or network condition each represents.

## Run and check

```powershell
Set-Location lessons/S1/042-tcp-state-reliability-and-streams/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Reliable stream transfer

Build a framed TCP transfer of length-prefixed messages and verify behavior under chunked writes and delayed reads.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
