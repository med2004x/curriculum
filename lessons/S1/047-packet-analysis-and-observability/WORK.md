# Lesson 047 — Problem Set

**Topic:** Packet Analysis and Observability.
**Your files:** `code/exercise.go`. `demo_*` and `reference/` are teaching examples; implement the assigned task files.

## Problems

1. Build the smallest useful version: Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets b.

2. Extend it with one explicit boundary case: In the same capture, identify the first HTTP request packet (containing the actual `GET`/`POST` line) and the first response packet (containing `HTTP/1.1 ...`), and compute the time delta between them — this is your directly measured "server think time" for that specific request.

3. Rebuild the same behavior with one changed constraint: Deliberately introduce artificial packet loss to observe retransmission directly: using a tool like `tc` (Linux traffic control, e.g.

Do not open the reference implementation until your own attempt has been saved.

Start with a prediction. Implement the TODOs without opening completed examples. Check one normal input, one boundary, and one failure case. The specifications and TODO prompts are repeated in comments in the student files.

## Run and check

```powershell
Set-Location lessons/S1/047-packet-analysis-and-observability/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Stretch problem

### Packet reasoning lab

Capture or inspect a small HTTP/TCP exchange and annotate the packet sequence, connection state, timing, and observable failure point.

## Turn in

In `EVIDENCE.md`, record your prediction, exact command, observed output, one failed attempt and its repair, and the invariant or reasoning behind the result.
