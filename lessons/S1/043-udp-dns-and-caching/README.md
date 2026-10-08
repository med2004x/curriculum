# Lesson 043 — UDP, DNS, and Caching

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand UDP as TCP's deliberately unreliable counterpart and why that's sometimes the right tradeoff, and implement a minimal DNS client to see name resolution as an actual protocol exchange rather than a black-box library call.

**The contract.** UDP preserves datagram boundaries for each delivered message but does not guarantee delivery, ordering, or duplicate suppression. It has no transport-level handshake. Applications that need reliability, congestion control, or ordering must obtain those properties from another layer or tolerate their absence.

**Worked example.** Query an A record and an AAAA record with a DNS client. Identify question name/type, answer data, TTL, and response code. Compare a successful answer with a deliberately nonexistent name that returns `NXDOMAIN`.

## 2. Mechanism and vocabulary

**Mechanism.** DNS is a message protocol. Traditional DNS queries commonly use UDP; TCP is also specified and used, including when a response is truncated or transport policy requires it. DNS can also be carried by encrypted transports such as DNS over TLS or DNS over HTTPS. Avoid treating UDP as the only DNS transport.

**Second point.** DNS caching stores answers according to their TTL and cache policy. A resolver may return a remaining TTL that decreases for a cached answer, but one observed TTL change alone does not prove which upstream path was used; caches, resolver policy, and timing affect the observation.

UDP checksum behavior, DNS truncation, retry policy, and transport choice are separate concerns. A lost UDP DNS response may trigger a retry at the resolver/client; it is not retransmitted by UDP itself.

**Worked example.** Query through the same resolver twice and compare answers and TTLs. Treat a decreasing TTL as evidence consistent with a cached response, not conclusive proof; the resolver could query upstream again, return a different answer, or apply its own TTL policy.

## 3. Worked case from the reference material

Worked example: Use `dig` (or `nslookup`) to query a real domain's A record (its IPv4 address) and observe the raw output, including the TTL value returned.

## 4. Verification, failure, and evidence

**Verification.** Record the query type, resolver, answer set, TTL, and response code. Compare the result with a second query using the same resolver and note that different resolvers or times can legitimately yield different answers.

**Failure mode.** Distinguish an `NXDOMAIN` response from a valid response with no records of the requested type (`NOERROR` with an empty answer), a timeout, and a malformed response.

**Completion criterion.** You can explain UDP's limits, parse DNS response fields, distinguish negative answers from failures, and describe the role and limitations of resolver caching.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
