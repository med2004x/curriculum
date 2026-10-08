// STUDENT PROBLEM SET — Lesson 043.
// Read the lecture, then solve these tasks before opening any worked examples:
// 1. Build the smallest useful version: Use `dig` (or `nslookup`) to query a real domain's A record (its IPv4 address) and observe the raw output, including the TTL value returned.
// 2. Extend it with one explicit boundary case: Run the same query again immediately and note whether the TTL you observe has decreased (most resolvers report the *remaining* TTL of a cached answer, decreasing in real time, which is itself direct evidence the answer came from cache rather than a fresh upstream query).
// 3. Rebuild the same behavior with one changed constraint: Use `dig +trace` (which performs and displays the full recursive resolution process, from root servers down to the authoritative answer) on a real domain, and identify, in the output, the transition from root server responses to TLD server responses to the final authoritative answer — write down, in your own words, which server in the trace was the *authorit. Do not open the reference implementation until your own attempt has been saved.
// Run the problem-set command in ../WORK.md and leave the supplied checks unchanged.
// PROBLEM STARTER — Lesson 043: UDP, DNS, and Caching
// Your task: define cache hit/miss behavior and expiration policy.
// Read ../README.md for the concept; follow ../WORK.md for the problem and checks. demo_* files are worked examples, not your submission.
package main

import "net"

func cachedLookup(cache map[string]string, name string) (string, bool) {
    // TODO: implement the assigned behavior.
    return "", false
}

func main() {
	_, _ = net.ResolveUDPAddr("udp", "127.0.0.1:5353")
}

// Keep starter imports available while the student implementation is unfinished.
var _ = net.ResolveUDPAddr
