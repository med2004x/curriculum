# Lesson 047 — Packet Analysis and Observability Problems

Read [README.md](README.md), then solve these problems in order.

**Edit:** `code/exercise.go`.

## Problems

1. Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets b.

2. In the same capture, identify the first HTTP request packet (containing the actual `GET`/`POST` line) and the first response packet (containing `HTTP/1.1 ...`), and compute the time delta between them — this is your directly measured "server think time" for that specific request.

3. Deliberately introduce artificial packet loss to observe retransmission directly: using a tool like `tc` (Linux traffic control, e.g.

## Run and check

```powershell
Set-Location lessons/S1/047-packet-analysis-and-observability/code
go run exercise.go
```

Add `exercise_test.go` with normal, boundary, and failure cases; a successful build alone is not a passing assignment.

## Optional challenge

### Packet reasoning lab

Capture or inspect a small HTTP/TCP exchange and annotate the packet sequence, connection state, timing, and observable failure point.

## Submit

Record the commands and results, one error you fixed, and the reason your solution is correct in [EVIDENCE.md](EVIDENCE.md).
