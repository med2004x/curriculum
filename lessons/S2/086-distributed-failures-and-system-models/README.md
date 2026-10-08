# Lesson 086 — Distributed Failures and System Models

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand why "partial failure" — the defining problem of distributed systems — has no equivalent in single-machine programming, and build the vocabulary (crash, omission, partition, delay) to reason precisely about what can go wrong.

**The contract.** A distributed client observes messages, responses, and elapsed time, not the remote process's internal state. A timeout means no response was observed before a deadline; the request may have been lost, delayed, processed with its response lost, or still be running.

**Worked example.** Start with the smallest concrete case described by the source material: For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself.

## 2. Mechanism and vocabulary

**Mechanism.** Define the system model: which components may crash, whether messages may be lost, duplicated, or reordered, what delay bounds exist, and how nodes recover. Safety and liveness claims depend on these assumptions.

**Second point.** A network partition separates communication paths without proving either endpoint has crashed. From one side, a sufficiently long delay can be indistinguishable from a failed or partitioned peer.

**Worked example.** For request lost, server crash before processing, server processes but response is lost, and response delayed: state the client-visible result under timeout-only behavior and identify what extra operation ID/status lookup could clarify.

## 3. Worked case from the reference material

Worked example: For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself.

## 4. Verification, failure, and evidence

**Verification.** Use a controlled client/server model to record request send, server-side effect, response loss/delay, and client timeout. Compare client symptoms and state which internal outcomes remain ambiguous.

**Failure mode.** Retry an append after the server committed but the response was lost. Trace the duplicate effect, then explain why an idempotency key or operation-status protocol is needed before retrying safely.

**Completion criterion.** You can classify failure modes, state the system model, explain timeout ambiguity, and identify when idempotency or reconciliation is needed to make retries safe.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
