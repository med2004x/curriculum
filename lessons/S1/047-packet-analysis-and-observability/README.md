# Lesson 047 — Packet Analysis and Observability

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Use packet captures as direct evidence when diagnosing network problems, rather than reasoning purely from application-level symptoms — and specifically learn to recognize retransmissions and handshake anomalies in raw capture output.

**The contract.** **Why packet captures are the ground truth for network problems.** Application-level logs tell you what your code observed (a timeout, an error, a slow response) but not *why* at the network level — was it DNS resolution taking too long (Lesson 3), a slow TCP handshake (Lesson 2), packet loss and retransmission, or the application server itself being slow to.

**Worked example.** Start with the smallest concrete case described by the source material: Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets b.

## 2. Mechanism and vocabulary

**Mechanism.** A packet capture shows exactly what went over the wire, with precise timestamps, letting you attribute delay to a specific phase rather than guessing from an aggregate "it was slow" symptom.

**Second point.** **Reading a capture: the essentials.** Tools like `tcpdump` (command-line) and Wireshark (GUI, generally easier for visual pattern recognition, though `tcpdump` is more often available on remote servers you're debugging) show each packet's source/destination, protocol, and (for TCP) flags (`SYN`, `ACK`, `FIN`, `RST`) and sequence/acknowledgment numbers.

**Worked example.** In the same capture, identify the first HTTP request packet (containing the actual `GET`/`POST` line) and the first response packet (containing `HTTP/1.1 ...`), and compute the time delta between them — this is your directly measured "server think time" for that specific request.

## 3. Worked case from the reference material

Worked example: Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets b.

## 4. Verification, failure, and evidence

**Verification.** For step 2, report the actual measured server-think-time in milliseconds for a real request, and for step 4, report the actual retransmission-attributable delay in milliseconds, both as concrete numbers extracted from real capture timestamps — not estimates.

**Failure mode.** Attempt to diagnose a "slow request" using *only* application-level logging (e.g. a log line printed when your Lesson 4 server starts and finishes handling a request, with no packet capture) for the same lossy-network scenario from step 3.

**Completion criterion.** You've captured and correctly identified a real three-way handshake and the server-think-time delta for an actual request, you've deliberately induced and observed real TCP retransmissions using artificial packet loss, and you can explain — using the failure drill's specific ambiguity — why application-level logging alone cannot distinguish network-layer del.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
