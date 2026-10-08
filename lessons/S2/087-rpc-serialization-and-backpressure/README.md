# Lesson 087 — RPC, Serialization, and Backpressure

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build a minimal RPC (Remote Procedure Call) system with proper timeouts and retries, and solve Lesson 1's exact duplicate-request problem with idempotency — turning the previous lesson's abstract danger into a concretely handled case.

**The contract.** An RPC call serializes a request, sends it over a network, waits for a response, and decodes it. The call may time out after the server has performed its side effect; transport failure cannot by itself reveal whether processing occurred. Define method/version compatibility and explicit error semantics.

**Worked example.** Start with the smallest concrete case described by the source material: Build a minimal RPC client/server (Go's built-in `net/rpc` package is a reasonable, quick way to get a working RPC mechanism without hand-rolling wire-format parsing, though you're welcome to build directly on networking Lesson 1's raw sockets/framing if you want the deeper exercise) with at least one method that takes an argument and returns a result.

## 2. Mechanism and vocabulary

**Mechanism.** Use a context/deadline for bounded client waiting, but remember cancellation does not guarantee the remote server stops. The server must observe cancellation where possible, and side-effecting methods need an idempotency/reconciliation strategy if callers retry.

**Second point.** Serialization is a protocol contract: field names/numbers, types, optional values, versioning, and size limits must be compatible across sender and receiver. Unknown/optional fields and schema evolution should have documented behavior.

**Worked example.** Add a deadline to the RPC client, bound the server's work queue, and define what happens when the queue is full: reject, shed load, or wait for a bounded duration. Expose queue depth/rejection metrics without allowing unbounded accumulation.

## 3. Worked case from the reference material

Worked example: Build a minimal RPC client/server (Go's built-in `net/rpc` package is a reasonable, quick way to get a working RPC mechanism without hand-rolling wire-format parsing, though you're welcome to build directly on networking Lesson 1's raw sockets/framing if you want the deeper exercise) with at least one method that takes an argument and returns a result.

## 4. Verification, failure, and evidence

**Verification.** Test normal response, timeout, server error, malformed payload, full queue, and duplicate request ID. Record client-visible outcomes and server-side side-effect count; specify the conditions under which a retry is safe.

**Failure mode.** Disable request-ID deduplication and simulate a processed request whose response is lost. Show that a retry can duplicate the effect, then restore deduplication and explain its retention/state assumptions.

**Completion criterion.** You can state the RPC wire/error contract, bound waiting and queue growth, and explain why idempotency makes selected retries safe without promising exactly-once transport delivery.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
