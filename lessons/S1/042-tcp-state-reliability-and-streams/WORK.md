# Lesson 042 — Problem Set

**Topic:** TCP State, Reliability, and Streams.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Annotate a TCP handshake trace with both initial sequence numbers and acknowledgment values.

2. Annotate a data trace with byte ranges, cumulative acknowledgments, and any retransmission.

3. Compare a receiver-window change with a congestion-window change and explain which endpoint or network condition each represents.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/042-tcp-state-reliability-and-streams/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Reliable stream transfer

Build a framed TCP transfer of length-prefixed messages and verify behavior under chunked writes and delayed reads.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
