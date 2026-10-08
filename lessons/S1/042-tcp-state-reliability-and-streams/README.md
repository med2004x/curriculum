# Lesson 042 — TCP State, Reliability, and Streams

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand precisely what TCP guarantees (reliable, ordered, byte-stream delivery) and the mechanisms — the three-way handshake, retransmission, windowing — that provide those guarantees over an unreliable network layer beneath it.

**The contract.** TCP presents a reliable, ordered byte stream while a connection is functioning; it does not preserve application write boundaries. TCP detects loss through acknowledgments and retransmits data, numbers bytes in sequence space, and uses flow control and congestion control for distinct purposes. IP itself does not promise TCP-style reliable ordered delivery.

**Worked example.** Read a TCP exchange in a packet trace: client SYN with initial sequence number `x`; server SYN-ACK acknowledging `x+1` and providing sequence number `y`; client ACK acknowledging `y+1`. SYN consumes sequence space. Capture details can be affected by offload and capture location, so interpret flags and sequence numbers rather than assuming one packet equals one application write.

## 2. Mechanism and vocabulary

**Mechanism.** TCP sequence numbers identify byte positions. Cumulative acknowledgments indicate the next expected sequence number; loss can be inferred from duplicate acknowledgments or timeout and trigger retransmission. TCP reorders received bytes before presenting them to the application.

**Second point.** Flow control uses the receiver-advertised window to limit data in flight based on receiver capacity. Congestion control limits sending based on inferred network capacity. They solve different problems.

Each TCP direction closes independently. FIN indicates no more bytes will be sent in that direction; ACK confirms sequence-space receipt. A common close involves separate FIN/ACK exchanges, while RST aborts a connection. Simultaneous close and implementation details can change the exact trace.

**Worked example.** In a packet trace, identify one data segment, its sequence range, the acknowledgment that covers it, and any retransmission. Then identify how each direction's shutdown is represented and whether the connection ends through FIN-based close or reset.

## 3. Worked case from the reference material

For a trace, record sender/receiver, flags, sequence number, acknowledgment number, advertised window, and payload length for selected packets. Explain the state transition represented by each.

Do not infer application message boundaries from the TCP packet boundaries; segmentation and coalescing are independent of application writes.

## 4. Verification, failure, and evidence

**Verification.** Use a packet trace or a controlled explanation of a trace to identify handshake, data sequence/acknowledgment, a retransmission if present, advertised receive window, and teardown. State what is visible and what cannot be inferred from the trace alone.

**Failure mode.** Consider a trace with a missing byte range followed by later data. Explain why TCP can acknowledge some later sequence space while the application still cannot read past the missing earlier byte in order.

**Completion criterion.** You can read sequence and acknowledgment numbers, distinguish retransmission from flow/congestion control, and explain why TCP exposes ordered bytes rather than packets or writes.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
