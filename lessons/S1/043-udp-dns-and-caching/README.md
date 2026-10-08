# Lesson 043 — UDP, DNS, and Caching

## Threads
**Primary:** Networking/Linux

**Supporting:** datagrams; name resolution; TTLs

**Math companion:** probability and caching

**Time:** 3–4 hours. **Prerequisites:** Lesson 042.

**Time:** 3–4 hours. **Prerequisites:** Lesson 042.

## What you will be able to do

1. Explain UDP datagrams and what UDP does not guarantee.
2. Read a DNS question and response for record type, answer, TTL, and response code.
3. Distinguish a DNS negative answer from a timeout or transport failure.
4. Explain how caching can reduce repeated lookup work and why TTL observations depend on the resolver path.

## Retrieve first (no notes)

1. What is a TCP byte stream?
2. What does TCP sequence numbering support?
3. How does TCP distinguish graceful close from reset?
4. Why does an application need framing over a byte stream?

## 1. Problem and model

**What problem this solves.** Understand UDP as TCP's deliberately unreliable counterpart and why that's sometimes the right tradeoff, and implement a minimal DNS client to see name resolution as an actual protocol exchange rather than a black-box library call.

**The contract.** UDP preserves datagram boundaries for each delivered message but does not guarantee delivery, ordering, or duplicate suppression. It has no transport-level handshake. Applications that need reliability, congestion control, or ordering must obtain those properties from another layer or tolerate their absence.

**Worked example.** Query an A record and an AAAA record with a DNS client. Identify question name/type, answer data, TTL, and response code. Compare a successful answer with a deliberately nonexistent name that returns `NXDOMAIN`.

Self-check: a DNS name can have multiple addresses or no address of the requested type; do not assume one name maps to exactly one IP.

## 2. Mechanism and vocabulary

**Mechanism.** DNS is a message protocol. Traditional DNS queries commonly use UDP; TCP is also specified and used, including when a response is truncated or transport policy requires it. DNS can also be carried by encrypted transports such as DNS over TLS or DNS over HTTPS. Avoid treating UDP as the only DNS transport.

**Second point.** DNS caching stores answers according to their TTL and cache policy. A resolver may return a remaining TTL that decreases for a cached answer, but one observed TTL change alone does not prove which upstream path was used; caches, resolver policy, and timing affect the observation.

UDP checksum behavior, DNS truncation, retry policy, and transport choice are separate concerns. A lost UDP DNS response may trigger a retry at the resolver/client; it is not retransmitted by UDP itself.

**Worked example.** Query through the same resolver twice and compare answers and TTLs. Treat a decreasing TTL as evidence consistent with a cached response, not conclusive proof; the resolver could query upstream again, return a different answer, or apply its own TTL policy.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Use `dig` (or `nslookup`) to query a real domain's A record (its IPv4 address) and observe the raw output, including the TTL value returned.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record the query type, resolver, answer set, TTL, and response code. Compare the result with a second query using the same resolver and note that different resolvers or times can legitimately yield different answers.

**Failure mode.** Distinguish an `NXDOMAIN` response from a valid response with no records of the requested type (`NOERROR` with an empty answer), a timeout, and a malformed response.

**Completion criterion.** You can explain UDP's limits, parse DNS response fields, distinguish negative answers from failures, and describe the role and limitations of resolver caching.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict what UDP does if a datagram is lost or reordered.
2. Predict the difference between querying A and AAAA records.
3. Predict the response distinction between `NXDOMAIN` and an empty successful answer.
4. Predict what a timeout tells you and what it does not tell you about name existence.
5. Predict how a cached TTL may change across repeated queries.
6. Predict why two resolvers may return different valid address sets.
7. Predict why DNS may use TCP after a truncated UDP response.
8. State which transport guarantee comes from UDP and which behavior belongs to DNS/client policy.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Use `dig` (or `nslookup`) to query a real domain's A record (its IPv4 address) and observe the raw output, including the TTL value returned.

2. Extend it with one explicit boundary case: Run the same query again immediately and note whether the TTL you observe has decreased (most resolvers report the *remaining* TTL of a cached answer, decreasing in real time, which is itself direct evidence the answer came from cache rather than a fresh upstream query).

3. Rebuild the same behavior with one changed constraint: Use `dig +trace` (which performs and displays the full recursive resolution process, from root servers down to the authoritative answer) on a real domain, and identify, in the output, the transition from root server responses to TLD server responses to the final authoritative answer — write down, in your own words, which server in the trace was the *authorit.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Modify your step 3 client to query a domain that doesn't exist (e.g. a deliberately made-up subdomain), and observe the response's flags/response-code field (an `NXDOMAIN` response, encoded in the DNS response header) rather than a valid IP address.

2. Parse this response code explicitly in your client and confirm you can distinguish "this domain doesn't exist" from "no response arrived" (a timeout) and from "a valid IP was returned" — three genuinely different outcomes a real DNS client must handle differently, and a common source of bugs in hand-rolled or poorly-tested DNS handling code that only tests t.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 043 — UDP, DNS, and Caching: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What delivery guarantees does UDP provide?
2. What does a DNS A record contain compared with an AAAA record?
3. What does a DNS TTL represent?
4. How does `NXDOMAIN` differ from `NOERROR` with an empty answer?
5. Why can a DNS query use TCP?
6. Why does a timeout not prove that a domain is nonexistent?

## Evidence to record

Create `lesson043/EVIDENCE.md` with exactly this structure:

```text
# Lesson 043 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Claiming UDP retransmits lost data.
- Treating DNS as exclusively UDP-based.
- Assuming a repeated query must return the same address set.
- Treating a decreasing TTL as conclusive proof of a particular cache path.
- Confusing `NXDOMAIN`, an empty answer, and a transport timeout.

## Precision checks

1. Identify the transport and record type for each query.
2. Record answer data, TTL, and DNS response code.
3. Distinguish authoritative response semantics from local resolver cache behavior.
4. Classify negative answer, empty answer, timeout, and malformed response separately.
5. Explain which reliability behavior is implemented by the client rather than UDP.

DNS answer semantics and transport behavior are separate layers; record both when diagnosing a lookup.
## University revision bridge

For the LIG track, connect datagram semantics to transport design and DNS caching to TTL-based state, while distinguishing resolver policy from protocol guarantees.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 043 — UDP, DNS, and Caching without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've traced a real domain's full recursive resolution path using `dig +trace` and correctly identified the authoritative answer within it, your from-scratch DNS client correctly parses a real response to extract an IP address, and you've handled the `NXDOMAIN` case explicitly and can explain why distinguishing "doesn't exist" from "timed out" from "resolve. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
