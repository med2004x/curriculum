# Lesson 047 — Packet Analysis and Observability

## Threads
**Primary:** Networking/Linux

**Supporting:** pcap thinking; tracing; latency evidence

**Math companion:** statistics and distributions

**Time:** 3–4 hours. **Prerequisites:** Lesson 046.

**Time:** 3–4 hours. **Prerequisites:** Lesson 046, # Lesson 046 — TLS, Cryptography, and Trust.

## What you will be able to do

1. Explain the central model of # 047 — Packet Analysis and Observability and name the boundary where it applies.
2. Trace one concrete case of # 047 — Packet Analysis and Observability before executing it.
3. Rebuild the core exercise from an empty workspace without copying the reference implementation.
4. Produce terminal evidence for one success case and one failure case.

## Retrieve first (no notes)

1. State the main concept you learned in Lesson 046, # Lesson 046 — TLS, Cryptography, and Trust.
2. Give one boundary case from the previous lesson.
3. Name one command, formula, invariant, or debugging technique used in the previous lesson.
4. Predict what would change if one input from the previous lesson changed.

## 1. Problem and model

**What problem this solves.** Use packet captures as direct evidence when diagnosing network problems, rather than reasoning purely from application-level symptoms — and specifically learn to recognize retransmissions and handshake anomalies in raw capture output.

**The contract.** **Why packet captures are the ground truth for network problems.** Application-level logs tell you what your code observed (a timeout, an error, a slow response) but not *why* at the network level — was it DNS resolution taking too long (Lesson 3), a slow TCP handshake (Lesson 2), packet loss and retransmission, or the application server itself being slow to.

**Worked example.** Start with the smallest concrete case described by the source material: Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets b.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** A packet capture shows exactly what went over the wire, with precise timestamps, letting you attribute delay to a specific phase rather than guessing from an aggregate "it was slow" symptom.

**Second point.** **Reading a capture: the essentials.** Tools like `tcpdump` (command-line) and Wireshark (GUI, generally easier for visual pattern recognition, though `tcpdump` is more often available on remote servers you're debugging) show each packet's source/destination, protocol, and (for TCP) flags (`SYN`, `ACK`, `FIN`, `RST`) and sequence/acknowledgment numbers.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** In the same capture, identify the first HTTP request packet (containing the actual `GET`/`POST` line) and the first response packet (containing `HTTP/1.1 ...`), and compute the time delta between them — this is your directly measured "server think time" for that specific request.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets b.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** For step 2, report the actual measured server-think-time in milliseconds for a real request, and for step 4, report the actual retransmission-attributable delay in milliseconds, both as concrete numbers extracted from real capture timestamps — not estimates.

**Failure mode.** Attempt to diagnose a "slow request" using *only* application-level logging (e.g. a log line printed when your Lesson 4 server starts and finishes handling a request, with no packet capture) for the same lossy-network scenario from step 3.

**Completion criterion.** You've captured and correctly identified a real three-way handshake and the server-think-time delta for an actual request, you've deliberately induced and observed real TCP retransmissions using artificial packet loss, and you can explain — using the failure drill's specific ambiguity — why application-level logging alone cannot distinguish network-layer del.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the result of the smallest # 047 — Packet Analysis and Observability case.
2. Predict one boundary case and state exactly which observation would confirm it.
3. Predict the first failure signal if one assumption is broken.
4. Predict which state, file, process, query, or interface is responsible for the observable result.
5. Predict what changes when one constraint is changed and all other inputs stay fixed.
6. Predict the verification command before you run it.
7. Predict one debugging step that would reduce the problem to a smaller case.
8. Write the prediction down before executing any of the cases.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets b.

2. Extend it with one explicit boundary case: In the same capture, identify the first HTTP request packet (containing the actual `GET`/`POST` line) and the first response packet (containing `HTTP/1.1 ...`), and compute the time delta between them — this is your directly measured "server think time" for that specific request.

3. Rebuild the same behavior with one changed constraint: Deliberately introduce artificial packet loss to observe retransmission directly: using a tool like `tc` (Linux traffic control, e.g.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Attempt to diagnose a "slow request" using *only* application-level logging (e.g. a log line printed when your Lesson 4 server starts and finishes handling a request, with no packet capture) for the same lossy-network scenario from step 3.

2. Confirm that your application-level logs show the request as slow, but provide no way to distinguish "the network was lossy and retransmission ate the time" from "my server code itself was slow to process the request" — both would produce an identical symptom (a large gap between request-received and response-sent) if the retransmission delay is happening *b.

3. Explain why this specific ambiguity is exactly what motivates packet-level debugging as a distinct skill from application-log analysis: the two failure modes are indistinguishable from inside the application, but trivially distinguishable from a packet capture that shows the network layer independently of what your code was doing.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 047 — Packet Analysis and Observability: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What is the main objective stated for # 047 — Packet Analysis and Observability?
2. What two concrete practices are required in the assigned material for # 047 — Packet Analysis and Observability?
3. What does the lesson ask you to reproduce before making changes?
4. What failure mode is explicitly used as a diagnostic exercise?
5. What evidence does the verification step require you to record?
6. What condition in the "Done when" criterion proves mastery of this lesson?

## Evidence to record

Create `lesson047/EVIDENCE.md` with exactly this structure:

```text
# Lesson 047 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Copying the reference before forming a prediction for # 047 — Packet Analysis and Observability.
- Changing several variables at once, which destroys the evidence trail.
- Recording a final result without recording the command or input that produced it.
- Treating a passing case as proof that boundary cases are correct.
- Using terminology loosely when the source material defines a precise concept.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 047 — Packet Analysis and Observability without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've captured and correctly identified a real three-way handshake and the server-think-time delta for an actual request, you've deliberately induced and observed real TCP retransmissions using artificial packet loss, and you can explain — using the failure drill's specific ambiguity — why application-level logging alone cannot distinguish network-layer del. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
