# Lesson 042 — TCP State, Reliability, and Streams

## Threads
**Primary:** Networking/Go

**Supporting:** connection state; retransmission; byte streams

**Math companion:** probability of loss

**Time:** 3–4 hours. **Prerequisites:** Lesson 041.

**Time:** 3–4 hours. **Prerequisites:** Lesson 041.

## What you will be able to do

1. Explain TCP's ordered byte-stream abstraction and how sequence numbers represent byte positions.
2. Trace the SYN, SYN-ACK, ACK handshake and the purpose of acknowledgments.
3. Distinguish retransmission, receiver flow control, and network congestion control.
4. Explain graceful FIN-based shutdown and how it differs from an RST.

## Retrieve first (no notes)

1. Why does TCP require application-level framing?
2. What happens if a receiver assumes one read returns a complete message?
3. What does a length-prefix protocol need to validate?
4. What does premature EOF mean for a framed message?

## 1. Problem and model

**What problem this solves.** Understand precisely what TCP guarantees (reliable, ordered, byte-stream delivery) and the mechanisms — the three-way handshake, retransmission, windowing — that provide those guarantees over an unreliable network layer beneath it.

**The contract.** TCP presents a reliable, ordered byte stream while a connection is functioning; it does not preserve application write boundaries. TCP detects loss through acknowledgments and retransmits data, numbers bytes in sequence space, and uses flow control and congestion control for distinct purposes. IP itself does not promise TCP-style reliable ordered delivery.

**Worked example.** Read a TCP exchange in a packet trace: client SYN with initial sequence number `x`; server SYN-ACK acknowledging `x+1` and providing sequence number `y`; client ACK acknowledging `y+1`. SYN consumes sequence space. Capture details can be affected by offload and capture location, so interpret flags and sequence numbers rather than assuming one packet equals one application write.

Self-check: acknowledgments refer to the next byte expected, not to a message boundary.

## 2. Mechanism and vocabulary

**Mechanism.** TCP sequence numbers identify byte positions. Cumulative acknowledgments indicate the next expected sequence number; loss can be inferred from duplicate acknowledgments or timeout and trigger retransmission. TCP reorders received bytes before presenting them to the application.

**Second point.** Flow control uses the receiver-advertised window to limit data in flight based on receiver capacity. Congestion control limits sending based on inferred network capacity. They solve different problems.

Each TCP direction closes independently. FIN indicates no more bytes will be sent in that direction; ACK confirms sequence-space receipt. A common close involves separate FIN/ACK exchanges, while RST aborts a connection. Simultaneous close and implementation details can change the exact trace.

**Worked example.** In a packet trace, identify one data segment, its sequence range, the acknowledgment that covers it, and any retransmission. Then identify how each direction's shutdown is represented and whether the connection ends through FIN-based close or reset.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

For a trace, record sender/receiver, flags, sequence number, acknowledgment number, advertised window, and payload length for selected packets. Explain the state transition represented by each.

Do not infer application message boundaries from the TCP packet boundaries; segmentation and coalescing are independent of application writes.

Self-check: an acknowledgment confirms receipt of sequence space, not that the remote application has processed the bytes.

## 4. Verification, failure, and evidence

**Verification.** Use a packet trace or a controlled explanation of a trace to identify handshake, data sequence/acknowledgment, a retransmission if present, advertised receive window, and teardown. State what is visible and what cannot be inferred from the trace alone.

**Failure mode.** Consider a trace with a missing byte range followed by later data. Explain why TCP can acknowledge some later sequence space while the application still cannot read past the missing earlier byte in order.

**Completion criterion.** You can read sequence and acknowledgment numbers, distinguish retransmission from flow/congestion control, and explain why TCP exposes ordered bytes rather than packets or writes.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the acknowledgment number after receiving a SYN with sequence `x`.
2. Predict the acknowledgment for a segment carrying `m` bytes starting at sequence `s`.
3. Predict whether an application can read later bytes before an earlier sequence gap is filled.
4. Predict how receiver flow control differs from congestion control.
5. Predict which direction a FIN closes.
6. Predict how an RST differs from graceful FIN shutdown.
7. Predict whether a packet trace can reveal exactly how many application `Write` calls occurred.
8. State which packet or sequence evidence supports each prediction.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Annotate a TCP handshake trace with both initial sequence numbers and acknowledgment values.

2. Annotate a data trace with byte ranges, cumulative acknowledgments, and any retransmission.

3. Compare a receiver-window change with a congestion-window change and explain which endpoint or network condition each represents.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Interpret a trace with later bytes received while an earlier range is missing.

2. Explain why the TCP application sees an ordered stream only after the missing range is recovered or the connection fails.

3. Explain why a packet acknowledgment does not prove application-level processing.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one trace condition: introduce a missing sequence range, vary the advertised receive window, or close one direction first. Explain which TCP behavior changes and which stream guarantees remain.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What do TCP sequence numbers count?
2. What does a cumulative acknowledgment indicate?
3. How do flow control and congestion control differ?
4. What is a retransmission responding to?
5. How do FIN and RST differ?
6. Does TCP preserve sender write boundaries?

## Evidence to record

Create `lesson042/EVIDENCE.md` with exactly this structure:

```text
# Lesson 042 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating an ACK as confirmation that the application processed data.
- Confusing the advertised receive window with the congestion window.
- Assuming one packet corresponds to one application write.
- Treating FIN and RST as equivalent close behavior.
- Assuming a capture point shows the exact wire segmentation when offload may affect observation.

## Precision checks

1. Trace a handshake and explain sequence-space consumption.
2. Map sequence/acknowledgment values to byte ranges.
3. Distinguish loss recovery, receiver flow control, and congestion control.
4. Explain ordered delivery across a sequence gap.
5. Describe directional shutdown and reset semantics.

TCP state explains how an unreliable network path is presented as an ordered stream; application framing remains a separate layer.
## University revision bridge

For the LIG track, connect sequence spaces and state transitions to the protocol model, then distinguish reliability, ordering, receiver flow control, and congestion control.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 042 — TCP State, Reliability, and Streams without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've identified a real three-way handshake and connection teardown in an actual packet capture, you've built a working (if simplified) reliable-delivery mechanism over UDP and confirmed it survives deliberate packet loss, and you can explain — using your own failure-drill result — why reliable delivery and ordered delivery are distinct guarantees that both. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
