# Lesson 041 — Sockets and the OS Boundary

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build a working TCP client/server, and understand why "sending a message" over TCP requires application-level framing — TCP itself has no concept of message boundaries, only a continuous byte stream.

**The contract.** A socket is an OS-managed communication endpoint. A TCP server binds a local address, listens, and accepts connections; a client connects to a server address. Each accepted connection is a separate full-duplex byte stream. `read`/`write` or Go's `net.Conn` operations transfer bytes, not application messages.

**Worked example.** Exchange a framed message using a four-byte unsigned network-order length followed by exactly that many payload bytes. The receiver must first read the full header, validate a maximum length, then read the full payload; one read can return fewer bytes than requested.

## 2. Mechanism and vocabulary

**Mechanism.** **TCP is a byte stream, not a message protocol.** This is the single most important, most commonly misunderstood fact in this lesson.

**Second point.** If you `write()` "hello" and then `write()` "world" on the sending side, the receiving side's `read()` calls are **not** guaranteed to return "hello" and "world" as two separate reads — TCP may deliver them as one read of "helloworld", or split unpredictably across multiple reads, or combine multiple sender-side writes into fewer receiver-side reads.

**Worked example.** Send two writes and receive them with a small buffer. The receiver may see both writes together or split either write across multiple reads; only the application framing rule determines message boundaries.

## 3. Worked case from the reference material

Worked example: open a listener, accept one connection, parse one length-prefixed message, and respond with a framed reply. Close the connection and listener on every exit path.

Trace listener state, accepted connection state, header/body byte counts, and close behavior.

## 4. Verification, failure, and evidence

**Verification.** Test zero-length, small, maximum permitted, oversized, and truncated messages. Compare the received payload bytes exactly with the sent payload; report short reads as valid progress until the declared length is satisfied or the peer closes.

**Failure mode.** Write a receiver that assumes one `Read` returns a full header or payload. Explain why that assumption is invalid even when a short-message test happens to pass.

**Completion criterion.** You can explain the socket lifecycle, frame messages over a byte stream, handle partial reads and premature close, and validate declared lengths before allocating or reading payloads.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
