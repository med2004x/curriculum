# Lesson 041 — Sockets and the OS Boundary

## Threads
**Primary:** Networking/OS/Go

**Supporting:** socket lifecycle; file descriptor relation

**Math companion:** functions and state machines

**Time:** 3–4 hours. **Prerequisites:** Lesson 040.

**Time:** 3–4 hours. **Prerequisites:** Lesson 040.

## What you will be able to do

1. Explain how a socket endpoint is created and connected through the operating-system networking interface.
2. Trace the server lifecycle from bind/listen/accept and the client lifecycle from connect.
3. Explain why TCP reads do not correspond to sender write calls.
4. Implement a framing rule that can read a complete application message across partial reads.

## Retrieve first (no notes)

1. What does Kahn's algorithm emit, and how does it identify a cycle?
2. What is the difference between a topological order and critical path duration?
3. How does earliest-finish time depend on prerequisite tasks?
4. Why should an engineering report lead with the result?

## 1. Problem and model

**What problem this solves.** Build a working TCP client/server, and understand why "sending a message" over TCP requires application-level framing — TCP itself has no concept of message boundaries, only a continuous byte stream.

**The contract.** A socket is an OS-managed communication endpoint. A TCP server binds a local address, listens, and accepts connections; a client connects to a server address. Each accepted connection is a separate full-duplex byte stream. `read`/`write` or Go's `net.Conn` operations transfer bytes, not application messages.

**Worked example.** Exchange a framed message using a four-byte unsigned network-order length followed by exactly that many payload bytes. The receiver must first read the full header, validate a maximum length, then read the full payload; one read can return fewer bytes than requested.

Self-check: define the empty-message case, maximum payload length, and behavior if the peer closes before all declared bytes arrive.

## 2. Mechanism and vocabulary

**Mechanism.** **TCP is a byte stream, not a message protocol.** This is the single most important, most commonly misunderstood fact in this lesson.

**Second point.** If you `write()` "hello" and then `write()` "world" on the sending side, the receiving side's `read()` calls are **not** guaranteed to return "hello" and "world" as two separate reads — TCP may deliver them as one read of "helloworld", or split unpredictably across multiple reads, or combine multiple sender-side writes into fewer receiver-side reads.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Send two writes and receive them with a small buffer. The receiver may see both writes together or split either write across multiple reads; only the application framing rule determines message boundaries.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: open a listener, accept one connection, parse one length-prefixed message, and respond with a framed reply. Close the connection and listener on every exit path.

Trace listener state, accepted connection state, header/body byte counts, and close behavior.

Self-check: a TCP connection is ordered and reliable while established, but neither it nor `Write` defines an application message boundary.

## 4. Verification, failure, and evidence

**Verification.** Test zero-length, small, maximum permitted, oversized, and truncated messages. Compare the received payload bytes exactly with the sent payload; report short reads as valid progress until the declared length is satisfied or the peer closes.

**Failure mode.** Write a receiver that assumes one `Read` returns a full header or payload. Explain why that assumption is invalid even when a short-message test happens to pass.

**Completion criterion.** You can explain the socket lifecycle, frame messages over a byte stream, handle partial reads and premature close, and validate declared lengths before allocating or reading payloads.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict whether two writes map to exactly two reads.
2. Predict what `Read` returns when fewer bytes are currently available than requested.
3. Predict the receiver behavior when EOF arrives before the declared payload length.
4. Predict how an oversized length prefix should be handled.
5. Predict the listener, accepted-connection, and client socket roles.
6. Predict whether TCP preserves application message boundaries.
7. Predict which resources should be closed after a failed accept or read.
8. State the framing invariant before interpreting a read trace.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Write a minimal TCP server (bind, listen, accept) and client (connect) in Go or C that exchanges a single short message.

2. Extend it with one explicit boundary case: Confirm it works for a small message.

3. Rebuild the same behavior with one changed constraint: Demonstrate the "no message boundaries" fact directly: modify the client to send two separate messages via two separate `write()`/`Write()` calls in quick succession (e.g. "hello" then "world", with no delay between them), and have the server do a single `read()` immediately after accepting the connection, printing exactly what it received in that one read.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Write a *naive* receiver that assumes a single `read()` call always returns a complete message (no length-prefix framing, no read-loop — just one `read()` call per expected message) and test it against step 4's large-message, small-buffer-size scenario.

2. Confirm it receives only a fragment of the intended message (whatever fit in the first `read()` call) and either processes that truncated fragment as if it were complete, or errors out unexpectedly depending on what the fragment happens to contain.

3. Explain why this naive version might have appeared to work correctly during earlier, smaller-message testing (steps 1-2), lulling a developer into believing the assumption was safe, when in fact it was only ever true by coincidence of small message size and favorable timing, not because TCP guarantees it.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one framing rule: use a delimiter with escaping instead of a length prefix, or lower the maximum payload. Explain how the receiver recognizes message completion and handles malformed input.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What are the server and client socket lifecycle steps for TCP?
2. Why are TCP reads not message reads?
3. What framing fields identify a length-prefixed message?
4. How should the receiver handle a partial header or payload?
5. What should happen if EOF arrives before the declared length?
6. Why must the declared length be bounded?

## Evidence to record

Create `lesson041/EVIDENCE.md` with exactly this structure:

```text
# Lesson 041 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Assuming one `Read` fills the supplied buffer.
- Assuming TCP preserves boundaries between `Write` calls.
- Allocating an unbounded payload based on an untrusted length prefix.
- Treating premature EOF as a complete message.
- Forgetting to close the listener or accepted connection on error paths.

## Precision checks

1. Trace bind/listen/accept and connect on the correct endpoints.
2. State the framing contract and maximum payload.
3. Explain how short reads are accumulated.
4. Distinguish complete message from truncated message.
5. Explain why a TCP stream provides no message boundaries.

The OS provides a reliable byte stream; the application must define, validate, and parse message framing.
## University revision bridge

For the LIG track, connect the socket API to the OS boundary and distinguish transport guarantees from application-protocol framing.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the TCP socket lifecycle, frame messages with a validated bounded length prefix, handle partial reads and early EOF, and explain why one read cannot be assumed to equal one application message.
