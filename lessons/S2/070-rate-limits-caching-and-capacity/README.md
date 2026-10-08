# Lesson 070 — Rate Limits, Caching, and Capacity

## Threads
**Primary:** API/Performance/Math

**Supporting:** token buckets; TTLs; admission control

**Math companion:** probability and queues

**Time:** 3–4 hours. **Prerequisites:** Lesson 069.

**Time:** 3–4 hours. **Prerequisites:** Lesson 069.

## What you will be able to do

1. Implement and compare fixed-window and token-bucket rate limiters.
2. Define the rate-limit key and explain how identity, IP, and shared proxies affect fairness.
3. Return useful limit/reset metadata without leaking sensitive identity data.
4. Explain how caching changes load, freshness, and capacity assumptions.

## Retrieve first (no notes)

1. How do authentication and authorization differ?
2. Why is a role insufficient to authorize access to a particular record?
3. What must be checked before trusting a signed token?
4. Why should bearer credentials not appear in logs?

## 1. Problem and model

**What problem this solves.** Implement API rate limiting using the token bucket algorithm, understand why it's preferred over naive fixed-window counting, and correctly communicate limit state to clients via standard headers so well-behaved clients can self-regulate rather than just getting rejected.

**The contract.** State the policy in terms of key, capacity, refill/window rate, burst behavior, and response. A token bucket with capacity `B` and refill rate `r` allows bursts up to available tokens while limiting sustained average rate to `r`. A distributed service needs shared or partitioned limiter state; process-local counters do not enforce a global limit across replicas.

**Worked example.** Start with the smallest concrete case described by the source material: Implement fixed-window rate limiting (a simple per-client counter reset every N seconds) and demonstrate its boundary-gaming flaw directly: send a burst of requests timed to straddle a window boundary (most just before the reset, the rest just after) and confirm you can exceed the intended rate within a short overall time span, despite each individual window.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Fixed windows reset a counter at boundaries and permit a burst on each side of a reset. A token bucket stores tokens up to capacity and replenishes as time passes; request processing refills based on elapsed time, clamps at capacity, then atomically consumes or rejects.

**Second point.** Rate-limit responses commonly use `429 Too Many Requests`; `Retry-After` can communicate a retry delay. Limit/reset headers are useful metadata, but exact standardized header names and semantics depend on the API's chosen specification; define them consistently.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Use authenticated user or API-key identity as the primary limiter key when available, optionally adding an IP-level abuse limit. Document shared-NAT behavior and protect the key store from unbounded attacker-controlled cardinality.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Implement fixed-window rate limiting (a simple per-client counter reset every N seconds) and demonstrate its boundary-gaming flaw directly: send a burst of requests timed to straddle a window boundary (most just before the reset, the rest just after) and confirm you can exceed the intended rate within a short overall time span, despite each individual window.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** With an explicit clock and policy parameters, calculate tokens before/after each request and compare the fixed-window and token-bucket allowed counts for the same burst. Test concurrent requests against the same key and state whether the limiter update is atomic.

**Failure mode.** Key only by IP and simulate two users behind one shared address; then key by authenticated identity and explain the fairness change. Also test spoofable forwarding headers: only trust forwarded client IP metadata from configured proxies.

**Completion criterion.** You can model bucket state, compare burst behavior, choose an appropriate limiter key, and explain local versus distributed enforcement and client retry metadata.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict the bucket balance after a burst and elapsed refill interval.
2. Predict fixed-window allowance for requests straddling a boundary.
3. Predict what happens when a token bucket reaches capacity during idle time.
4. Predict how two users behind one NAT interact with an IP-only key.
5. Predict what changes when limiter state is local to each service replica.
6. Predict which metadata helps a client retry responsibly.
7. Predict the risk of trusting a client-supplied forwarding header directly.
8. State the atomicity requirement for simultaneous requests using one key.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Implement a fixed-window counter and reproduce the boundary burst with a controllable clock.

2. Implement a token bucket with capacity, refill rate, and last-refill timestamp.

3. Test burst, idle refill, empty bucket, shared-key concurrency, and separate logical client identities.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Key by client IP and show two users behind one NAT consuming the same allowance.

2. Switch to authenticated identity/API key and compare per-principal fairness while retaining an IP-level abuse safeguard if required.

3. Explain why forwarding headers must come only from explicitly trusted proxies.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one policy: use an authenticated principal plus route rather than IP, or lower the bucket capacity while keeping refill rate fixed. Recompute the burst allowance and discuss fairness/capacity.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does bucket capacity control?
2. What does refill rate control?
3. Why can fixed windows allow boundary bursts?
4. What limiter key is appropriate for authenticated API clients?
5. Why are process-local limits insufficient across replicas?
6. What does `Retry-After` communicate?

## Evidence to record

Create `lesson070/EVIDENCE.md` with exactly this structure:

```text
# Lesson 070 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Ignoring the fixed-window boundary burst.
- Keying only by IP despite shared NATs or spoofable proxy headers.
- Assuming a per-process counter is global across replicas.
- Updating token state non-atomically under concurrent load.
- Returning a retry instruction that contradicts the actual refill/reset policy.

## Precision checks

1. State limiter key, capacity, refill/window rate, and burst behavior.
2. Compute bucket state after a sequence of timestamps and requests.
3. Explain identity fairness and proxy trust.
4. Describe distributed limiter state requirements.
5. Document the client's retry metadata and cache freshness policy if caching is involved.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 070 — Rate Limits, Caching, and Capacity without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've directly demonstrated fixed-window rate limiting's boundary-gaming flaw with a real, exploitable test case, you've implemented and verified token bucket limiting correctly closes that gap using the same test scenario, you've implemented and tested the standard rate-limit response headers with a client that successfully self-regulates using them, and y. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
