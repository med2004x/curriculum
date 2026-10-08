# Lesson 070 — Rate Limits, Caching, and Capacity

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Implement API rate limiting using the token bucket algorithm, understand why it's preferred over naive fixed-window counting, and correctly communicate limit state to clients via standard headers so well-behaved clients can self-regulate rather than just getting rejected.

**The contract.** State the policy in terms of key, capacity, refill/window rate, burst behavior, and response. A token bucket with capacity `B` and refill rate `r` allows bursts up to available tokens while limiting sustained average rate to `r`. A distributed service needs shared or partitioned limiter state; process-local counters do not enforce a global limit across replicas.

**Worked example.** Start with the smallest concrete case described by the source material: Implement fixed-window rate limiting (a simple per-client counter reset every N seconds) and demonstrate its boundary-gaming flaw directly: send a burst of requests timed to straddle a window boundary (most just before the reset, the rest just after) and confirm you can exceed the intended rate within a short overall time span, despite each individual window.

## 2. Mechanism and vocabulary

**Mechanism.** Fixed windows reset a counter at boundaries and permit a burst on each side of a reset. A token bucket stores tokens up to capacity and replenishes as time passes; request processing refills based on elapsed time, clamps at capacity, then atomically consumes or rejects.

**Second point.** Rate-limit responses commonly use `429 Too Many Requests`; `Retry-After` can communicate a retry delay. Limit/reset headers are useful metadata, but exact standardized header names and semantics depend on the API's chosen specification; define them consistently.

**Worked example.** Use authenticated user or API-key identity as the primary limiter key when available, optionally adding an IP-level abuse limit. Document shared-NAT behavior and protect the key store from unbounded attacker-controlled cardinality.

## 3. Worked case from the reference material

Worked example: Implement fixed-window rate limiting (a simple per-client counter reset every N seconds) and demonstrate its boundary-gaming flaw directly: send a burst of requests timed to straddle a window boundary (most just before the reset, the rest just after) and confirm you can exceed the intended rate within a short overall time span, despite each individual window.

## 4. Verification, failure, and evidence

**Verification.** With an explicit clock and policy parameters, calculate tokens before/after each request and compare the fixed-window and token-bucket allowed counts for the same burst. Test concurrent requests against the same key and state whether the limiter update is atomic.

**Failure mode.** Key only by IP and simulate two users behind one shared address; then key by authenticated identity and explain the fairness change. Also test spoofable forwarding headers: only trust forwarded client IP metadata from configured proxies.

**Completion criterion.** You can model bucket state, compare burst behavior, choose an appropriate limiter key, and explain local versus distributed enforcement and client retry metadata.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
